package specification

import (
	"gorm.io/gorm"
)

// BiggerSpecification generate bigger query
func BiggerSpecification(column, value string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(column+" > ?", value)
	}
}
