package service

import (
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

// CreateCourse create a course
func CreateCourse(topic, courseType string, file multipart.File, header *multipart.FileHeader, start, end time.Time) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	logger.Debug(header)

	if header != nil {
		course := &gorm.Course{
			Topic:       topic,
			Information: header.Filename,
			Type:        courseType,
			Start:       start,
			End:         end,
		}
		gorm.CourseDao.New(tx, course)

		f, err := os.OpenFile("./assets/course/"+typecast.ToString(course.ID), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		io.Copy(f, file)
	} else {
		course := &gorm.Course{
			Topic:       topic,
			Information: "",
			Type:        courseType,
			Start:       start,
			End:         end,
		}
		gorm.CourseDao.New(tx, course)
	}

	return "success", nil
}

// GetCourse get course list
func GetCourse(account, start, length, search string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	var (
		courses *[]gorm.Course
		filered int
	)

	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		courses = gorm.CourseDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("start", specification.OrderDirectionDESC),
			specification.LikeSpecification([]string{"topic", "information", "type", "start", "end"}, search),
			specification.IsNullSpecification("deleted_at"),
		)

		filered = gorm.CourseDao.Count(
			tx,
			specification.LikeSpecification([]string{"topic", "information", "type", "start", "end"}, search),
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		courses = gorm.CourseDao.Query(
			tx,
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
		"recordsFiltered": filered,
	}

	return result, nil
}

// GetInformation get the information of course
func GetInformation(courseID string) (result map[string]string, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	course := gorm.CourseDao.GetByID(tx, typecast.StringToUint(courseID))
	if course == nil {
		return nil, errors.NotFoundError(courseID)
	}
	if _, err := os.Stat("./assets/course/" + courseID); os.IsNotExist(err) {
		return nil, errors.NotFoundError(courseID)
	}

	return map[string]string{
		"Path": "./assets/course/" + courseID,
		"Name": course.Information,
	}, nil
}

// SingUpCourse student sign up course
func SingUpCourse(account, courseID, meal string) (result interface{}, e *errors.Error) {
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

	course := gorm.CourseDao.Query(
		tx,
		specification.IDSpecification(courseID),
		specification.IsNullSpecification("deleted_at"),
		// specification.BiggerSpecification("start", time.Now().String()),
	)

	if len(*course) == 0 {
		return nil, errors.NotFoundError("course ID " + courseID)
	}

	srudentCourse := &gorm.StudentCourse{
		StudentID: student.ID,
		CourseID:  typecast.StringToUint(courseID),
		Meal:      meal,
	}

	gorm.StudentCourseDao.New(tx, srudentCourse)

	return "success", nil
}

// GetStudentCourseList get the list of student course
func GetStudentCourseList(account, start, length string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	var (
		studentCourses *[]gorm.StudentCourse
		filtered       int
	)

	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		studentCourses = gorm.StudentCourseDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_course`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
		)

		filtered = gorm.StudentCourseDao.Count(
			tx,
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		student := gorm.StudentDao.GetByAccount(tx, account)
		studentCourses = gorm.StudentCourseDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_course`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
			specification.StudentSpecification(student.ID),
		)
	}

	total := gorm.StudentCourseDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.StudentCoursesDTO(studentCourses),
		"recordsTotal":    total,
		"recordsFiltered": filtered,
	}

	return
}

// UpdateStudentCourseReview update student-course review
func UpdateStudentCourseReview(id, review string) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	studentCourse := gorm.StudentCourseDao.GetByID(tx, typecast.StringToUint(id))
	studentCourse.Review = review
	gorm.StudentCourseDao.Update(tx, studentCourse)

	return
}

// UpdateStudentCourseStatus update student-course status
func UpdateStudentCourseStatus(id, status, comment string) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	studentCourse := gorm.StudentCourseDao.GetByID(tx, typecast.StringToUint(id))
	studentCourse.Status = status
	studentCourse.Comment = comment
	gorm.StudentCourseDao.Update(tx, studentCourse)

	return
}

// DeleteCourse delete course
func DeleteCourse(courseID string) (result string, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	gorm.CourseDao.Delete(tx, typecast.StringToUint(courseID))

	return "success", nil
}

// UpdateCourse update course
func UpdateCourse(courseID, topic, courseType string, file multipart.File, header *multipart.FileHeader, start, end time.Time) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	course := gorm.CourseDao.GetByID(tx, typecast.StringToUint(courseID))
	course.Topic = topic
	course.Type = courseType
	course.Start = start
	course.End = end

	if file != nil {
		course.Information = header.Filename

		var fileName = "./assets/course/" + typecast.ToString(course.ID)
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			if err := os.Remove(fileName); err != nil {
				panic(err)
			}
		}

		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		io.Copy(f, file)
	}

	gorm.CourseDao.Update(tx, course)

	return "success", nil
}
