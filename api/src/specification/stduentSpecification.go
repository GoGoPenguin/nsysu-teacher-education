package specification

import (
	"gorm.io/gorm"
)

// StudentSpecification generate student id query
func StudentSpecification(id uint) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("student_id = ?", id)
	}
}
