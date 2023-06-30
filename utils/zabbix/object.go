package zabbix

type Authenticated struct {
	Result string `json:"result"`
}

type Host struct {
	HostId string `json:"hostid,omitempty"`
	Host   string `json:"host,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Item struct {
	HostId string `json:"hostid,omitempty"`
	ItemId string `json:"itemid,omitempty"`
	Key    string `json:"key_,omitempty"`
	Name   string `json:"name,omitempty"`
}

type History struct {
	ItemId string `json:"itemid,omitempty"`
	Value  string `json:"value,omitempty"`
}
