package auditUtils

import (
	"encoding/json"
	"gorm.io/gorm"
	"nashrul-be/crm/utils/localtime"
)

type Auditor interface {
	LogPresentation() (map[string]any, error)
	PrimaryKey() string
	EntityName() string
	Copy() Auditor
	PrimaryFields() map[string]any
}

func NewAudit(actor Actor, action, entity, entityId string, dataBefore, dataAfter any) (Result, error) {
	audit := Result{
		DateTime: localtime.Now(),
		Username: actor.Identity(),
		Action:   action,
		Entity:   entity,
		EntityID: entityId,
	}
	if dataBefore != nil {
		dataBeforeJson, err := json.Marshal(dataBefore)
		if err != nil {
			return Result{}, err
		}
		audit.DataBefore = string(dataBeforeJson)
	}
	if dataAfter != nil {
		dataAfterJson, err := json.Marshal(dataAfter)
		if err != nil {
			return Result{}, err
		}
		audit.DataAfter = string(dataAfterJson)
	}
	return audit, nil
}

func Create(actor Actor, data Auditor) (Result, error) {
	dataAfter, err := data.LogPresentation()
	if err != nil {
		return Result{}, err
	}
	return NewAudit(actor, "CREATE", data.EntityName(), data.PrimaryKey(), nil, dataAfter)
}

func Update(tx *gorm.DB, actor Actor, newData Auditor) (Result, error) {
	oldData := newData.Copy()
	if err := tx.Where(newData.PrimaryFields()).Find(&oldData).Error; err != nil {
		return Result{}, err
	}
	return UpdateWithOldData(actor, newData, oldData)
}

func UpdateWithOldData(actor Actor, newData Auditor, oldData Auditor) (Result, error) {
	updatedColumn, err := UpdatedColumns(oldData, newData)
	if err != nil {
		return Result{}, err
	}
	dataAfter, err := newData.LogPresentation()
	if err != nil {
		return Result{}, err
	}
	dataBefore, err := oldData.LogPresentation()
	if err != nil {
		return Result{}, err
	}
	dataAfter = ExtractColumns(dataAfter, updatedColumn)
	dataBefore = ExtractColumns(dataBefore, updatedColumn)
	return NewAudit(actor, "UPDATE", newData.EntityName(), newData.PrimaryKey(), dataBefore, dataAfter)
}

func Delete(tx *gorm.DB, actor Actor, data Auditor) (Result, error) {
	dataCopy := data.Copy()
	if err := tx.Where(data.PrimaryFields()).Find(&dataCopy).Error; err != nil {
		return Result{}, err
	}
	dataBefore, err := dataCopy.LogPresentation()
	if err != nil {
		return Result{}, err
	}
	return NewAudit(actor, "DELETE", data.EntityName(), data.PrimaryKey(), dataBefore, nil)
}
