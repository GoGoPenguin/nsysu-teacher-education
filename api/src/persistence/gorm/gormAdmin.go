package gorm

import (
	"github.com/jinzhu/gorm"
)

// Admin model
type Admin struct {
	gorm.Model
	Account  string `gorm:"column:account; unique_index"`
	Password string `gorm:"column:password;"`
}

type adminDao struct {
	table string
	Role  string
}

// AdminDao user data acces object
var AdminDao = &adminDao{
	table: "admin",
	Role:  "admin",
}

// New a record
func (dao *adminDao) New(tx *gorm.DB, user *Admin) {
	err := tx.Table(dao.table).
		Create(user).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *adminDao) GetByID(tx *gorm.DB, id uint) *Admin {
	result := Admin{}
	err := tx.Table(dao.table).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// GetByAccount get a record by id
func (dao *adminDao) GetByAccount(tx *gorm.DB, acount string) *Admin {
	result := Admin{}
	err := tx.Table(dao.table).
		Where("account = ?", acount).
		Where("deleted_at IS NULL").
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// Query custom query
func (dao *adminDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]Admin {
	var result []Admin
	err := tx.Table(dao.table).
		Scopes(funcs...).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}