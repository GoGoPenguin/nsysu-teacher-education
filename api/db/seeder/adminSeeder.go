package main

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/utils/hash"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

func adminSeeder() {
	tx := gorm.DB().Begin()

	defer func() {
		if r := recover(); r != nil {
			logger.Error(r)
			tx.Rollback()
		}
	}()

	admin := &gorm.Admin{
		Account:  "admin",
		Password: hash.New("P@ssword"),
	}

	if gorm.AdminDao.GetByAccount(tx, admin.Account) == nil {
		gorm.AdminDao.New(tx, admin)
	}

	tx.Commit()
}
