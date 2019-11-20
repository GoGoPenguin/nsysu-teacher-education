package gorm

import (
	"github.com/jinzhu/gorm"
)

// StudentLeture model
type StudentLeture struct {
	gorm.Model
	StudentID uint    `gorm:"column:student_id;"`
	Student   Student `gorm:"foreignkey:StudentID;"`
	LetureID  uint    `gorm:"column:leture_id;"`
	Leture    Leture  `gorm:"foreignkey:LetureID;"`
	Pass      bool    `gorm:"column:pass;"`
}

type studentLetureDao struct {
	table        string
	Meat         string
	Vegetable    string
	StatusPass   string
	StatusFailed string
}

// StudentLetureDao student leture data access object
var StudentLetureDao = &studentLetureDao{
	table: "student_leture",
}

// New a record
func (dao *studentLetureDao) New(tx *gorm.DB, studentLeture *StudentLeture) {
	err := tx.Table(dao.table).
		Create(studentLeture).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *studentLetureDao) GetByID(tx *gorm.DB, id uint) *StudentLeture {
	result := StudentLeture{
		Model: gorm.Model{
			ID: id,
		},
	}
	err := tx.Table(dao.table).
		Find(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// Update record
func (dao *studentLetureDao) Update(tx *gorm.DB, studentLeture *StudentLeture) {
	err := tx.Model(&studentLeture).
		Updates(map[string]interface{}{
			"StudentID": studentLeture.StudentID,
			"LetureID":  studentLeture.LetureID,
			"Meal":      studentLeture.Pass,
		}).Error

	if gorm.IsRecordNotFoundError(err) {
		return
	}
	if err != nil {
		panic(err)
	}
}

// Count get total count
func (dao *studentLetureDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int {
	var count int
	tx.Table(dao.table).
		Scopes(funcs...).
		Count(&count)

	return count
}

// Query custom query
func (dao *studentLetureDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]StudentLeture {
	var result []StudentLeture
	err := tx.Table(dao.table).
		Select("*").
		Scopes(funcs...).
		Find(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}
