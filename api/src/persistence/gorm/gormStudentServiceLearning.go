package gorm

import (
	"errors"

	"gorm.io/gorm"
)

// StudentServiceLearning model
type StudentServiceLearning struct {
	gorm.Model
	StudentID         uint            `gorm:"column:student_id;"`
	Student           Student         `gorm:"foreignkey:StudentID;"`
	ServiceLearningID uint            `gorm:"column:service_learning_id;"`
	ServiceLearning   ServiceLearning `gorm:"foreignkey:ServiceLearningID;"`
	Status            string          `gorm:"column:status"`
	Review            string          `gorm:"column:review"`
	Reference         string          `gorm:"column:reference"`
	Hours             *uint           `gorm:"column:hours"`
	Comment           string          `gorm:"column:comment"`
}

type studentServiceLearningDao struct {
	table        string
	StatusPass   string
	StatusFailed string
}

// StudentServiceLearningDao student service-learning data access object
var StudentServiceLearningDao = &studentServiceLearningDao{
	table:        "student_service_learning",
	StatusPass:   "pass",
	StatusFailed: "failed",
}

// New a record
func (dao *studentServiceLearningDao) New(tx *gorm.DB, user *StudentServiceLearning) {
	err := tx.Table(dao.table).
		Create(user).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *studentServiceLearningDao) GetByID(tx *gorm.DB, id uint) *StudentServiceLearning {
	result := StudentServiceLearning{}
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

// Update record
func (dao *studentServiceLearningDao) Update(tx *gorm.DB, studentServiceLearning *StudentServiceLearning) {
	err := tx.Model(&studentServiceLearning).
		Updates(map[string]interface{}{
			"StudentID":         studentServiceLearning.StudentID,
			"ServiceLearningID": studentServiceLearning.ServiceLearningID,
			"Status":            studentServiceLearning.Status,
			"Review":            studentServiceLearning.Review,
			"Reference":         studentServiceLearning.Reference,
			"Hours":             studentServiceLearning.Hours,
			"Comment":           studentServiceLearning.Comment,
		}).Error

	if err != nil {
		panic(err)
	}
}

// Count get total count
func (dao *studentServiceLearningDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int64 {
	var result []StudentServiceLearning
	count := tx.
		Joins("Student").
		Joins("ServiceLearning").
		Table(dao.table).
		Scopes(funcs...).
		Find(&result).RowsAffected

	return count
}

// Query custom query
func (dao *studentServiceLearningDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]StudentServiceLearning {
	var result []StudentServiceLearning
	err := tx.
		Joins("Student").
		Joins("ServiceLearning").
		Table(dao.table).
		Scopes(funcs...).
		Find(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}
