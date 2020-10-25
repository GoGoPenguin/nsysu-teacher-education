package specification

import (
	"github.com/jinzhu/gorm"
)

// PreloadSpecification generate preload query
func PreloadSpecification(column string, conditions ...interface{}) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Preload(column, conditions...)
	}
}
