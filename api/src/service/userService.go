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
		user := &gorm.User{
			Name:     line[0],
			Account:  line[1],
			Password: hash.New(line[2]),
			Role:     gorm.UserDao.Roletudent,
		}

		gorm.UserDao.New(tx, user)
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

	// result = []map[string]interface{}{}
	users := gorm.UserDao.Query(
		tx,
		specification.PaginationSpecification(typecast.StringToInt(start), typecast.StringToInt(length)),
		specification.RoleSpecification(gorm.UserDao.Roletudent),
		specification.IsNullSpecification("deleted_at"),
	)
	total := gorm.UserDao.Count(
		tx,
		specification.RoleSpecification(gorm.UserDao.Roletudent),
		specification.IsNullSpecification("deleted_at"),
	)

	result = map[string]interface{}{
		"list":            assembler.UsersDTO(users),
		"recordsTotal":    total,
		"recordsFiltered": total,
	}

	return result, nil
}
