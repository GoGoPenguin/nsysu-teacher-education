package main

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

func letureSeeder() {
	tx := gorm.DB().Begin()

	letureData := []map[string]interface{}{
		map[string]interface{}{
			"ID":        uint(1),
			"Name":      "自然科學領域化學專長",
			"MinCredit": uint(42),
			"Comment":   "適合化學系所等",
		},
	}

	for _, data := range letureData {
		leture := &gorm.Leture{
			Name:      data["Name"].(string),
			MinCredit: data["MinCredit"].(uint),
			Comment:   data["Comment"].(string),
		}
		leture.ID = data["ID"].(uint)

		if gorm.LetureDao.GetByName(tx, leture.Name) == nil {
			gorm.LetureDao.New(tx, leture)
		}
	}

	categoryData := []map[string]interface{}{
		map[string]interface{}{
			"ID":        uint(1),
			"LetureID":  uint(1),
			"Name":      "領域核心課程",
			"MinCredit": uint(4),
			"MinType":   uint(0),
		},
	}

	for _, data := range categoryData {
		letureCategory := &gorm.LetureCategory{
			LetureID:  data["LetureID"].(uint),
			Name:      data["Name"].(string),
			MinCredit: data["MinCredit"].(uint),
			MinType:   data["MinType"].(uint),
		}
		letureCategory.ID = data["ID"].(uint)

		if gorm.LetureCategoryDao.GetByLetureAndName(tx, letureCategory.LetureID, letureCategory.Name) == nil {
			gorm.LetureCategoryDao.New(tx, letureCategory)
		}
	}

	typeData := []map[string]interface{}{
		map[string]interface{}{
			"ID":               uint(1),
			"LetureCategoryID": uint(1),
			"Name":             "探究與實作",
			"MinCredit":        uint(0),
		},
	}

	for _, data := range typeData {
		letureType := &gorm.LetureType{
			LetureCategoryID: data["LetureCategoryID"].(uint),
			Name:             data["Name"].(string),
			MinCredit:        data["MinCredit"].(uint),
		}
		letureType.ID = data["ID"].(uint)

		if gorm.LetureTypeDao.GetByCategoryAndName(tx, letureType.LetureCategoryID, letureType.Name) == nil {
			gorm.LetureTypeDao.New(tx, letureType)
		}
	}

	groupData := []map[string]interface{}{
		map[string]interface{}{
			"ID":           uint(1),
			"LetureTypeID": uint(1),
			"MinCredit":    uint(0),
		},
	}

	for _, data := range groupData {
		subjectGroup := &gorm.SubjectGroup{
			LetureTypeID: data["LetureTypeID"].(uint),
			MinCredit:    data["MinCredit"].(uint),
		}
		subjectGroup.ID = data["ID"].(uint)

		if gorm.SubjectGroupDao.GetByIDAndType(tx, subjectGroup.ID, subjectGroup.LetureTypeID) == nil {
			gorm.SubjectGroupDao.New(tx, subjectGroup)
		}
	}

	subjectData := []map[string]interface{}{
		map[string]interface{}{
			"ID":             uint(1),
			"SubjectGroupID": uint(1),
			"Name":           "生活科技概論",
			"Credit":         uint(3),
			"Compulsory":     false,
		},
	}

	for _, data := range subjectData {
		subject := &gorm.Subject{
			SubjectGroupID: data["SubjectGroupID"].(uint),
			Name:           data["Name"].(string),
			Credit:         data["Credit"].(uint),
			Compulsory:     data["Compulsory"].(bool),
		}
		subject.ID = data["ID"].(uint)

		if gorm.SubjectDao.GetByName(tx, subject.Name) == nil {
			gorm.SubjectDao.New(tx, subject)
		}
	}

	tx.Commit()
}
