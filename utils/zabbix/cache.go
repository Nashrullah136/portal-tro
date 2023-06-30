package zabbix

import "nashrul-be/crm/entities"

var (
	template  []entities.ServerUtilization
	itemIds   []string
	lastValue []entities.ServerUtilization
)

type Cache interface {
	SetTemplate(template []entities.ServerUtilization) error
	GetTemplate() ([]entities.ServerUtilization, error)
	SetItemIds([]string) error
	GetItemIds() ([]string, error)
	SetLastValue([]entities.ServerUtilization) error
	GetLastValue() ([]entities.ServerUtilization, error)
}

func NewCache() Cache {
	return cache{}
}

type cache struct {
}

func (c cache) SetTemplate(newTemplate []entities.ServerUtilization) error {
	template = newTemplate
	return nil
}

func (c cache) GetTemplate() ([]entities.ServerUtilization, error) {
	return template, nil
}

func (c cache) SetItemIds(strings []string) error {
	itemIds = strings
	return nil
}

func (c cache) GetItemIds() ([]string, error) {
	return itemIds, nil
}

func (c cache) SetLastValue(utilization []entities.ServerUtilization) error {
	lastValue = utilization
	return nil
}

func (c cache) GetLastValue() ([]entities.ServerUtilization, error) {
	return lastValue, nil
}
