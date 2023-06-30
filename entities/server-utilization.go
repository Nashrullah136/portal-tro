package entities

type ServerUtilization struct {
	Hostname      string            `json:"hostname,omitempty"`
	CpuPercentage string            `json:"cpu_percentage,omitempty"`
	MemoryUsage   string            `json:"memory_usage,omitempty"`
	SystemUptime  string            `json:"system_uptime,omitempty"`
	Disks         map[string]string `json:"disks,omitempty"`
}
