package gorm

import (
	"github.com/jinzhu/gorm"
)

// Subject model
type Subject struct {
	gorm.Model
	SubjectGroupID uint   `gorm:"column:subject_group_id;"`
	Name           string `gorm:"column:name;"`
	Credit         uint   `gorm:"column:credit;"`
	Compulsory     bool   `grom:"column:compulsory;"`
	Status         string `gorm:"column:status;"`
}

type subjectDao struct {
	table string
}

// SubjectDao subject data acces object
var SubjectDao = &subjectDao{
	table: "subject",
}

// New a record
func (dao *subjectDao) New(tx *gorm.DB, subject *Subject) {
	err := tx.Table(dao.table).
		Create(subject).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *subjectDao) GetByID(tx *gorm.DB, id uint) *Subject {
	result := Subject{
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

// GetByIDAndType get a record by name
func (dao *subjectDao) GetByName(tx *gorm.DB, name string) *Subject {
	result := Subject{}
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
func (dao *subjectDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]Subject {
	var result []Subject
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
