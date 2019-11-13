package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentCoursesDTO student-course data transfer object
func StudentCoursesDTO(studentCourses *[]gorm.StudentCourse) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, studentCourse := range *studentCourses {
		result = append(result, map[string]interface{}{
			"ID": studentCourse.ID,
			"Student": map[string]interface{}{
				"Name":    studentCourse.Student.Name,
				"Account": studentCourse.Student.Account,
				"Major":   studentCourse.Student.Major,
				"Number":  studentCourse.Student.Number,
			},
			"Course": map[string]interface{}{
				"ID":    studentCourse.Course.ID,
				"Topic": studentCourse.Course.Topic,
				"Type":  studentCourse.Course.Type,
				"Start": studentCourse.Course.Start,
				"End":   studentCourse.Course.End,
			},
			"Meal":    studentCourse.Meal,
			"Status":  studentCourse.Status,
			"Review":  studentCourse.Review,
			"Comment": studentCourse.Comment,
		})
	}
	return result
}
