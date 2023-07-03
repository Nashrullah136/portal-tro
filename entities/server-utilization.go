package entities

type ServerUtilization struct {
	Hostname      string            `json:"hostname"`
	CpuPercentage string            `json:"cpu_percentage"`
	MemoryUsage   string            `json:"memory_usage"`
	SystemUptime  string            `json:"system_uptime"`
	Disks         map[string]string `json:"disks"`
}
