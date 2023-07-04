package server_utilization

import (
	"nashrul-be/crm/entities"
	"strconv"
)

type ThresholdFunc func(utilization entities.ServerUtilization) bool

func CheckCPU(utilization entities.ServerUtilization) bool {
	cpu, _ := strconv.ParseFloat(utilization.CpuPercentage, 32)
	return cpu < 90
}

func CheckMemory(utilization entities.ServerUtilization) bool {
	memory, _ := strconv.ParseFloat(utilization.MemoryUsage, 32)
	return memory < 90
}

func CheckUptime(utilization entities.ServerUtilization) bool {
	upTime, _ := strconv.Atoi(utilization.SystemUptime)
	return upTime < 3*2628000
}

func CheckDisk(utilization entities.ServerUtilization) bool {
	for _, disk := range utilization.Disks {
		percent, _ := strconv.ParseFloat(disk, 32)
		if percent >= 95 {
			return false
		}
	}
	return true
}
