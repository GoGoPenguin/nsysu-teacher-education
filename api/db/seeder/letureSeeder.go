package main

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
)

func letureSeeder() {
	tx := gorm.DB().Begin()

	letureData := []map[string]interface{}{
		map[string]interface{}{"ID": uint(1), "Name": "自然科學領域化學專長", "MinCredit": uint(42), "Comment": "適合化學系所等"},
		map[string]interface{}{"ID": uint(2), "Name": "藝術領域藝術生活科－表演藝術專長", "MinCredit": uint(38), "Comment": "適合劇場藝術學系所等"},
		map[string]interface{}{"ID": uint(3), "Name": "藝術領域表演藝術專長", "MinCredit": uint(44), "Comment": "適合劇場藝術學系所等"},
		map[string]interface{}{"ID": uint(4), "Name": "語文領域國語文專長", "MinCredit": uint(48), "Comment": "適合中國文學系所等"},
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

		map[string]interface{}{"ID": uint(6), "LetureID": uint(3), "Name": "領域核心", "MinCredit": uint(3), "MinType": uint(0)},
		map[string]interface{}{"ID": uint(7), "LetureID": uint(3), "Name": "表演藝術專長課程", "MinCredit": uint(34), "MinType": uint(0)},

		map[string]interface{}{"ID": uint(8), "LetureID": uint(4), "Name": "主修專長課程", "MinCredit": uint(0), "MinType": uint(0)},
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

		map[string]interface{}{"ID": uint(13), "LetureCategoryID": uint(6), "Name": "藝術領域核心", "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(14), "LetureCategoryID": uint(7), "Name": "理解與應用", "MinCredit": uint(12)},
		map[string]interface{}{"ID": uint(15), "LetureCategoryID": uint(7), "Name": "實踐與展現", "MinCredit": uint(20)},
		map[string]interface{}{"ID": uint(16), "LetureCategoryID": uint(7), "Name": "教學知能", "MinCredit": uint(2)},

		map[string]interface{}{"ID": uint(17), "LetureCategoryID": uint(8), "Name": "語言知能課群", "MinCredit": uint(10)},
		map[string]interface{}{"ID": uint(18), "LetureCategoryID": uint(8), "Name": "文學知能課群", "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(19), "LetureCategoryID": uint(8), "Name": "哲學知能課群", "MinCredit": uint(7)},
		map[string]interface{}{"ID": uint(20), "LetureCategoryID": uint(8), "Name": "國學知能課群", "MinCredit": uint(7)},
		map[string]interface{}{"ID": uint(21), "LetureCategoryID": uint(8), "Name": "語文應用、創作、傳播與相關教學知能課群", "MinCredit": uint(6)},
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

		map[string]interface{}{"ID": uint(14), "LetureTypeID": uint(13), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(15), "LetureTypeID": uint(14), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(16), "LetureTypeID": uint(15), "MinCredit": uint(0)},
		map[string]interface{}{"ID": uint(17), "LetureTypeID": uint(16), "MinCredit": uint(0)},

		map[string]interface{}{"ID": uint(18), "LetureTypeID": uint(17), "MinCredit": uint(5)},
		map[string]interface{}{"ID": uint(19), "LetureTypeID": uint(17), "MinCredit": uint(2)},
		map[string]interface{}{"ID": uint(20), "LetureTypeID": uint(18), "MinCredit": uint(10)},
		map[string]interface{}{"ID": uint(21), "LetureTypeID": uint(18), "MinCredit": uint(8)},
		map[string]interface{}{"ID": uint(22), "LetureTypeID": uint(19), "MinCredit": uint(3)},
		map[string]interface{}{"ID": uint(23), "LetureTypeID": uint(19), "MinCredit": uint(2)},
		map[string]interface{}{"ID": uint(24), "LetureTypeID": uint(20), "MinCredit": uint(2)},
		map[string]interface{}{"ID": uint(25), "LetureTypeID": uint(20), "MinCredit": uint(2)},
		map[string]interface{}{"ID": uint(26), "LetureTypeID": uint(21), "MinCredit": uint(0)},
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

		map[string]interface{}{"SubjectGroupID": uint(14), "Name": "藝術概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(14), "Name": "劇場美學", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "表演技藝導論", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "劇場設計導論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "西洋戲劇史(一)", "Credit": uint(2), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "西洋戲劇史(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "中國戲劇史(一)", "Credit": uint(2), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "中國戲劇史(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "中國戲劇史(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "西洋藝術史", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "劇本導讀", "Credit": uint(2), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "劇本解析", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "臺灣劇場", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "現當代華語文戲劇選讀", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(15), "Name": "臺灣傳統戲曲", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "劇場製作基礎", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "劇場製作基礎", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "導演概論", "Credit": uint(2), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "導演(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "基礎表演", "Credit": uint(3), "Compulsory": true},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "傳統戲曲表演身段(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "觀點技巧表演", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "基礎發聲", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "聲音訓練", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "歌唱技巧", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "肢體開發(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "肢體開發(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "舞蹈技巧", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "舞蹈創作", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "排演", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "燈光設計及技術", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "劇場服裝技術", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "人物服裝畫", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "劇場服裝設計", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "舞臺化妝與造型", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "基礎設計", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "進階技術繪圖", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(16), "Name": "模型製作", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(17), "Name": "創作性戲劇", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(17), "Name": "戲劇治療與PBT設計", "Credit": uint(3), "Compulsory": false},

		map[string]interface{}{"SubjectGroupID": uint(18), "Name": "語言學概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(18), "Name": "文字學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(18), "Name": "聲韻學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(18), "Name": "訓詁學", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(19), "Name": "現代漢語語法", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(19), "Name": "古代漢語語法", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(19), "Name": "語意學概論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(19), "Name": "語言教學概論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(19), "Name": "閩南語讀書音", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(19), "Name": "閩南語概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(19), "Name": "客家話概論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "文學概論(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "文學概論(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "中國文學史(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "中國文學史(二)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "歷代文選及習作(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "歷代文選及習作(二)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "詩選及習作(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "詩選及習作(二)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "詞曲選及習作(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(20), "Name": "詞曲選及習作(二)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "詩經", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "楚辭", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "昭明文選", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "杜甫詩", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "王安石詩", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "韓柳文", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "東坡詞", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "明清小品選", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "歐蘇文", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "陶淵明詩文選讀", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "古典小說選", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "古典戲劇選", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "紅樓夢", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "文心雕龍", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "中國文學批評史", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "臺灣文學作品選讀", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "臺灣古典文學概論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "詩品", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(21), "Name": "中國五四時期的文本文化", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(22), "Name": "中國思想史(一)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(22), "Name": "中國思想史(二)", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "老子", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "莊子", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "荀子", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "墨子", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "韓非子", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "先秦子學思想綜論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "佛學概論", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "傳習錄", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(23), "Name": "近思錄", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(24), "Name": "國學導讀(一)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(24), "Name": "國學導讀(二)", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(25), "Name": "易經", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(25), "Name": "左傳", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(25), "Name": "論語", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(25), "Name": "孟子", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(25), "Name": "禮記", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(25), "Name": "史記", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(25), "Name": "韓詩外傳", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "應用文習作", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "現代詩選與寫作", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "現代散文選與寫作", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "現代小說選與寫作", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "兒童文學", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "紀錄片文獻選讀", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "現代文學理論", "Credit": uint(2), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "影像與文學", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "文學閱讀與生命書寫", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "地方文化典藏與報導應用", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "閩南民間文學與文化采風", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "臺灣語言踏查之旅", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "語言風格與創作", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "古典文化與現代生活", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "臺灣文化民俗誌", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "歲時節慶", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "海洋文化民俗誌", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "性別文化民俗誌", "Credit": uint(3), "Compulsory": false},
		map[string]interface{}{"SubjectGroupID": uint(26), "Name": "生命禮儀", "Credit": uint(3), "Compulsory": false},
	}

	for _, data := range subjectData {
		subject := &gorm.Subject{
			SubjectGroupID: data["SubjectGroupID"].(uint),
			Name:           data["Name"].(string),
			Credit:         data["Credit"].(uint),
			Compulsory:     data["Compulsory"].(bool),
		}

		if gorm.SubjectDao.GetByNameAndGroup(tx, subject.Name, subject.SubjectGroupID) == nil {
			gorm.SubjectDao.New(tx, subject)
		}
	}

	tx.Commit()
}
