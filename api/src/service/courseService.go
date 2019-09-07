package service

import (
	"time"

	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

// CreateCourse create a course
func CreateCourse(topic, information, courseType string, start, end time.Time) (result interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	course := &gorm.Course{
		Topic:       topic,
		Information: information,
		Type:        courseType,
		Start:       start,
		End:         end,
	}
	gorm.CourseDao.New(tx, course)

	return "success", nil
}
