package service

import (
	"time"

	"github.com/nsysu/teacher-education/src/assembler"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/specification"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/typecast"
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

// GetServiceLearningList get service-learning list
func GetServiceLearningList(account, start, length string) (result map[string]interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	var serviceLearnings *[]gorm.ServiceLearning
	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		serviceLearnings = gorm.ServiceLearningDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`service_learning`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("`service_learning`.deleted_at"),
		)
	} else {
		student := gorm.StudentDao.GetByAccount(tx, account)
		serviceLearnings = gorm.ServiceLearningDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`service_learning`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
			specification.StudentSpecification(student.ID),
		)
	}

	total := gorm.ServiceLearningDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.ServiceLearningDTO(serviceLearnings),
		"recordsTotal":    total,
		"recordsFiltered": total,
	}

	return
}
