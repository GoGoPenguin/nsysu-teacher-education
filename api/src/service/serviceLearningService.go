package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/nsysu/teacher-education/src/assembler"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/specification"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/typecast"
)

// CreateServieLearning create service-learning
func CreateServieLearning(serviceType, content, session string, hours uint, start, end time.Time) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
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
func GetServiceLearningList(account, start, length, search string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	if search == "同時認列教育實習服務暨志工服務" {
		search = "both"
	} else if search == "實習服務" {
		search = "internship"
	} else if search == "志工服務" {
		search = "volunteer"
	}

	var serviceLearnings *[]gorm.ServiceLearning
	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		serviceLearnings = gorm.ServiceLearningDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("start", specification.OrderDirectionDESC),
			specification.LikeSpecification([]string{"type", "content", "start", "end", "hours"}, search),
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		serviceLearnings = gorm.ServiceLearningDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.BiggerSpecification("start", time.Now().String()),
			specification.OrderSpecification("start", specification.OrderDirectionASC),
			specification.IsNullSpecification("deleted_at"),
		)
	}

	total := gorm.ServiceLearningDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	list := assembler.ServiceLearningDTO(serviceLearnings)
	result = map[string]interface{}{
		"list":            list,
		"recordsTotal":    total,
		"recordsFiltered": len(list),
	}

	return
}

// SingUpServiceLearning sudent sign up service-learning
func SingUpServiceLearning(account, serviceLearningID string) (result interface{}, e *errors.Error) {
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

	serviceLearning := gorm.ServiceLearningDao.Query(
		tx,
		specification.IDSpecification(serviceLearningID),
		specification.IsNullSpecification("deleted_at"),
		specification.BiggerSpecification("End", time.Now().String()),
	)

	if len(*serviceLearning) == 0 {
		return nil, errors.NotFoundError("service-learning ID " + serviceLearningID)
	}

	studentServiceLearning := &gorm.StudentServiceLearning{
		StudentID:         student.ID,
		ServiceLearningID: typecast.StringToUint(serviceLearningID),
	}

	gorm.StudentServiceLearningDao.New(tx, studentServiceLearning)

	return "success", nil
}

// GetSutdentServiceLearningList get the list of student service-learning
func GetSutdentServiceLearningList(account, start, length string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	var studentServiceLearnings *[]gorm.StudentServiceLearning
	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		studentServiceLearnings = gorm.StudentServiceLearningDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_service_learning`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		student := gorm.StudentDao.GetByAccount(tx, account)
		studentServiceLearnings = gorm.StudentServiceLearningDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_service_learning`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
			specification.StudentSpecification(student.ID),
		)
	}

	total := gorm.StudentServiceLearningDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.StudentServiceLearningsDTO(studentServiceLearnings),
		"recordsTotal":    total,
		"recordsFiltered": len(*studentServiceLearnings),
	}

	return
}

// UpdateStudentServiceLearning upload student service-learning review or reference file
func UpdateStudentServiceLearning(reference, review multipart.File, studentServiceLearningID, referenceFileName, reviewFileName string) (result string, e *errors.Error) {
	tx := gorm.DB().Begin()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	studentServiceLearning := gorm.StudentServiceLearningDao.GetByID(tx, typecast.StringToUint(studentServiceLearningID))

	if reference != nil {
		fileName := fmt.Sprintf("./assets/service-learning/%s-Reference", studentServiceLearningID)

		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			if err := os.Remove(fileName); err != nil {
				panic(err)
			}
		}

		file, err := os.OpenFile(
			fileName,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			0666,
		)
		if err != nil {
			panic(err)
		}
		io.Copy(file, reference)
		defer file.Close()

		studentServiceLearning.Reference = referenceFileName
		gorm.StudentServiceLearningDao.Update(tx, studentServiceLearning)
	}

	if review != nil {
		fileName := fmt.Sprintf("./assets/service-learning/%s-Review", studentServiceLearningID)

		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			if err := os.Remove(fileName); err != nil {
				panic(err)
			}
		}

		file, err := os.OpenFile(
			fileName,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			0666,
		)
		if err != nil {
			panic(err)
		}
		io.Copy(file, review)
		defer file.Close()

		studentServiceLearning.Review = reviewFileName
		gorm.StudentServiceLearningDao.Update(tx, studentServiceLearning)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return "success", nil
}

// UpdateStudentServiceLearningStatus update student-service-learning status
func UpdateStudentServiceLearningStatus(studentServiceLearningID, status, comment string) (result string, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	studentServiceLearning := gorm.StudentServiceLearningDao.GetByID(tx, typecast.StringToUint(studentServiceLearningID))
	studentServiceLearning.Status = status
	studentServiceLearning.Comment = comment
	gorm.StudentServiceLearningDao.Update(tx, studentServiceLearning)

	return "success", nil
}

// GetStudentServiceLearningFile get student-service-learning refernce or review file
func GetStudentServiceLearningFile(studentServiceLearningID, file string) (result map[string]string, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	studentServiceLearning := gorm.StudentServiceLearningDao.GetByID(tx, typecast.StringToUint(studentServiceLearningID))
	if studentServiceLearning == nil {
		return nil, errors.NotFoundError(file)
	}

	var (
		filePath string
		fileName string
	)

	if file == "reference" {
		fileName = studentServiceLearning.Reference
		filePath = fmt.Sprintf("./assets/service-learning/%s-Reference", studentServiceLearningID)
	} else {
		fileName = studentServiceLearning.Review
		filePath = fmt.Sprintf("./assets/service-learning/%s-Review", studentServiceLearningID)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, errors.NotFoundError(fileName)
	}

	return map[string]string{
		"Path": filePath,
		"Name": fileName,
	}, nil
}

// DeleteServiceLearning delete service-learning
func DeleteServiceLearning(serviceLearnginID string) (result string, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	gorm.ServiceLearningDao.Delete(tx, typecast.StringToUint(serviceLearnginID))

	return "success", nil
}

// UpdateServieLearning update service-learning
func UpdateServieLearning(serviceLearningID, serviceType, content, session string, hours uint, start, end time.Time) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	serviceLearning := gorm.ServiceLearningDao.GetByID(tx, typecast.StringToUint(serviceLearningID))
	serviceLearning.Type = serviceType
	serviceLearning.Content = content
	serviceLearning.Session = session
	serviceLearning.Hours = hours
	serviceLearning.Start = start
	serviceLearning.End = end

	gorm.ServiceLearningDao.Update(tx, serviceLearning)

	return "success", nil
}
