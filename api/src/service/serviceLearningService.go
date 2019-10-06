package service

import (
	"time"

	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

// CreateServieLearning create service-learning
func CreateServieLearning(serviceType, content, session string, hours uint, start, end time.Time) (result interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	serviceLearning := &gorm.ServiceLearning{
		Type:    serviceType,
		Content: content,
		Session: session,
		Hours:   hours,
		Start:   start,
		End:     end,
	}
	gorm.ServiceLearningDao.New(tx, serviceLearning)

	return "success", nil
}
