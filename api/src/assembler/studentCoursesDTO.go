package assembler

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

// StudentCoursesDTO student-course data transfer object
func StudentCoursesDTO(studentCourses *[]gorm.StudentCourse) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, studentCourse := range *studentCourses {
		result = append(result, map[string]interface{}{
			"Student": studentCourse.Student.Name,
			"Course":  studentCourse.Course.Topic,
			"Meal":    studentCourse.Meal,
			"Status":  studentCourse.Status,
			"Review":  studentCourse.Review,
			"Comment": studentCourse.Comment,
		})
	}
	return result
}
