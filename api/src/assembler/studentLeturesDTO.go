package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentLeturesDTO student-leture data transfer object
func StudentLeturesDTO(studentLetures *[]gorm.StudentLeture) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, studentLeture := range *studentLetures {
		temp := map[string]interface{}{
			"ID": studentLeture.ID,

			"Leture": map[string]interface{}{
				"ID":        studentLeture.Leture.ID,
				"Name":      studentLeture.Leture.Name,
				"MinCredit": studentLeture.Leture.MinCredit,
				"Comment":   studentLeture.Leture.Comment,
				"Status":    studentLeture.Leture.Status,
			},
			"Pass": studentLeture.Pass,
		}
		if studentLeture.Student.ID != 0 {
			temp["Student"] = map[string]interface{}{
				"ID":      studentLeture.Student.ID,
				"Name":    studentLeture.Student.Name,
				"Account": studentLeture.Student.Account,
				"Major":   studentLeture.Student.Major,
				"Number":  studentLeture.Student.Number,
			}
		}
		result = append(result, temp)
	}
	return result
}
