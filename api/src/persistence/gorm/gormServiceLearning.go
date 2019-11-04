package gorm

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceLearning model
type ServiceLearning struct {
	gorm.Model
	Type    string    `gorm:"column:type;"`
	Content string    `gorm:"column:content;"`
	Session string    `gorm:"column:session;"`
	Hours   uint      `gormn:"column:hours"`
	Start   time.Time `gorm:"column:start"`
	End     time.Time `gorm:"column:end"`
}

type serviceLearningDao struct {
	table          string
	TypeInternship string
	TypeVolunteer  string
	TypeBoth       string
}

// ServiceLearningDao user data acces object
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
func (dao *serviceLearningDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int {
	var count int
	tx.Table(dao.table).
		Scopes(funcs...).
		Count(&count)

	return count
}

// Query custom query
func (dao *serviceLearningDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]ServiceLearning {
	var result []ServiceLearning
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
