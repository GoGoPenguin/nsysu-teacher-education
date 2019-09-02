package specification

import (
	"github.com/jinzhu/gorm"
)

// PaginationSpecification generate pagination query
func PaginationSpecification(idIndex, orderDirection string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if orderDirection == OrderDirectionASC {
			return tx.Where(""+IDColumn+" > ?", idIndex).Order(IDColumn + " " + orderDirection)
		}
		return tx.Where(""+IDColumn+" < ?", idIndex).Order(IDColumn + " " + orderDirection)
	}
}
