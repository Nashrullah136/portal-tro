package db

import "gorm.io/gorm"

type Transactor interface {
	Commit() *gorm.DB
	Rollback() *gorm.DB
	GetDB() *gorm.DB
}

func NewTransactor(db *gorm.DB) Transactor {
	return transaction{db: db}
}

type transaction struct {
	db *gorm.DB
}

func (t transaction) Commit() *gorm.DB {
	return t.db.Commit()
}

func (t transaction) Rollback() *gorm.DB {
	return t.db.Rollback()
}

func (t transaction) GetDB() *gorm.DB {
	return t.db
}
