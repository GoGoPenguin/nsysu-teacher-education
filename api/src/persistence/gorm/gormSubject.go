package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// Subject model
type Subject struct {
	gorm.Model
	SubjectGroupID uint           `gorm:"column:subject_group_id;"`
	Name           string         `gorm:"column:name;"`
	Credit         uint           `gorm:"column:credit;"`
	Compulsory     bool           `gorm:"column:compulsory;"`
	Status         string         `gorm:"column:status; default:'enable'"`
	StudentSubject StudentSubject `gorm:"foreignkey:SubjectID; PRELOAD:false"`
}

type subjectDao struct {
	table   string
	Enable  string
	Disable string
}

// SubjectDao subject data access object
var SubjectDao = &subjectDao{
	table:   "subject",
	Enable:  "enable",
	Disable: "disable",
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
		First(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// GetByIDAndType get a record by name and group
func (dao *subjectDao) GetByNameAndGroup(tx *gorm.DB, name string, groupID uint) *Subject {
	result := Subject{}
	err := tx.Table(dao.table).
		Where("name = ?", name).
		Where("subject_group_id = ?", groupID).
		Where("deleted_at IS NULL").
		First(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
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

	if err != nil {
		panic(err)
	}
	return &result
}
