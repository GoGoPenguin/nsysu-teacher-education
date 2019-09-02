package specification

import (
	"github.com/jinzhu/gorm"
)

// RoleSpecification generate roel query
func RoleSpecification(role string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("role = ?", role)
	}
}
