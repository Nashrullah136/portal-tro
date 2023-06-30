package server_utilization

import (
	"nashrul-be/crm/entities"
	"strconv"
	"time"
)

type ThresholdFunc func(utilization entities.ServerUtilization) bool

func CheckCPU(utilization entities.ServerUtilization) bool {
	cpu, _ := strconv.Atoi(utilization.CpuPercentage)
	return cpu < 90
}

func CheckMemory(utilization entities.ServerUtilization) bool {
	memory, _ := strconv.Atoi(utilization.MemoryUsage)
	return memory < 90
}

func CheckUptime(utilization entities.ServerUtilization) bool {
	upTime, _ := strconv.Atoi(utilization.SystemUptime)
	return upTime < int(30*24*time.Hour)
}

func CheckDisk(utilization entities.ServerUtilization) bool {
	for _, disk := range utilization.Disks {
		percent, _ := strconv.Atoi(disk)
		if percent >= 95 {
			return false
		}
	}
	return true
}
