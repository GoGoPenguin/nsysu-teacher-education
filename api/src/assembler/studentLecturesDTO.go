package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentLecturesDTO student-lecture data transfer object
func StudentLecturesDTO(studentLectures *[]gorm.StudentLecture) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, studentLecture := range *studentLectures {
		if studentLecture.Lecture.ID != 0 {
			temp := map[string]interface{}{
				"ID": studentLecture.ID,

				"Lecture": map[string]interface{}{
					"ID":        studentLecture.Lecture.ID,
					"Name":      studentLecture.Lecture.Name,
					"MinCredit": studentLecture.Lecture.MinCredit,
					"Comment":   studentLecture.Lecture.Comment,
					"Status":    studentLecture.Lecture.Status,
				},
				"Pass": studentLecture.Pass,
			}
			if studentLecture.Student.ID != 0 {
				temp["Student"] = map[string]interface{}{
					"ID":      studentLecture.Student.ID,
					"Name":    studentLecture.Student.Name,
					"Account": studentLecture.Student.Account,
					"Major":   studentLecture.Student.Major,
					"Number":  studentLecture.Student.Number,
				}
			}
			result = append(result, temp)
		}
	}
	return result
}
