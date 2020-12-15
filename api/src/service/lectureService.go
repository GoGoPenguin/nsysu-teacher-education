package service

import (
	"github.com/nsysu/teacher-education/src/assembler"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/specification"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/typecast"
	"gorm.io/gorm/clause"
)

// GetLectures get lecture list
func GetLectures(account, start, length, search string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	var (
		lectures *[]gorm.Lecture
		filtered int64
	)

	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		lectures = gorm.LectureDao.Query(
			tx,
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("created_at", specification.OrderDirectionASC),
			specification.LikeSpecification([]string{"concat(name,comment,min_credit,status,comment)"}, search),
			specification.IsNullSpecification("deleted_at"),
		)

		filtered = gorm.LectureDao.Count(
			tx,
			specification.LikeSpecification([]string{"concat(name,comment,min_credit,status,comment)"}, search),
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		lectures = gorm.LectureDao.Query(
			tx,
			specification.OrderSpecification("created_at", specification.OrderDirectionASC),
			specification.StatusSpecification(gorm.LectureDao.Enable),
			specification.IsNullSpecification("deleted_at"),
		)
	}

	total := gorm.LectureDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.LecturesDTO(lectures),
		"recordsTotal":    total,
		"recordsFiltered": filtered,
	}

	return result, nil

}

// GetLectureDetail get lecture detail
func GetLectureDetail(lectureID string) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	tx = tx.Preload("Categories.Types.Groups.Subjects.StudentSubject").Preload(clause.Associations)
	lecture := gorm.LectureDao.GetByID(tx, typecast.StringToUint(lectureID))
	return lecture, nil
}

// SingUpLecture student sign up lecture
func SingUpLecture(account, lectureID string) (result interface{}, e *errors.Error) {
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

	lecture := gorm.LectureDao.Query(
		tx,
		specification.IDSpecification(lectureID),
		specification.IsNullSpecification("deleted_at"),
		specification.StatusSpecification(gorm.LectureDao.Enable),
	)

	if len(*lecture) == 0 {
		return nil, errors.NotFoundError("service-learning ID " + lectureID)
	}

	studentLecture := &gorm.StudentLecture{
		StudentID: student.ID,
		LectureID: typecast.StringToUint(lectureID),
		Pass:      false,
	}

	gorm.StudentLectureDao.New(tx, studentLecture)

	for _, category := range (*lecture)[0].Categories {
		for _, lectureType := range category.Types {
			for _, group := range lectureType.Groups {
				for _, subject := range group.Subjects {
					studentSubject := &gorm.StudentSubject{
						StudentLectureID: studentLecture.ID,
						SubjectID:        subject.ID,
					}
					gorm.StudentSubjectDao.New(tx, studentSubject)
				}
			}
		}
	}
	tx.Rollback()

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	return "success", nil
}

// GetStudentLectureList get the list of student lecture
func GetStudentLectureList(account, start, length string) (result map[string]interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	var (
		studentLectures *[]gorm.StudentLecture
		filtered        int64
	)

	if operator := gorm.AdminDao.GetByAccount(tx, account); operator != nil {
		studentLectures = gorm.StudentLectureDao.Query(
			tx.Preload("Lecture").Preload("Student"),
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_lecture`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
		)

		filtered = gorm.StudentLectureDao.Count(
			tx,
			specification.IsNullSpecification("deleted_at"),
		)
	} else {
		student := gorm.StudentDao.GetByAccount(tx, account)
		studentLectures = gorm.StudentLectureDao.Query(
			tx.Preload("Lecture"),
			specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
			specification.OrderSpecification("`student_lecture`."+specification.IDColumn, specification.OrderDirectionDESC),
			specification.IsNullSpecification("deleted_at"),
			specification.StudentSpecification(student.ID),
		)
	}

	total := gorm.StudentLectureDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.StudentLecturesDTO(studentLectures),
		"recordsTotal":    total,
		"recordsFiltered": filtered,
	}

	return
}

// GetStudentLectureDetail get studnet lecture detail
func GetStudentLectureDetail(studentLectureID string) (result interface{}, e *errors.Error) {
	tx := gorm.DB().Preload("Lecture.Categories.Types.Groups.Subjects.StudentSubject")

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	studentLecture := gorm.StudentLectureDao.GetByID(tx, typecast.StringToUint(studentLectureID))

	return assembler.StudentLecturesDetailDTO(studentLecture), nil
}

// UpdateStudentSubject update student subject
func UpdateStudentSubject(account, studentLectureID, subjectID, name, year, semester, credit, score string) (result interface{}, e *errors.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = errors.UnexpectedError()
		}
	}()

	if gorm.StudentDao.GetByAccount(tx, account) == nil {
		return nil, errors.NotFoundError("Student " + account)
	}

	studentSubject := gorm.StudentSubjectDao.GetByLectureAndSubject(tx, typecast.StringToUint(studentLectureID), typecast.StringToUint(subjectID))

	if studentSubject == nil {
		return nil, errors.NotFoundError("Student Subject")
	}

	studentSubject.Name = name
	studentSubject.Year = year
	studentSubject.Semester = semester
	studentSubject.Credit = credit
	studentSubject.Score = score

	gorm.StudentSubjectDao.Update(tx, studentSubject)

	return "success", nil
}

// UpdateStudentLecturePass update student lecture pass
func UpdateStudentLecturePass(account, lectureID, pass string) (result interface{}, e *errors.Error) {
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

	studentLecture := gorm.StudentLectureDao.GetByLectureAndStudent(tx, typecast.StringToUint(lectureID), student.ID)
	if studentLecture == nil {
		return nil, errors.NotFoundError("Student Lecture")
	}

	studentLecture.Pass = typecast.StringToBool(pass)
	gorm.StudentLectureDao.Update(tx, studentLecture)

	return "success", nil
}
