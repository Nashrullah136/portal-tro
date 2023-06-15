package entities

import (
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"log"
	"time"
)

const ADMIN = "admin"

// User TODO: add user log presentation
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

func (u *User) LogPresentation() (result map[string]any, err error) {
	if err = mapstructure.Decode(u, &result); err != nil {
		return result, err
	}
	result["password"] = "-"
	return result, nil
}

func (u *User) PrimaryKey() string {
	return u.Username
}

func (u *User) EntityName() string {
	return "USER"
}

func (u *User) Copy() Auditor {
	result := *u
	return &result
}

func (u *User) PrimaryFields() map[string]any {
	return map[string]any{
		"username": u.Username,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	actor, err := ExtractActorFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	u.CreatedBy = actor.Username
	u.UpdatedBy = actor.Username
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	return AuditCreate(tx, u)
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if err := AuditUpdate(tx, u); err != nil {
		return err
	}
	actor, err := ExtractActorFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	u.UpdatedBy = actor.Username
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
	return AuditDelete(tx, u)
}
