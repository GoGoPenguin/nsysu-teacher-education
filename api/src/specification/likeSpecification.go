package specification

import (
	"gorm.io/gorm"
)

// LikeSpecification generate is null query
func LikeSpecification(columns []string, value string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		var sql string
		for _, column := range columns {
			sql += column + " LIKE '%" + value + "%' OR "
		}
		return tx.Where(sql[0 : len(sql)-4])
	}
}
