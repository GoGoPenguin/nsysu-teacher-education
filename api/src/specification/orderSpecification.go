package specification

import (
	"github.com/jinzhu/gorm"
)

// OrderSpecification generate limt query
func OrderSpecification(column, direction string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(column + " " + direction)
	}
}
