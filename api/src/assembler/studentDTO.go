package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentDTO studnet data transfer object
func StudentDTO(student *gorm.Student) map[string]interface{} {
	return map[string]interface{}{
		"Name":      student.Name,
		"Account":   student.Account,
		"Major":     student.Major,
		"Number":    student.Number,
		"CreatedAt": student.CreatedAt,
	}
}
