package main

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

func letureSeeder() {
	tx := gorm.DB().Begin()

	letureData := []map[string]interface{}{
		map[string]interface{}{"ID": uint(1), "Name": "自然科學領域化學專長", "MinCredit": uint(42), "Comment": "適合化學系所等"},
		map[string]interface{}{"ID": uint(2), "Name": "藝術領域藝術生活科－表演藝術專長", "MinCredit": uint(38), "Comment": "劇場藝術學系所等"},
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
		map[string]interface{}{"ID": uint(1), "LetureID": uint(1), "Name": "領域核心課程", "MinCredit": uint(4), "MinType": uint(0)},
		map[string]interface{}{"ID": uint(2), "LetureID": uint(1), "Name": "領域內跨科課程", "MinCredit": uint(8), "MinType": uint(2)},
		map[string]interface{}{"ID": uint(3), "LetureID": uint(1), "Name": "化學專長課程", "MinCredit": uint(30), "MinType": uint(0)},

		map[string]interface{}{"ID": uint(4), "LetureID": uint(2), "Name": "領域核心課程", "MinCredit": uint(3), "MinType": uint(0)},
		map[string]interface{}{"ID": uint(5), "LetureID": uint(2), "Name": "表演藝術專長課程", "MinCredit": uint(30), "MinType": uint(0)},
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
		map[string]interface{}{"ID": uint(1), "LetureCategoryID": uint(1), "Name": "探究與實作", "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(2), "LetureCategoryID": uint(2), "Name": "生物專長", "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(3), "LetureCategoryID": uint(2), "Name": "物理專長", "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(4), "LetureCategoryID": uint(2), "Name": "地球科學專長", "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(5), "LetureCategoryID": uint(3), "Name": "化學基本知識", "MinCredit": uint(15)},
		map[string]interface{}{"ID": uint(6), "LetureCategoryID": uint(3), "Name": "化學實驗能力", "MinCredit": uint(7)},
		map[string]interface{}{"ID": uint(7), "LetureCategoryID": uint(3), "Name": "跨學科與應用知識", "MinCredit": uint(2)},

		map[string]interface{}{"ID": uint(8), "LetureCategoryID": uint(4), "Name": "藝術領域核心", "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(9), "LetureCategoryID": uint(5), "Name": "理解與應用", "MinCredit": uint(10)},
		map[string]interface{}{"ID": uint(10), "LetureCategoryID": uint(5), "Name": "實踐與展現", "MinCredit": uint(12)},
		map[string]interface{}{"ID": uint(11), "LetureCategoryID": uint(5), "Name": "表演藝術進階跨科/跨領域", "MinCredit": uint(6)},
		map[string]interface{}{"ID": uint(12), "LetureCategoryID": uint(5), "Name": "教學知能", "MinCredit": uint(2)},
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
		map[string]interface{}{"ID": uint(1), "LetureTypeID": uint(1), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(2), "LetureTypeID": uint(2), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(3), "LetureTypeID": uint(3), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(4), "LetureTypeID": uint(4), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(5), "LetureTypeID": uint(5), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(6), "LetureTypeID": uint(6), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(7), "LetureTypeID": uint(7), "MinCredit": uint(0)},

		map[string]interface{}{"ID": uint(8), "LetureTypeID": uint(8), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(9), "LetureTypeID": uint(9), "MinCredit": uint(4)},
		map[string]interface{}{"ID": uint(10), "LetureTypeID": uint(9), "MinCredit": uint(6)},
		map[string]interface{}{"ID": uint(11), "LetureTypeID": uint(10), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(12), "LetureTypeID": uint(11), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(13), "LetureTypeID": uint(12), "MinCredit": uint(0)},
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
		map[string]interface{}{"SubjectGroupID": uint(1), "Name": "生活科技概論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(1), "Name": "專題研究(一)(二)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(1), "Name": "探究與實作", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(2), "Name": "普通生物學(一)(二)", "Credit": uint(4), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(2), "Name": "遺傳學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(2), "Name": "生態學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(2), "Name": "動物生理學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(2), "Name": "植物生理學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(2), "Name": "細胞生物學", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(3), "Name": "普通物理及實驗", "Credit": uint(4), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(3), "Name": "力學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(3), "Name": "電磁學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(3), "Name": "光學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(3), "Name": "熱統計物理", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(3), "Name": "近代物理", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(3), "Name": "物理發展史", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(4), "Name": "天文學導論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(4), "Name": "地質學", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(4), "Name": "基礎海洋學(海洋生態學)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(4), "Name": "地球物理概論(海洋物理概論)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(4), "Name": "環境變遷與生態保育", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "普通化學(一)(二)", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "有機化學(一)(二)", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "分析化學", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "無機化學(一)(二)", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "物理化學(一)(二)", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "儀器分析(一)(二)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "有機化學反應", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "有機光譜概論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "有機合成", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "材料化學導論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "原子光譜分析技術", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "工業質譜分析應用", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "奈米生醫分析", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "化學及生物感測器", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "生物無機化學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "材料化學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "有機金屬化學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "高分子化學導論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "群論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "奈米薄層結構分析", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "化學數學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "核磁共振光譜與影像導論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "氣膠科學導論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "化學實驗之程式應用", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(5), "Name": "量子化學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "普通化學實驗(一)", "Credit": uint(1), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "普通化學實驗(二)", "Credit": uint(1), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "有機化學實驗(一)", "Credit": uint(1), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "有機化學實驗(二)", "Credit": uint(1), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "分析化學實驗", "Credit": uint(1), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "物理化學實驗(一)", "Credit": uint(1), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "物理化學實驗(二)", "Credit": uint(1), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "儀器分析實驗(一)", "Credit": uint(1), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(6), "Name": "儀器分析實驗(二)", "Credit": uint(1), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "生物化學(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "生物統計學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "生物科學研究法", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "應用生物方法學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "應用生物實務", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "奈米科技概論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "材料科學導論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "材料熱力學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "材料物理性質", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "材料物理性質", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "高分子分析", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "工業化學講座", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(7), "Name": "食品安全分析概論", "Credit": uint(2), "Compulsory": false},

		map[string]interface{}{"SubjectGroupID": uint(8), "Name": "藝術概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(8), "Name": "劇場美學", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "表演技藝導論", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "劇場設計導論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "西洋戲劇史(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "西洋戲劇史(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "中國戲劇史(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "中國戲劇史(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "中西舞蹈史", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(9), "Name": "西洋藝術史", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "劇本導讀", "Credit": uint(2), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "導演概論", "Credit": uint(2), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "劇本解析", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "戲劇評論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "臺灣劇場", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "現當代華語文戲劇選讀", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "臺灣傳統戲曲", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(10), "Name": "劇本創作(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "劇場製作基礎", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "基礎表演", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "導演(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "觀點技巧表演", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "進階表演(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "傳統戲曲表演身段(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "基礎發聲", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "聲音訓練", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "歌唱技巧", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "肢體開發(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "肢體開發(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "舞蹈技巧", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "舞蹈創作", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "燈光設計及技術", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "劇場服裝技術", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "人物服裝畫", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "劇場服裝設計", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "舞臺化妝與造型", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "基礎設計", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "進階技術繪圖", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(11), "Name": "模型製作", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "音樂劇表演", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "音樂劇製作", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "創意表演", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "排演", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "劇場管理導論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "藝術管理概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "藝術行銷概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "藝術與文化環境概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "(碩)文化創意產業", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "(碩)藝術教育與社區文化", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "(碩)展演設計論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "(碩)展演設計(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(12), "Name": "(碩)經典服裝設計", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(13), "Name": "創作性戲劇", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(13), "Name": "戲劇治療與PBT設計", "Credit": uint(3), "Compulsory": false},
	}

	for _, data := range subjectData {
		subject := &gorm.Subject{
			SubjectGroupID: data["SubjectGroupID"].(uint),
			Name:           data["Name"].(string),
			Credit:         data["Credit"].(uint),
			Compulsory:     data["Compulsory"].(bool),
		}

		if gorm.SubjectDao.GetByName(tx, subject.Name) == nil {
			gorm.SubjectDao.New(tx, subject)
		}
	}

	tx.Commit()
}
