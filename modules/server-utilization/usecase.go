package server_utilization

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/maps"
	"nashrul-be/crm/utils/zabbix"
	"regexp"
)

const ZabbixTemplate = "zabbix-template"
const ZabbixLastValue = "zabbix-last-value"
const ZabbixItemIds = "zabbix-item-ids"

type UseCase interface {
	RefreshHostList() error
	GetLastData() (map[string][]entities.ServerUtilization, error)
	IsValid(utilization entities.ServerUtilization) bool
}

func NewUseCase(client *redis.Client, api zabbix.API) UseCase {
	return useCase{
		redisConn: client,
		zabbixApi: api,
	}
}

type useCase struct {
	redisConn *redis.Client
	zabbixApi zabbix.API
}

func (uc useCase) RefreshHostList() error {
	diskRegex := regexp.MustCompile(`vfs.fs.size\[(.+),pused]`)
	hosts, err := uc.zabbixApi.GetAllHost()
	if err != nil {
		return err
	}
	serverUtils := make(map[string]entities.ServerUtilization, len(hosts))
	hostIds := make([]string, len(hosts))
	for index, host := range hosts {
		hostIds[index] = host.HostId
		serverUtils[host.HostId] = entities.ServerUtilization{Hostname: host.Host}
	}
	items, err := uc.zabbixApi.GetItemFromHosts(hostIds)
	if err != nil {
		return err
	}
	itemIds := make([]string, len(items))
	for index, item := range items {
		itemIds[index] = item.ItemId
		switch {
		case item.Key == "system.cpu.util":
			temp := serverUtils[item.HostId]
			temp.CpuPercentage = item.ItemId
			serverUtils[item.HostId] = temp
		case item.Key == "vm.memory.util":
			temp := serverUtils[item.HostId]
			temp.MemoryUsage = item.ItemId
			serverUtils[item.HostId] = temp
		case item.Key == "system.uptime":
			temp := serverUtils[item.HostId]
			temp.SystemUptime = item.ItemId
			serverUtils[item.HostId] = temp
		case diskRegex.Match([]byte(item.Key)):
			temp := serverUtils[item.HostId]
			diskLabel := diskRegex.FindStringSubmatch(item.Key)[1]
			temp.Disks[diskLabel] = item.ItemId
			serverUtils[item.HostId] = temp
		}
	}
	values := maps.Values(serverUtils)
	valuesJson, err := json.Marshal(values)
	if err != nil {
		return err
	}
	if status := uc.redisConn.Set(context.Background(), ZabbixTemplate, valuesJson, 0); status.Err() != nil {
		return err
	}
	itemIdsJson, err := json.Marshal(itemIds)
	if err != nil {
		return err
	}
	if status := uc.redisConn.Set(context.Background(), ZabbixItemIds, itemIdsJson, 0); status.Err() != nil {
		return err
	}
	return nil
}

func (uc useCase) GetLastData() (map[string][]entities.ServerUtilization, error) {
	jsonLastValue := uc.redisConn.Get(context.Background(), ZabbixLastValue)
	if jsonLastValue.Err() != nil {
		return nil, jsonLastValue.Err()
	}
	serverUtils := make([]entities.ServerUtilization, 18)
	if err := json.Unmarshal([]byte(jsonLastValue.String()), &serverUtils); err != nil {
		return nil, err
	}
	threshold := make([]entities.ServerUtilization, 18)
	safe := make([]entities.ServerUtilization, 18)
	for _, serverUtil := range serverUtils {
		if uc.IsValid(serverUtil) {
			safe = append(safe, serverUtil)
		} else {
			threshold = append(threshold, serverUtil)
		}
	}
	return map[string][]entities.ServerUtilization{
		"threshold": threshold,
		"safe":      safe,
	}, nil
}

func (uc useCase) IsValid(utilization entities.ServerUtilization) bool {
	checkFunc := []ThresholdFunc{CheckCPU, CheckMemory, CheckUptime, CheckDisk}
	for _, check := range checkFunc {
		if !check(utilization) {
			return false
		}
	}
	return true
}
