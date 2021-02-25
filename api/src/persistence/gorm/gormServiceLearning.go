package gorm

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// ServiceLearning model
type ServiceLearning struct {
	gorm.Model
	Type      string    `gorm:"column:type;"`
	Content   string    `gorm:"column:content;"`
	Session   string    `gorm:"column:session;"`
	Hours     uint      `gorm:"column:hours"`
	Start     time.Time `gorm:"column:start"`
	End       time.Time `gorm:"column:end"`
	Show      *bool     `gorm:"column:show;"`
	CreatedBy *uint     `gorm:"column:created_by;"`
	Student   Student   `gorm:"foreignkey:CreatedBy;"`
}

type serviceLearningDao struct {
	table          string
	TypeInternship string
	TypeVolunteer  string
	TypeBoth       string
}

// ServiceLearningDao service-learning data access object
var ServiceLearningDao = &serviceLearningDao{
	table:          "service_learning",
	TypeInternship: "internship",
	TypeVolunteer:  "volunteer",
	TypeBoth:       "both",
}

// New a record
func (dao *serviceLearningDao) New(tx *gorm.DB, serviceLearning *ServiceLearning) {
	err := tx.Table(dao.table).
		Create(serviceLearning).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *serviceLearningDao) GetByID(tx *gorm.DB, id uint) *ServiceLearning {
	result := ServiceLearning{}
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

// Count get total count
func (dao *serviceLearningDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int64 {
	var result []ServiceLearning
	count := tx.Table(dao.table).
		Scopes(funcs...).
		Find(&result).RowsAffected

	return count
}

// Query custom query
func (dao *serviceLearningDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]ServiceLearning {
	var result []ServiceLearning
	err := tx.Table(dao.table).
		Scopes(funcs...).
		Find(&result).Error

	if err != nil {
		panic(err)
	}
	return &result
}

// Delete a service-learning
func (dao *serviceLearningDao) Delete(tx *gorm.DB, id uint) {
	attrs := map[string]interface{}{
		"deleted_at": time.Now(),
	}
	err := tx.Table(dao.table).
		Where("id = ?", id).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

// Update a record
func (dao *serviceLearningDao) Update(tx *gorm.DB, serviceLearning *ServiceLearning) {
	attrs := map[string]interface{}{
		"Type":    serviceLearning.Type,
		"Content": serviceLearning.Content,
		"Session": serviceLearning.Session,
		"Hours":   serviceLearning.Hours,
		"Start":   serviceLearning.Start,
		"End":     serviceLearning.End,
		"Show":    serviceLearning.Show,
	}
	err := tx.Model(serviceLearning).
		Where("id = ?", serviceLearning.ID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}
