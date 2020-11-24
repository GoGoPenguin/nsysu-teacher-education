package specification

import (
	"gorm.io/gorm"
)

// IDSpecification generate id query
func IDSpecification(id string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id = ?", id)
	}
}
