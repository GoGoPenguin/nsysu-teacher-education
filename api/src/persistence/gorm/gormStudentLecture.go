package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// StudentLecture model
type StudentLecture struct {
	gorm.Model
	StudentID uint    `gorm:"column:student_id;"`
	Student   Student `gorm:"foreignkey:StudentID;"`
	LectureID uint    `gorm:"column:lecture_id;"`
	Lecture   Lecture `gorm:"foreignkey:LectureID;"`
	Pass      bool    `gorm:"column:pass;"`
}

type studentLectureDao struct {
	table        string
	Meat         string
	Vegetable    string
	StatusPass   string
	StatusFailed string
}

// StudentLectureDao student lecture data access object
var StudentLectureDao = &studentLectureDao{
	table: "student_lecture",
}

// New a record
func (dao *studentLectureDao) New(tx *gorm.DB, studentLecture *StudentLecture) {
	err := tx.Table(dao.table).
		Create(studentLecture).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *studentLectureDao) GetByID(tx *gorm.DB, id uint) *StudentLecture {
	result := StudentLecture{
		Model: gorm.Model{
			ID: id,
		},
	}
	err := tx.Table(dao.table).
		First(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// GetByID get a record by id
func (dao *studentLectureDao) GetByLectureAndStudent(tx *gorm.DB, lectureID, studentID uint) *StudentLecture {
	var result StudentLecture

	err := tx.Table(dao.table).
		Where(&StudentLecture{
			StudentID: studentID,
			LectureID: lectureID,
		}).
		First(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// Update record
func (dao *studentLectureDao) Update(tx *gorm.DB, studentLecture *StudentLecture) {
	err := tx.Model(&studentLecture).
		Updates(map[string]interface{}{
			"Pass": studentLecture.Pass,
		}).Error

	if err != nil {
		panic(err)
	}
}

// Count get total count
func (dao *studentLectureDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int64 {
	var result []StudentLecture
	count := tx.Table(dao.table).
		Select("*").
		Scopes(funcs...).
		Find(&result).RowsAffected

	return count
}

// Query custom query
func (dao *studentLectureDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]StudentLecture {
	var result []StudentLecture
	err := tx.Table(dao.table).
		Select("*").
		Scopes(funcs...).
		Find(&result).Error

	if err != nil {
		panic(err)
	}
	return &result
}
