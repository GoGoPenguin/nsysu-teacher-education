package specification

import (
	"gorm.io/gorm"
)

// PaginationSpecification generate pagination query
func PaginationSpecification(start, length int) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(start).Limit(length)
	}
}
