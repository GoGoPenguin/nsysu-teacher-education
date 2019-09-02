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
func GetStudents(index, count string) (result interface{}, e *error.Error) {
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
		specification.PaginationSpecification(index, specification.OrderDirectionASC),
		specification.RoleSpecification(gorm.UserDao.Roletudent),
		specification.IsNullSpecification("deleted_at"),
		specification.LimitSpecification(count),
	)

	result = map[string]interface{}{
		"List":   assembler.UsersDTO(users),
		"LastID": (*users)[len(*users)-1].ID,
	}

	return result, nil
}
