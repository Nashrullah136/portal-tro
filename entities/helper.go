package entities

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

func SystemContext() context.Context {
	ctx := context.Background()
	user := User{
		Username: "SYSTEM",
	}
	ctx = context.WithValue(ctx, "actor", user)
	return ctx
}

func ExtractActorFromContext(ctx context.Context) (User, error) {
	actorCtx := ctx.Value("actor")
	if actorCtx == nil {
		return User{}, errors.New("actor doesn't exist")
	}
	actor, ok := actorCtx.(User)
	if !ok {
		return User{}, errors.New("actor is not valid")
	}
	return actor, nil
}

func CreateAudit(tx *gorm.DB, action, entity, entityId string, dataBefore, dataAfter any) error {
	actor, err := ExtractActorFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err.Error())
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
			log.Println(err.Error())
			return err
		}
		audit.DataBefore = string(dataBeforeJson)
	}
	if dataAfter != nil {
		dataAfterJson, err := json.Marshal(dataAfter)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		audit.DataAfter = string(dataAfterJson)
	}
	if err := tx.Create(&audit).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
