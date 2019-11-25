package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentsDTO studnet data transfer object
func StudentsDTO(students *[]gorm.Student) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, student := range *students {
		result = append(result, map[string]interface{}{
			"Name":      student.Name,
			"Account":   student.Account,
			"Major":     student.Major,
			"Number":    student.Number,
			"CreatedAt": student.CreatedAt,
		})
	}
	return result
}
