package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// LeturesDTO leture data transfer object
func LeturesDTO(letures *[]gorm.Leture) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, leture := range *letures {
		result = append(result, map[string]interface{}{
			"ID":        leture.ID,
			"Name":      leture.Name,
			"MinCredit": leture.MinCredit,
			"Comment":   leture.Comment,
			"Status":    leture.Status,
		})
	}
	return result
}
