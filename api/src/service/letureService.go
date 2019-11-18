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

	var (
		letures  *[]gorm.Leture
		filtered int
	)

	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		letures = gorm.LetureDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("created_at", specification.OrderDirectionASC),
			specification.LikeSpecification([]string{"name", "comment", "min_credit", "status"}, search),
			specification.IsNullSpecification("deleted_at"),
		)

		filtered = gorm.LetureDao.Count(
			tx,
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
		"recordsFiltered": filtered,
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
	tx := gorm.DB().Set("gorm:auto_preload", true).Begin()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
			tx.Rollback()
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

	for _, category := range (*leture)[0].Categories {
		for _, letureType := range category.Types {
			for _, group := range letureType.Groups {
				for _, subject := range group.Subjects {
					studentSubject := &gorm.StudentSubject{
						StudentLetureID: studentLeture.ID,
						SubjectID:       subject.ID,
						Pass:            false,
					}
					gorm.StudentSubjectDao.New(tx, studentSubject)
				}
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	return "success", nil
}

// GetSutdentLetureList get the list of student leture
func GetSutdentLetureList(account, start, length string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	var (
		studentLetures *[]gorm.StudentLeture
		filtered       int
	)

	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		studentLetures = gorm.StudentLetureDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_leture`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
		)

		filtered = gorm.StudentLetureDao.Count(
			tx,
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		student := gorm.StudentDao.GetByAccount(tx, account)
		studentLetures = gorm.StudentLetureDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_leture`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
			specification.StudentSpecification(student.ID),
		)
	}

	total := gorm.StudentLetureDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.StudentLeturesDTO(studentLetures),
		"recordsTotal":    total,
		"recordsFiltered": filtered,
	}

	return
}

// GetStudentLetureDetail get studnet leture detail
func GetStudentLetureDetail(studentLetureID string) (result interface{}, e *errors.Error) {
	tx := gorm.DB().Preload("Leture.Categories.Types.Groups.Subjects.StudentSubject").Preload("Student")

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	studentLeture := gorm.StudentLetureDao.GetByID(tx, typecast.StringToUint(studentLetureID))

	return assembler.StudentLeturesDetailDTO(studentLeture), nil
}
