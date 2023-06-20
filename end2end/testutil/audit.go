package testutil

import (
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"nashrul-be/crm/entities"
	"testing"
)

func AssertAudit(t *testing.T, username, action, entity, entityID string, dataBefore, dataAfter map[string]any) {
	var (
		audit       entities.Audit
		auditBefore map[string]any
		auditAfter  map[string]any
	)
	db, _ := GetConn()
	if err := db.Order("date_time desc").First(&audit).Error; err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(t, username, audit.Username)
	assert.Equal(t, action, audit.Action)
	assert.Equal(t, entity, audit.Entity)
	assert.Equal(t, entityID, audit.EntityID)
	if dataBefore == nil {
		assert.Equal(t, "", audit.DataBefore)
	} else {
		if err := json.Unmarshal([]byte(audit.DataBefore), &auditBefore); err != nil {
			t.Fatal(err)
		}
		AssertMap(t, dataBefore, auditBefore)
	}
	if dataAfter == nil {
		assert.Equal(t, "", audit.DataAfter)
	} else {
		if err := json.Unmarshal([]byte(audit.DataAfter), &auditAfter); err != nil {
			t.Fatal(err)
		}
		AssertMap(t, dataAfter, dataAfter)
	}
}

func AssertMap(t *testing.T, source map[string]any, destiny map[string]any) {
	if len(source) != len(destiny) {
		t.Fatal("map len not same")
	}
	for key, sourceValue := range source {
		destinyValue, exist := destiny[key]
		if !exist {
			t.Fatalf("key %s doesn't exist in destiny map", key)
		}
		assert.EqualValues(t, sourceValue, destinyValue)
	}
}
