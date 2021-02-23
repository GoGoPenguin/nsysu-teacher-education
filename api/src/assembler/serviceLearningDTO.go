package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// ServiceLearningDTO service-learning data transfer object
func ServiceLearningDTO(serviceLearnings *[]gorm.ServiceLearning) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, serviceLearning := range *serviceLearnings {
		result = append(result, map[string]interface{}{
			"ID":        serviceLearning.ID,
			"Type":      serviceLearning.Type,
			"Content":   serviceLearning.Content,
			"Session":   serviceLearning.Session,
			"Hours":     serviceLearning.Hours,
			"Start":     serviceLearning.Start,
			"End":       serviceLearning.End,
			"CreatedBy": serviceLearning.Student.Name,
		})
	}
	return result
}
