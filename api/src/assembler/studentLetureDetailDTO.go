package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentLeturesDetailDTO student-leture data transfer object
func StudentLeturesDetailDTO(studentLeture *gorm.StudentLeture) map[string]interface{} {
	categories := []map[string]interface{}{}

	for _, category := range studentLeture.Leture.Categories {
		letureTypes := []map[string]interface{}{}

		for _, letureType := range category.Types {
			groups := []map[string]interface{}{}

			for _, group := range letureType.Groups {
				subjects := []map[string]interface{}{}

				for _, subject := range group.Subjects {
					subjects = append(subjects, map[string]interface{}{
						"ID":         subject.ID,
						"Name":       subject.Name,
						"Credit":     subject.Credit,
						"Compulsory": subject.Compulsory,
						"Pass":       subject.StudentSubject.Pass,
						"Score":      subject.StudentSubject.Score,
					})
				}

				groups = append(groups, map[string]interface{}{
					"ID":        group.ID,
					"MinCredit": group.MinCredit,
					"Subjects":  subjects,
				})
			}

			letureTypes = append(letureTypes, map[string]interface{}{
				"ID":        letureType.ID,
				"Name":      letureType.Name,
				"MinCredit": letureType.MinCredit,
				"Groups":    groups,
			})
		}
		categories = append(categories, map[string]interface{}{
			"ID":        category.ID,
			"Name":      category.Name,
			"MinCredit": category.MinCredit,
			"MinType":   category.MinType,
			"Types":     letureTypes,
		})
	}

	result := map[string]interface{}{
		"ID": studentLeture.ID,
		"Leture": map[string]interface{}{
			"ID":         studentLeture.Leture.ID,
			"Name":       studentLeture.Leture.Name,
			"MinCredit":  studentLeture.Leture.MinCredit,
			"Comment":    studentLeture.Leture.Comment,
			"Categories": categories,
		},
		"Pass": studentLeture.Pass,
	}

	return result
}
