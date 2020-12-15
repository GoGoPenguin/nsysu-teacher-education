package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// LecturesDTO lecture data transfer object
func LecturesDTO(lectures *[]gorm.Lecture) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, lecture := range *lectures {
		result = append(result, map[string]interface{}{
			"ID":        lecture.ID,
			"Name":      lecture.Name,
			"MinCredit": lecture.MinCredit,
			"Comment":   lecture.Comment,
			"Status":    lecture.Status,
		})
	}
	return result
}
