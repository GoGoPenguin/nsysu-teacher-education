package main

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/utils/hash"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

func studentSeeder() {
	tx := gorm.DB().Begin()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			tx.Rollback()
		}
	}()

	student := &gorm.Student{
		Name:     "測試帳號",
		Account:  "test",
		Password: hash.New("test"),
		Major:    "test",
		Number:   "0",
	}

	if gorm.StudentDao.GetByAccount(tx, student.Account) == nil {
		gorm.StudentDao.New(tx, student)
	}

	tx.Commit()
}
