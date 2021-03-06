package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// Lecture model
type Lecture struct {
	gorm.Model
	Name       string            `gorm:"column:name;"`
	MinCredit  uint              `gorm:"column:min_credit;"`
	Comment    string            `gorm:"column:comment;"`
	Status     string            `gorm:"column:status; default:'enable'"`
	Categories []LectureCategory `gorm:"foreignkey:LectureID"`
}

type lectureDao struct {
	table   string
	Enable  string
	Disable string
}

// LectureDao lecture data access object
var LectureDao = &lectureDao{
	table:   "lecture",
	Enable:  "enable",
	Disable: "disable",
}

// New a record
func (dao *lectureDao) New(tx *gorm.DB, lecture *Lecture) {
	err := tx.Table(dao.table).
		Create(lecture).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *lectureDao) GetByID(tx *gorm.DB, id uint) *Lecture {
	result := Lecture{
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

// GetByAccount get a record by name
func (dao *lectureDao) GetByName(tx *gorm.DB, name string) *Lecture {
	result := Lecture{}
	err := tx.Table(dao.table).
		Where("name = ?", name).
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

// Count get total count
func (dao *lectureDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int64 {
	var result []Lecture
	count := tx.Table(dao.table).
		Scopes(funcs...).
		Find(&result).RowsAffected

	return count
}

// Query custom query
func (dao *lectureDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]Lecture {
	var result []Lecture
	err := tx.Table(dao.table).
		Scopes(funcs...).
		Find(&result).Error

	if err != nil {
		panic(err)
	}
	return &result
}
