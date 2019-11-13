package gorm

import (
	"github.com/jinzhu/gorm"
)

// Student model
type Student struct {
	gorm.Model
	Name     string `gorm:"column:name;"`
	Account  string `gorm:"column:account; unique_index"`
	Password string `gorm:"column:password;"`
	Major    string `gorm:"column:major;"`
	Number   string `gorm:"column:number"`
}

type studentDao struct {
	table string
	Role  string
}

// StudentDao user data access object
var StudentDao = studentDao{
	table: "student",
	Role:  "student",
}

// New a record
func (dao *studentDao) New(tx *gorm.DB, student *Student) {
	err := tx.Table(dao.table).
		Create(student).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *studentDao) GetByID(tx *gorm.DB, id uint) *Student {
	result := Student{}
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

// GetByAccount get a record by account
func (dao *studentDao) GetByAccount(tx *gorm.DB, account string) *Student {
	result := Student{}
	err := tx.Table(dao.table).
		Where("account = ?", account).
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

// Count get total count
func (dao *studentDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int {
	var count int
	tx.Table(dao.table).
		Scopes(funcs...).
		Count(&count)

	return count
}

// Query custom query
func (dao *studentDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]Student {
	var result []Student
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
