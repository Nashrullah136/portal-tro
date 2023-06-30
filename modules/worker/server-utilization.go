package worker

import (
	"context"
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/redis/go-redis/v9"
	"log"
	"nashrul-be/crm/entities"
	serverutilization "nashrul-be/crm/modules/server-utilization"
	"nashrul-be/crm/utils/zabbix"
	"time"
)

func UpdateLastDataServerUtil(client *redis.Client, api zabbix.API) error {
	wib, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(wib)
	_, err := s.Every(1).Minute().Do(func() {
		templateJson := client.Get(context.Background(), serverutilization.ZabbixTemplate)
		if templateJson.Err() != nil {
			log.Fatal("Failed to access redis")
		}
		var template []entities.ServerUtilization
		if err := json.Unmarshal([]byte(templateJson.String()), &template); err != nil {
			log.Fatal("Malformed zabbix template")
		}
		itemIdsJson := client.Get(context.Background(), serverutilization.ZabbixItemIds)
		if itemIdsJson.Err() != nil {
			log.Fatal("Failed to access redis")
		}
		var itemIds []string
		if err := json.Unmarshal([]byte(itemIdsJson.String()), &itemIds); err != nil {
			log.Fatal("Malformed zabbix item ids")
		}
		history, err := api.GetHistoryFromItem(itemIds)
		if err != nil {
			log.Fatal("Failed to access zabbix server")
		}
		historyMap := make(map[string]string)
		for _, val := range history {
			historyMap[val.ItemId] = val.Value
		}
		result := make([]entities.ServerUtilization, 18)
		for _, serverUtil := range template {
			temp := entities.ServerUtilization{
				CpuPercentage: historyMap[serverUtil.CpuPercentage],
				MemoryUsage:   historyMap[serverUtil.MemoryUsage],
				SystemUptime:  historyMap[serverUtil.SystemUptime],
			}
			for key, val := range serverUtil.Disks {
				temp.Disks[key] = historyMap[val]
			}
			result = append(result, temp)
		}
		resultJson, err := json.Marshal(result)
		if err != nil {
			log.Fatal("Failed to marshal result")
		}
		if status := client.Set(context.Background(), serverutilization.ZabbixLastValue, resultJson, 0); status.Err() != nil {
			log.Fatal("Failed to save last value result to redis")
		}
	})
	if err != nil {
		return err
	}
	s.StartAsync()
	return nil
}
