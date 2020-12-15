package specification

import (
	"gorm.io/gorm"
)

// StatusSpecification generate is null query
func StatusSpecification(status string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("status = ?", status)
	}
}
