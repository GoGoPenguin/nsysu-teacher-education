package gorm

import (
	"github.com/jinzhu/gorm"
)

// LetureType model
type LetureType struct {
	gorm.Model
	LetureCategoryID uint           `gorm:"column:leture_category_id;"`
	Name             string         `gorm:"column:name;"`
	MinCredit        uint           `gorm:"column:min_credit;"`
	Groups           []SubjectGroup `gorm:"foreignkey:LetureTypeID"`
}

type letureTypeDao struct {
	table string
}

// LetureTypeDao leture_type data acces object
var LetureTypeDao = &letureTypeDao{
	table: "leture_type",
}

// New a record
func (dao *letureTypeDao) New(tx *gorm.DB, letureType *LetureType) {
	err := tx.Table(dao.table).
		Create(letureType).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *letureTypeDao) GetByID(tx *gorm.DB, id uint) *LetureType {
	result := LetureType{}
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

// GetByAccount get a record by category and name
func (dao *letureTypeDao) GetByCategoryAndName(tx *gorm.DB, categoryID uint, name string) *LetureType {
	result := LetureType{}
	err := tx.Table(dao.table).
		Where("name = ?", name).
		Where("leture_category_id = ?", categoryID).
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
func (dao *letureTypeDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]LetureType {
	var result []LetureType
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
