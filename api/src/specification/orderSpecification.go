package specification

import (
	"gorm.io/gorm"
)

// OrderSpecification generate order query
func OrderSpecification(column, direction string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(column + " " + direction)
	}
}
