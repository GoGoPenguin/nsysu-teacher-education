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
	Categories []LetureCategory `gorm:"foreignkey:LetureID"`
}

type letureDao struct {
	table string
}

// LetureDao leture data acces object
var LetureDao = &letureDao{
	table: "leture",
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
		Where("deleted_at IS NULL").
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
