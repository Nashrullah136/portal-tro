package entities

import (
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"log"
	"nashrul-be/crm/utils/auditUtils"
	"time"
)

const ADMIN = "admin"
const USER = "user"

type User struct {
	ID        uint       `gorm:"autoIncrement" mapstructure:"-" json:"-"`
	Name      string     `mapstructure:"name,omitempty" json:"name,omitempty"`
	Username  string     `gorm:"primaryKey" mapstructure:"username" json:"username,omitempty"`
	Password  string     `mapstructure:"password,omitempty" json:"-"`
	RoleID    uint       `mapstructure:"role_id,omitempty" json:"role_id,omitempty"`
	Role      Role       `mapstructure:"-" json:"role,omitempty"`
	CreatedAt *time.Time `mapstructure:"-" json:"created_at,omitempty"`
	CreatedBy string     `mapstructure:"-" json:"created_by,omitempty"`
	UpdatedAt *time.Time `mapstructure:"-" json:"updated_at,omitempty"`
	UpdatedBy string     `mapstructure:"-" json:"updated_by,omitempty"`
}

func (u *User) Identity() string {
	return u.Username
}

func (u *User) IsNewUser() bool {
	return u.UpdatedAt.Sub(*u.CreatedAt) < 1*time.Second
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

func (u *User) Copy() auditUtils.Auditor {
	result := *u
	return &result
}

func (u *User) PrimaryFields() map[string]any {
	return map[string]any{
		"username": u.Username,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	actor, err := getUserFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	u.CreatedBy = actor.Identity()
	u.UpdatedBy = actor.Identity()
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	actor, err := getUserFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	auditResult, err := auditUtils.Create(&actor, u)
	if err != nil {
		return err
	}
	auditEntities := MapAuditResultToAuditEntities(auditResult)
	return tx.Create(&auditEntities).Error
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	actor, err := getUserFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	auditResult, err := auditUtils.Update(tx, &actor, u)
	if err != nil {
		return err
	}
	auditEntities := MapAuditResultToAuditEntities(auditResult)
	if err := tx.Create(&auditEntities).Error; err != nil {
		return err
	}
	u.UpdatedBy = actor.Identity()
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
	actor, err := getUserFromContext(tx.Statement.Context)
	if err != nil {
		log.Println(err)
		return err
	}
	auditResult, err := auditUtils.Delete(tx, &actor, u)
	if err != nil {
		return err
	}
	auditEntities := MapAuditResultToAuditEntities(auditResult)
	return tx.Create(&auditEntities).Error
}
