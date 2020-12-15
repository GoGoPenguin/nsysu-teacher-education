package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentServiceLearningsDTO student-service-learning data transfer object
func StudentServiceLearningsDTO(studentServiceLearnings *[]gorm.StudentServiceLearning) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, studentServiceLearning := range *studentServiceLearnings {
		if studentServiceLearning.ServiceLearning.ID != 0 && studentServiceLearning.Student.ID != 0 {
			result = append(result, map[string]interface{}{
				"ID": studentServiceLearning.ID,
				"Student": map[string]interface{}{
					"Name":    studentServiceLearning.Student.Name,
					"Account": studentServiceLearning.Student.Account,
					"Major":   studentServiceLearning.Student.Major,
					"Number":  studentServiceLearning.Student.Number,
				},
				"ServiceLearning": map[string]interface{}{
					"ID":      studentServiceLearning.ServiceLearning.ID,
					"Type":    studentServiceLearning.ServiceLearning.Type,
					"Content": studentServiceLearning.ServiceLearning.Content,
					"Session": studentServiceLearning.ServiceLearning.Session,
					"Hours":   studentServiceLearning.ServiceLearning.Hours,
					"Start":   studentServiceLearning.ServiceLearning.Start,
					"End":     studentServiceLearning.ServiceLearning.End,
				},
				"Status":    studentServiceLearning.Status,
				"Review":    studentServiceLearning.Review,
				"Reference": studentServiceLearning.Reference,
				"Hours":     studentServiceLearning.Hours,
				"Comment":   studentServiceLearning.Comment,
			})
		}
	}
	return result
}
