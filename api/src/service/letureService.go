package service

import (
	"github.com/nsysu/teacher-education/src/assembler"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/specification"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/typecast"
)

// GetLetures get leture list
func GetLetures(account, start, length, search string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	var letures *[]gorm.Leture
	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		letures = gorm.LetureDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("created_at", specification.OrderDirectionASC),
			specification.LikeSpecification([]string{"name", "comment", "min_credit", "status"}, search),
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		letures = gorm.LetureDao.Query(
			tx,
			specification.OrderSpecification("created_at", specification.OrderDirectionASC),
			specification.StatusSpecification(gorm.LetureDao.Enable),
			specification.IsNullSpecification("deleted_at"),
		)
	}

	total := gorm.LetureDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.LeturesDTO(letures),
		"recordsTotal":    total,
		"recordsFiltered": len(*letures),
	}

	return result, nil

}

// GetLetureDetail get leture detail
func GetLetureDetail(letureID string) (result interface{}, e *errors.Error) {
	tx := gorm.DB().Set("gorm:auto_preload", true)

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	leture := gorm.LetureDao.GetByID(tx, typecast.StringToUint(letureID))
	return leture, nil
}

// SingUpLeture sudent sign up leture
func SingUpLeture(account, letureID string) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	student := gorm.StudentDao.GetByAccount(tx, account)

	if student == nil {
		return nil, errors.NotFoundError("Student " + account)
	}

	leture := gorm.LetureDao.Query(
		tx,
		specification.IDSpecification(letureID),
		specification.IsNullSpecification("deleted_at"),
		specification.StatusSpecification(gorm.LetureDao.Enable),
	)

	if len(*leture) == 0 {
		return nil, errors.NotFoundError("service-learning ID " + letureID)
	}

	studentLeture := &gorm.StudentLeture{
		StudentID: student.ID,
		LetureID:  typecast.StringToUint(letureID),
		Pass:      false,
	}

	gorm.StudentLetureDao.New(tx, studentLeture)

	return "success", nil
}
