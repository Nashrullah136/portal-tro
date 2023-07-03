package zabbix

const MethodHostGet = "host.get"
const MethodItemGet = "item.get"
const MethodHistoryGet = "history.get"

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
	if err = a.server.Do(MethodHostGet, request, &result); err != nil {
		return nil, err
	}
	return result, err
}

func (a api) GetItemFromHosts(hostIds []string) (result []Item, err error) {
	var temp []Item
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
				"system.uptime",
			},
		},
		"searchByAny": true,
	}
	if err = a.server.Do(MethodItemGet, request, &temp); err != nil {
		return nil, err
	}
	result = append(result, temp...)
	request = map[string]any{
		"output": []string{
			"hostid",
			"itemid",
			"key_",
			"name",
		},
		"hostids": hostIds,
		"search": map[string]any{
			"key_": []string{
				"vfs.fs.size[*",
				"*:,pused]",
			},
		},
		"searchWildcardsEnabled": true,
	}
	if err = a.server.Do(MethodItemGet, request, &temp); err != nil {
		return nil, err
	}
	result = append(result, temp...)
	return result, nil
}

func (a api) GetHistoryFromItem(itemIds []string) (result []History, err error) {
	request := map[string]any{
		"output": []string{
			"itemid",
			"lastvalue",
		},
		"itemids": itemIds,
	}
	if err = a.server.Do(MethodItemGet, request, &result); err != nil {
		return nil, err
	}
	return result, err
}
