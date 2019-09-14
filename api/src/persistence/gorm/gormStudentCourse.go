package gorm

import (
	"github.com/jinzhu/gorm"
)

// StudentCourse model
type StudentCourse struct {
	gorm.Model
	StudentID uint   `gorm:"column:student_id;"`
	CourseID  uint   `gorm:"column:course_id;"`
	Meal      string `gorm:"column:meal;"`
	Status    string `gorm:"column:status"`
	Review    string `gorm:"column:review"`
	Comment   string `gorm:"column:comment"`
}

type studentCourseDao struct {
	table        string
	Meat         string
	Vegetable    string
	StatusPass   string
	StatusFailed string
}

// StudentCourseDao user data acces object
var StudentCourseDao = &studentCourseDao{
	table:        "student_course",
	Meat:         "meate",
	Vegetable:    "vegetable",
	StatusPass:   "pass",
	StatusFailed: "failed",
}

// New a record
func (dao *studentCourseDao) New(tx *gorm.DB, user *StudentCourse) {
	err := tx.Table(dao.table).
		Create(user).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *studentCourseDao) GetByID(tx *gorm.DB, id uint) *StudentCourse {
	result := StudentCourse{}
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

// Count get total count
func (dao *studentCourseDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int {
	var count int
	tx.Table(dao.table).
		Scopes(funcs...).
		Count(&count)

	return count
}

// Query custom query
func (dao *studentCourseDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]Course {
	var result []Course
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
