package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// LectureType model
type LectureType struct {
	gorm.Model
	LectureCategoryID uint           `gorm:"column:lecture_category_id;"`
	Name              string         `gorm:"column:name;"`
	MinCredit         uint           `gorm:"column:min_credit;"`
	Groups            []SubjectGroup `gorm:"foreignkey:LectureTypeID"`
}

type lectureTypeDao struct {
	table string
}

// LectureTypeDao lecture_type data access object
var LectureTypeDao = &lectureTypeDao{
	table: "lecture_type",
}

// New a record
func (dao *lectureTypeDao) New(tx *gorm.DB, lectureType *LectureType) {
	err := tx.Table(dao.table).
		Create(lectureType).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *lectureTypeDao) GetByID(tx *gorm.DB, id uint) *LectureType {
	result := LectureType{}
	err := tx.Table(dao.table).
		Where("id = ?", id).
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

// GetByAccount get a record by category and name
func (dao *lectureTypeDao) GetByCategoryAndName(tx *gorm.DB, categoryID uint, name string) *LectureType {
	result := LectureType{}
	err := tx.Table(dao.table).
		Where("name = ?", name).
		Where("lecture_category_id = ?", categoryID).
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
func (dao *lectureTypeDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]LectureType {
	var result []LectureType
	err := tx.Table(dao.table).
		Scopes(funcs...).
		Scan(&result).Error

	if err != nil {
		panic(err)
	}
	return &result
}
