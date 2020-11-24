package specification

import (
	"gorm.io/gorm"
)

// IsNullSpecification generate is null query
func IsNullSpecification(column string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(column + " IS NULL")
	}
}
