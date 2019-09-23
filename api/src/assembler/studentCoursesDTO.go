package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentCoursesDTO student-course data transfer object
func StudentCoursesDTO(studentCourses *[]gorm.StudentCourse) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, studentCourse := range *studentCourses {
		result = append(result, map[string]interface{}{
			"Student": map[string]interface{}{
				"Name":    studentCourse.Student.Name,
				"ACcount": studentCourse.Student.Account,
				"Major":   studentCourse.Student.Major,
				"Number":  studentCourse.Student.Number,
			},
			"Course": map[string]interface{}{
				"Topic": studentCourse.Course.Topic,
				"Type":  studentCourse.Course.Type,
				"Start": studentCourse.Course.Start.String(),
				"End":   studentCourse.Course.End.String(),
			},
			"Meal":    studentCourse.Meal,
			"Status":  studentCourse.Status,
			"Review":  studentCourse.Review,
			"Comment": studentCourse.Comment,
		})
	}
	return result
}
