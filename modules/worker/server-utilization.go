package worker

import (
	"github.com/go-co-op/gocron"
	"log"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/zabbix"
	"time"
)

func UpdateLastDataServerUtil(cache zabbix.Cache, api zabbix.API) error {
	wib, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(wib)
	_, err := s.Every(1).Minute().Do(func() {
		template, err := cache.GetTemplate()
		if err != nil {
			log.Fatal("can't get template for server util")
		}
		itemIds, err := cache.GetItemIds()
		if err != nil {
			log.Fatal("can't get item ids for server util")
		}
		result := make([]entities.ServerUtilization, 18)
		history, err := api.GetHistoryFromItem(itemIds)
		if err != nil {
			log.Fatal("Failed to access zabbix server")
		}
		historyMap := make(map[string]string)
		for _, val := range history {
			historyMap[val.ItemId] = val.Value
		}
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
		if err := cache.SetLastValue(result); err != nil {
			log.Fatal("Failed to save last value result to redis")
		}
	})
	if err != nil {
		return err
	}
	s.StartAsync()
	return nil
}
