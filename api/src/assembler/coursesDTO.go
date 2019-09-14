package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// CoursesDTO users data transfer object
func CoursesDTO(courses *[]gorm.Course) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, course := range *courses {
		result = append(result, map[string]interface{}{
			"ID":          course.ID,
			"Topic":       course.Topic,
			"Information": course.Information,
			"Type":        course.Type,
			"Start":       course.Start.String(),
			"End":         course.End.String(),
		})
	}
	return result
}
