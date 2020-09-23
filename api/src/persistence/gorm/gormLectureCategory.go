package gorm

import (
	"github.com/jinzhu/gorm"
)

// LectureCategory model
type LectureCategory struct {
	gorm.Model
	LectureID uint          `gorm:"column:lecture_id;"`
	Name      string        `gorm:"column:name;"`
	MinCredit uint          `gorm:"column:min_credit;"`
	MinType   uint          `gorm:"column:min_type;"`
	Types     []LectureType `gorm:"foreignkey:LectureCategoryID"`
}

type lectureCategoryDao struct {
	table string
}

// LectureCategoryDao lecture_category data access object
var LectureCategoryDao = &lectureCategoryDao{
	table: "lecture_category",
}

// New a record
func (dao *lectureCategoryDao) New(tx *gorm.DB, lectureCategory *LectureCategory) {
	err := tx.Table(dao.table).
		Create(lectureCategory).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *lectureCategoryDao) GetByID(tx *gorm.DB, id uint) *LectureCategory {
	result := LectureCategory{}
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

// GetByAccount get a record by lecture and name
func (dao *lectureCategoryDao) GetByLectureAndName(tx *gorm.DB, lectureID uint, name string) *LectureCategory {
	result := LectureCategory{}
	err := tx.Table(dao.table).
		Where("name = ?", name).
		Where("lecture_id = ?", lectureID).
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
func (dao *lectureCategoryDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]LectureCategory {
	var result []LectureCategory
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
