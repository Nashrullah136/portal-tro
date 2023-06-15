package entities

import (
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"log"
	"time"
)

const ADMIN = "admin"

type User struct {
	ID        uint      `mapstructure:"-" json:"-"`
	Name      string    `mapstructure:"name,omitempty" json:"name,omitempty"`
	Username  string    `gorm:"primaryKey" mapstructure:"-" json:"username,omitempty"`
	Password  string    `mapstructure:"password,omitempty" json:"password,omitempty"`
	RoleID    uint      `mapstructure:"role_id,omitempty" json:"role_id,omitempty"`
	Role      Role      `mapstructure:"-" json:"-"`
	CreatedAt time.Time `mapstructure:"-" json:"-"`
	CreatedBy string    `mapstructure:"-" json:"-"`
	UpdatedAt time.Time `mapstructure:"-" json:"-"`
	UpdatedBy string    `mapstructure:"-" json:"-"`
}

func (a *User) BeforeCreate(tx *gorm.DB) error {
	actor, err := ExtractActorFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	a.CreatedBy = actor.Username
	a.UpdatedBy = actor.Username
	return nil
}

func (a *User) AfterCreate(tx *gorm.DB) error {
	user := *a
	user.Password = "-"
	if err := CreateAudit(tx, "CREATE", "USER", a.Username, nil, user); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a *User) BeforeUpdate(tx *gorm.DB) error {
	var (
		after      map[string]any
		before     map[string]any
		beforeUser = User{Username: a.Username}
		afterUser  = *a
		sysCtx     = SystemContext()
	)
	if err := tx.WithContext(sysCtx).First(&beforeUser).Error; err != nil {
		log.Println(err)
		return err
	}
	if afterUser.Password != "" {
		afterUser.Password = "-"
		beforeUser.Password = "-"
	}
	if err := mapstructure.Decode(afterUser, &after); err != nil {
		log.Println(err)
		return err
	}
	if err := mapstructure.Decode(beforeUser, &before); err != nil {
		log.Println(err)
		return err
	}
	for key := range before {
		if _, ok := after[key]; !ok {
			delete(before, key)
		}
	}
	if err := CreateAudit(tx, "UPDATE", "USER", a.Username, before, after); err != nil {
		log.Println(err)
		return err
	}
	actor, err := ExtractActorFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	a.UpdatedBy = actor.Username
	return nil
}

func (a *User) BeforeDelete(tx *gorm.DB) error {
	user := *a
	sysCtx := SystemContext()
	if err := tx.WithContext(sysCtx).First(&user).Error; err != nil {
		log.Println(err)
		return err
	}
	user.Password = "-"
	if err := CreateAudit(tx, "DELETE", "USER", a.Username, user, nil); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
