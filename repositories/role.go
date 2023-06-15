package repositories

import (
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
)

type RoleRepositoryInterface interface {
	GetByID(id uint) (role entities.Role, err error)
}

func NewRoleRepository(db *gorm.DB) RoleRepositoryInterface {
	return roleRepository{db: db}
}

type roleRepository struct {
	db *gorm.DB
}

func (r roleRepository) GetByID(id uint) (role entities.Role, err error) {
	err = r.db.First(&role, id).Error
	return
}
