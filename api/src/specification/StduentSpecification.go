package specification

import (
	"github.com/jinzhu/gorm"
)

// StudentSpecification generate is null query
func StudentSpecification(id uint) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("student_id = ?", id)
	}
}
