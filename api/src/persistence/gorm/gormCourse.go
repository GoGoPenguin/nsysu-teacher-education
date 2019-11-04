package gorm

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Course model
type Course struct {
	gorm.Model
	Topic       string    `gorm:"column:topic;"`
	Information string    `gorm:"column:information;"`
	Type        string    `gorm:"column:type;"`
	Start       time.Time `gorm:"column:start"`
	End         time.Time `gorm:"column:end"`
}

type courseDao struct {
	table string
	TypeA string
	TypeB string
	TypeC string
}

// CourseDao user data acces object
var CourseDao = &courseDao{
	table: "course",
	TypeA: "A",
	TypeB: "B",
	TypeC: "C",
}

// New a record
func (dao *courseDao) New(tx *gorm.DB, course *Course) {
	err := tx.Table(dao.table).
		Create(course).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *courseDao) GetByID(tx *gorm.DB, id uint) *Course {
	result := Course{}
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

// GetByTopic get a record by topic
func (dao *courseDao) GetByTopic(tx *gorm.DB, topic string) *Course {
	result := Course{}
	err := tx.Table(dao.table).
		Where("topic = ?", topic).
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

// GetByInformation get a record by information
func (dao *courseDao) GetByInformation(tx *gorm.DB, information string) *Course {
	result := Course{}
	err := tx.Table(dao.table).
		Where("information = ?", information).
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
func (dao *courseDao) Count(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) int {
	var count int
	tx.Table(dao.table).
		Scopes(funcs...).
		Count(&count)

	return count
}

// Query custom query
func (dao *courseDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]Course {
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

// Delete a course
func (dao *courseDao) Delete(tx *gorm.DB, id uint) {
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
