package gorm

import (
	"github.com/jinzhu/gorm"
)

// LetureCategory model
type LetureCategory struct {
	gorm.Model
	LetureID  uint         `gorm:"column:leture_id;"`
	Name      string       `gorm:"column:name;"`
	MinCredit uint         `gorm:"column:min_credit;"`
	MinType   uint         `gorm:"column:min_type;"`
	Types     []LetureType `gorm:"foreignkey:LetureCategoryID"`
}

type letureCategoryDao struct {
	table string
}

// LetureCategoryDao leture_category data acces object
var LetureCategoryDao = &letureCategoryDao{
	table: "leture_category",
}

// New a record
func (dao *letureCategoryDao) New(tx *gorm.DB, letureCategory *LetureCategory) {
	err := tx.Table(dao.table).
		Create(letureCategory).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *letureCategoryDao) GetByID(tx *gorm.DB, id uint) *LetureCategory {
	result := LetureCategory{}
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

// GetByAccount get a record by leture and name
func (dao *letureCategoryDao) GetByLetureAndName(tx *gorm.DB, letureID uint, name string) *LetureCategory {
	result := LetureCategory{}
	err := tx.Table(dao.table).
		Where("name = ?", name).
		Where("leture_id = ?", letureID).
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
func (dao *letureCategoryDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]LetureCategory {
	var result []LetureCategory
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
