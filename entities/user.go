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
	Username  string    `gorm:"primaryKey" mapstructure:"username" json:"username,omitempty"`
	Password  string    `mapstructure:"password,omitempty" json:"-"`
	RoleID    uint      `mapstructure:"role_id,omitempty" json:"role_id,omitempty"`
	Role      Role      `mapstructure:"-" json:"role,omitempty"`
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
	audit, err := AuditCreate(tx, u)
	if err != nil {
		return err
	}
	return tx.Create(&audit).Error
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	audit, err := AuditUpdate(tx, u)
	if err != nil {
		return err
	}
	if err := tx.Create(&audit).Error; err != nil {
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
	audit, err := AuditDelete(tx, u)
	if err != nil {
		return err
	}
	return tx.Create(&audit).Error
}
