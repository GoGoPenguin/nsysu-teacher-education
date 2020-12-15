package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// SubjectGroup model
type SubjectGroup struct {
	gorm.Model
	LectureTypeID uint      `gorm:"column:lecture_type_id;"`
	MinCredit     uint      `gorm:"column:min_credit;"`
	Subjects      []Subject `gorm:"foreignkey:SubjectGroupID"`
}

type subjectGroupDao struct {
	table string
}

// SubjectGroupDao subject_group data access object
var SubjectGroupDao = &subjectGroupDao{
	table: "subject_group",
}

// New a record
func (dao *subjectGroupDao) New(tx *gorm.DB, subjectGroup *SubjectGroup) {
	err := tx.Table(dao.table).
		Create(subjectGroup).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *subjectGroupDao) GetByID(tx *gorm.DB, id uint) *SubjectGroup {
	result := SubjectGroup{
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

// GetByIDAndType get a record by type and id
func (dao *subjectGroupDao) GetByIDAndType(tx *gorm.DB, id, typeID uint) *SubjectGroup {
	result := SubjectGroup{}
	err := tx.Table(dao.table).
		Where("id = ?", id).
		Where("lecture_type_id = ?", typeID).
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
func (dao *subjectGroupDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]SubjectGroup {
	var result []SubjectGroup
	err := tx.Table(dao.table).
		Scopes(funcs...).
		Scan(&result).Error

	if err != nil {
		panic(err)
	}
	return &result
}
