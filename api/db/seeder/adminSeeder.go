package main

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/utils/hash"
)

func adminSeeder() {
	tx := gorm.DB().Begin()

	admin := &gorm.Admin{
		Account:  "admin",
		Password: hash.New("P@ssword"),
	}

	if gorm.AdminDao.GetByAccount(tx, admin.Account) == nil {
		gorm.AdminDao.New(tx, admin)
	}

	tx.Commit()
}
