package server_utilization

import (
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/maps"
	"nashrul-be/crm/utils/zabbix"
	"regexp"
)

type Controller interface {
	RefreshHostList() error
	GetLastData() (dto.BaseResponse, error)
	IsValid(utilization entities.ServerUtilization) bool
}

func NewController(cache zabbix.Cache, api zabbix.API) Controller {
	return controller{
		cache:     cache,
		zabbixApi: api,
	}
}

type controller struct {
	cache     zabbix.Cache
	zabbixApi zabbix.API
}

func (uc controller) RefreshHostList() error {
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
		temp := serverUtils[item.HostId]
		switch {
		case item.Key == "system.cpu.util":
			temp.CpuPercentage = item.ItemId
		case item.Key == "vm.memory.util":
			temp.MemoryUsage = item.ItemId
		case item.Key == "system.uptime":
			temp.SystemUptime = item.ItemId
		case diskRegex.Match([]byte(item.Key)):
			diskLabel := diskRegex.FindStringSubmatch(item.Key)[1]
			temp.Disks[diskLabel] = item.ItemId
		}
		serverUtils[item.HostId] = temp
	}
	values := maps.Values(serverUtils)
	if err = uc.cache.SetTemplate(values); err != nil {
		return err
	}
	if err = uc.cache.SetItemIds(itemIds); err != nil {
		return err
	}
	return nil
}

func (uc controller) GetLastData() (dto.BaseResponse, error) {
	serverUtils, err := uc.cache.GetLastValue()
	if err != nil {
		return dto.ErrorInternalServerError(), err
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
	result := map[string][]entities.ServerUtilization{
		"threshold": threshold,
		"safe":      safe,
	}
	return dto.Success("Success get latest data", result), nil
}

func (uc controller) IsValid(utilization entities.ServerUtilization) bool {
	checkFunc := []ThresholdFunc{CheckCPU, CheckMemory, CheckUptime, CheckDisk}
	for _, check := range checkFunc {
		if !check(utilization) {
			return false
		}
	}
	return true
}
