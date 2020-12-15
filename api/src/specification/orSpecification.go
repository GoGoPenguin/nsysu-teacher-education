package specification

import (
	"fmt"

	"gorm.io/gorm"
)

// OrSpecification generate or query
func OrSpecification(condition1, condition2 string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(fmt.Sprintf("(%s OR %s)", condition1, condition2))
	}
}
