package entities

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"time"
)

type Auditor interface {
	LogPresentation() (map[string]any, error)
	PrimaryKey() string
	EntityName() string
	Copy() Auditor
	PrimaryFields() map[string]any
}

func ExtractActorFromContext(ctx context.Context) (User, error) {
	actorCtx := ctx.Value("user")
	if actorCtx == nil {
		return User{}, errors.New("user doesn't exist")
	}
	actor, ok := actorCtx.(User)
	if !ok {
		return User{}, errors.New("user is not valid")
	}
	return actor, nil
}

func ExtractColumns(data map[string]any, columns []string) (result map[string]any) {
	result = make(map[string]any)
	for _, val := range columns {
		result[val] = data[val]
	}
	return
}

func UpdatedColumns(oldData Auditor, newData Auditor) (result []string, err error) {
	var (
		oldMap map[string]any
		newMap map[string]any
	)
	if err := mapstructure.Decode(oldData, &oldMap); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(newData, &newMap); err != nil {
		return nil, err
	}
	for key, newValue := range newMap {
		if oldMap[key] != newValue {
			result = append(result, key)
		}
	}
	return result, nil
}

func InsertAudit(tx *gorm.DB, action, entity, entityId string, dataBefore, dataAfter any) error {
	actor, err := ExtractActorFromContext(tx.Statement.Context)
	if err != nil {
		return err
	}
	audit := Audit{
		DateTime: time.Now(),
		Username: actor.Username,
		Action:   action,
		Entity:   entity,
		EntityID: entityId,
	}
	if dataBefore != nil {
		dataBeforeJson, err := json.Marshal(dataBefore)
		if err != nil {
			return err
		}
		audit.DataBefore = string(dataBeforeJson)
	}
	if dataAfter != nil {
		dataAfterJson, err := json.Marshal(dataAfter)
		if err != nil {
			return err
		}
		audit.DataAfter = string(dataAfterJson)
	}
	if err := tx.Create(&audit).Error; err != nil {
		return err
	}
	return nil
}

func AuditCreate(tx *gorm.DB, data Auditor) error {
	dataAfter, err := data.LogPresentation()
	if err != nil {
		return err
	}
	return InsertAudit(tx, "CREATE", data.EntityName(), data.PrimaryKey(), nil, dataAfter)
}

func AuditUpdate(tx *gorm.DB, data Auditor) error {
	dataCopy := data.Copy()
	if err := tx.Where(data.PrimaryFields()).Find(&dataCopy).Error; err != nil {
		return err
	}
	updatedColumn, err := UpdatedColumns(dataCopy, data)
	if err != nil {
		return err
	}
	dataAfter, err := data.LogPresentation()
	if err != nil {
		return err
	}
	dataBefore, err := dataCopy.LogPresentation()
	if err != nil {
		return err
	}
	dataAfter = ExtractColumns(dataAfter, updatedColumn)
	dataBefore = ExtractColumns(dataBefore, updatedColumn)
	return InsertAudit(tx, "UPDATE", data.EntityName(), data.PrimaryKey(), dataBefore, dataAfter)
}

func AuditDelete(tx *gorm.DB, data Auditor) error {
	dataCopy := data.Copy()
	if err := tx.Where(data.PrimaryFields()).Find(&dataCopy).Error; err != nil {
		return err
	}
	dataBefore, err := dataCopy.LogPresentation()
	if err != nil {
		return err
	}
	return InsertAudit(tx, "DELETE", data.EntityName(), data.PrimaryKey(), dataBefore, nil)
}
