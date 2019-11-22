package gorm

import (
	"github.com/jinzhu/gorm"
)

// StudentSubject model
type StudentSubject struct {
	gorm.Model
	StudentLetureID uint   `gorm:"column:student_leture_id;"`
	SubjectID       uint   `gorm:"column:subject_id;"`
	Name            string `gorm:"column:name;"`
	Year            string `gorm:"column:year;"`
	Semester        string `gorm:"column:semester;"`
	Credit          string `gorm:"column:credit;"`
	Score           string `gorm:"column:score;"`
}

type studentSubjectDao struct {
	table string
}

// StudentSubjectDao student subject data access object
var StudentSubjectDao = &studentSubjectDao{
	table: "student_subject",
}

// New a record
func (dao *studentSubjectDao) New(tx *gorm.DB, studentSubject *StudentSubject) {
	err := tx.Table(dao.table).
		Create(studentSubject).Error

	if err != nil {
		panic(err)
	}
}

// GetByID get a record by id
func (dao *studentSubjectDao) GetByID(tx *gorm.DB, id uint) *StudentSubject {
	result := StudentSubject{
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

// GetByID get a record by id
func (dao *studentSubjectDao) GetByLetureAndSubject(tx *gorm.DB, letureID, subjectID uint) *StudentSubject {
	var result StudentSubject

	err := tx.Table(dao.table).
		Where(&StudentSubject{
			StudentLetureID: letureID,
			SubjectID:       subjectID,
		}).
		Find(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return &result
}

// Query custom query
func (dao *studentSubjectDao) Query(tx *gorm.DB, funcs ...func(*gorm.DB) *gorm.DB) *[]StudentSubject {
	var result []StudentSubject
	err := tx.Table(dao.table).
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

// Update a record
func (dao *studentSubjectDao) Update(tx *gorm.DB, studentSubject *StudentSubject) {
	attrs := map[string]interface{}{
		"Name":     studentSubject.Name,
		"Year":     studentSubject.Year,
		"Semester": studentSubject.Semester,
		"Credit":   studentSubject.Credit,
		"Score":    studentSubject.Score,
	}
	err := tx.Model(studentSubject).
		Where("id = ?", studentSubject.ID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}
