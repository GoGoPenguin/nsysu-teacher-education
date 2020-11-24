package specification

import (
	"gorm.io/gorm"
)

// LimitSpecification generate limit query
func LimitSpecification(count int) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(count)
	}
}
