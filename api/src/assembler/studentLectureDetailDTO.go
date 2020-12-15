package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentLecturesDetailDTO student-lecture data transfer object
func StudentLecturesDetailDTO(studentLecture *gorm.StudentLecture) map[string]interface{} {
	categories := []map[string]interface{}{}

	for _, category := range studentLecture.Lecture.Categories {
		lectureTypes := []map[string]interface{}{}

		for _, lectureType := range category.Types {
			groups := []map[string]interface{}{}

			for _, group := range lectureType.Groups {
				subjects := []map[string]interface{}{}

				for _, subject := range group.Subjects {
					subjects = append(subjects, map[string]interface{}{
						"ID":            subject.ID,
						"Name":          subject.Name,
						"Credit":        subject.Credit,
						"Compulsory":    subject.Compulsory,
						"StudentName":   subject.StudentSubject.Name,
						"Year":          subject.StudentSubject.Year,
						"Semester":      subject.StudentSubject.Semester,
						"StudentCredit": subject.StudentSubject.Credit,
						"Score":         subject.StudentSubject.Score,
					})
				}

				groups = append(groups, map[string]interface{}{
					"ID":        group.ID,
					"MinCredit": group.MinCredit,
					"Subjects":  subjects,
				})
			}

			lectureTypes = append(lectureTypes, map[string]interface{}{
				"ID":        lectureType.ID,
				"Name":      lectureType.Name,
				"MinCredit": lectureType.MinCredit,
				"Groups":    groups,
			})
		}
		categories = append(categories, map[string]interface{}{
			"ID":        category.ID,
			"Name":      category.Name,
			"MinCredit": category.MinCredit,
			"MinType":   category.MinType,
			"Types":     lectureTypes,
		})
	}

	result := map[string]interface{}{
		"ID": studentLecture.ID,
		"Lecture": map[string]interface{}{
			"ID":         studentLecture.Lecture.ID,
			"Name":       studentLecture.Lecture.Name,
			"MinCredit":  studentLecture.Lecture.MinCredit,
			"Comment":    studentLecture.Lecture.Comment,
			"Categories": categories,
		},
		"Pass": studentLecture.Pass,
	}

	return result
}
