package service

import (
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/nsysu/teacher-education/src/assembler"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/specification"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/typecast"
)

// CreateCourse create a course
func CreateCourse(topic, courseType string, file multipart.File, header *multipart.FileHeader, start, end time.Time) (result interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	f, err := os.OpenFile("./assets/course/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, file)
	course := &gorm.Course{
		Topic:       topic,
		Information: header.Filename,
		Type:        courseType,
		Start:       start,
		End:         end,
	}
	gorm.CourseDao.New(tx, course)

	return "success", nil
}

// GetCourse get course list
func GetCourse(account, start, length string) (result map[string]interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	var courses *[]gorm.Course
	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		courses = gorm.CourseDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("start", specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		courses = gorm.CourseDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.BiggerSpecification("start", time.Now().String()),
			specification.OrderSpecification("start", specification.OrderDirectionASC),
			specification.IsNullSpecification("deleted_at"),
		)
	}

	total := gorm.CourseDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.CoursesDTO(courses),
		"recordsTotal":    total,
		"recordsFiltered": total,
	}

	return result, nil
}

// GetInformation get the information of course
func GetInformation(filename string) (result string, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	course := gorm.CourseDao.GetByInformation(tx, filename)
	if course == nil {
		return "", error.NotFoundError(filename)
	}
	if _, err := os.Stat("./assets/course/" + filename); os.IsNotExist(err) {
		return "", error.NotFoundError(filename)
	}

	return "./assets/course/" + filename, nil
}
