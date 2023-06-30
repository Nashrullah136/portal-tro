package zabbix

type API interface {
	GetAllHost() (result []Host, err error)
	GetItemFromHosts(hostIds []string) (result []Item, err error)
	GetHistoryFromItem(itemIds []string) (result []History, err error)
}

type api struct {
	server Server
}

func NewAPI(zabbixServer Server) API {
	return api{server: zabbixServer}
}

func (a api) GetAllHost() (result []Host, err error) {
	request := map[string]any{
		"output": []string{
			"hostid",
			"host",
			"error",
		},
	}
	if err = a.server.Do(request, &result); err != nil {
		return nil, err
	}
	return result, err
}

func (a api) GetItemFromHosts(hostIds []string) (result []Item, err error) {
	request := map[string]any{
		"output": []string{
			"hostid",
			"itemid",
			"key_",
			"name",
		},
		"hostids": hostIds,
		"search": map[string]any{
			"key_": []string{
				"system.cpu.util",
				"vm.memory.util",
				"vfs.fs.size[*,pused]",
				"system.uptime",
			},
			"searchWildcardsEnabled": true,
		},
	}
	if err = a.server.Do(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (a api) GetHistoryFromItem(itemIds []string) (result []History, err error) {
	request := map[string]any{
		"output": []string{
			"itemid",
			"value",
		},
		"history":   0,
		"itemids":   itemIds,
		"sortfield": "clock",
		"sortorder": "DESC",
		"limit":     1,
	}
	if err = a.server.Do(request, &result); err != nil {
		return nil, err
	}
	return result, err
}
