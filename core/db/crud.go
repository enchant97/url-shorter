package db

func CreateNewShort(shortID string, targetURL string) Short {
	shortRow := Short{ShortID: shortID, TargetURL: targetURL}
	DB.Create(&shortRow)
	return shortRow
}

func GetShortByShortID(shortID string) Short {
	var shortRow Short
	if err := DB.Where("short_id = ?", shortID).First(&shortRow).Error; err != nil {
		return Short{}
	}
	return shortRow
}

func ShortVisitCountIncr(shortID string) int {
	var shortRow Short
	if err := DB.Where("short_id = ?", shortID).First(&shortRow).Error; err == nil {
		shortRow.VisitCount++
		DB.Save(&shortRow)
		return shortRow.VisitCount
	}
	return -1
}
