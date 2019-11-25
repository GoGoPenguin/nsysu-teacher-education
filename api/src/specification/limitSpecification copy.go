package specification

import (
	"github.com/jinzhu/gorm"
)

// LimitSpecification generate limit query
func LimitSpecification(count string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(count)
	}
}
