package service

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/nsysu/teacher-education/src/assembler"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/specification"
	"github.com/nsysu/teacher-education/src/utils/hash"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"github.com/nsysu/teacher-education/src/utils/typecast"
)

// CreateStudents create students by csv file
func CreateStudents(file multipart.File) (result interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		student := &gorm.Student{
			Name:     line[0],
			Account:  line[1],
			Password: hash.New(line[2]),
			Major:    line[3],
			Number:   line[4],
		}

		gorm.StudentDao.New(tx, student)
	}

	return "success", nil
}

// GetStudents get user list
func GetStudents(start, length string) (result map[string]interface{}, e *error.Error) {
	tx := gorm.DB()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			e = error.UnexpectedError()
		}
	}()

	students := gorm.StudentDao.Query(
		tx,
		specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
		specification.OrderSpecification(specification.IDColumn, specification.OrderDirectionDESC),
		specification.IsNullSpecification("deleted_at"),
	)

	total := gorm.StudentDao.Count(
		tx,
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.StudentsDTO(students),
		"recordsTotal":    total,
		"recordsFiltered": total,
	}

	return result, nil
}
