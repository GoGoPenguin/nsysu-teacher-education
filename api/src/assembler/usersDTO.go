package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// UsersDTO users data transfer object
func UsersDTO(users *[]gorm.User) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, user := range *users {
		result = append(result, map[string]interface{}{
			"Name":      user.Name,
			"Account":   user.Account,
			"CreatedAt": user.CreatedAt,
		})
	}
	return result
}
