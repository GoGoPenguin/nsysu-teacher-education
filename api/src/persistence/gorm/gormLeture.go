package gorm

import (
	"github.com/jinzhu/gorm"
)

// Leture model
type Leture struct {
	gorm.Model
	Name       string           `gorm:"column:name;"`
	MinCredit  uint             `gorm:"column:min_credit;"`
	Comment    string           `gorm:"column:comment;"`
	Status     string           `gorm:"column:status;"`
	Categories []LetureCategory `gorm:"foreignkey:LetureID"`
}

type letureDao struct {
	table   string
	Enable  string
	Disable string
}

// LetureDao leture data acces object
var LetureDao = &letureDao{
	table:   "leture",
	Enable:  "enable",
	Disable: "disable",
}

// New a record
func (dao *letureDao) New(tx *gorm.DB, leture *Leture) {
	err := tx.Table(dao.table).
		Create(leture).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *letureDao) GetByID(tx *gorm.DB, id uint) *Leture {
	result := Leture{
		Model: gorm.Model{
			ID: id,
		},
	}
	err := tx.Table(dao.table).
		Find(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// GetByAccount get a record by name
func (dao *letureDao) GetByName(tx *gorm.DB, name string) *Leture {
	result := Leture{}
	err := tx.Table(dao.table).
		Where("name = ?", name).
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
func (dao *letureDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int {
	var count int
	tx.Table(dao.table).
		Scopes(funcs...).
		Count(&count)

	return count
}

// Query custom query
func (dao *letureDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]Leture {
	var result []Leture
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
