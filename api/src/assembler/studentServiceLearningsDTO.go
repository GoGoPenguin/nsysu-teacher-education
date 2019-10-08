package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentServiceLearningsDTO student-service-learning data transfer object
func StudentServiceLearningsDTO(studentServiceLearnings *[]gorm.StudentServiceLearning) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, studentServiceLearning := range *studentServiceLearnings {
		result = append(result, map[string]interface{}{
			"ID": studentServiceLearning.ID,
			"Student": map[string]interface{}{
				"Name":    studentServiceLearning.Student.Name,
				"Account": studentServiceLearning.Student.Account,
				"Major":   studentServiceLearning.Student.Major,
				"Number":  studentServiceLearning.Student.Number,
			},
			"ServiceLearning": map[string]interface{}{
				"Type":    studentServiceLearning.ServiceLearning.Type,
				"Content": studentServiceLearning.ServiceLearning.Content,
				"Session": studentServiceLearning.ServiceLearning.Session,
				"Hours":   studentServiceLearning.ServiceLearning.Hours,
				"Start":   studentServiceLearning.ServiceLearning.Start.String(),
				"End":     studentServiceLearning.ServiceLearning.End.String(),
			},
			"Status":    studentServiceLearning.Status,
			"Review":    studentServiceLearning.Review,
			"Reference": studentServiceLearning.Reference,
			"Comment":   studentServiceLearning.Comment,
		})
	}
	return result
}