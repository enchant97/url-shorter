package db

import "github.com/enchant97/url-shorter/core"

func CreateNewShort(short core.Short) (Short, error) {
	shortRow := Short{ShortID: short.ShortID, TargetURL: short.TargetURL}
	db := DB.Create(&shortRow)
	return shortRow, db.Error
}

func GetShortByShortID(shortID string) Short {
	var shortRow Short
	if err := DB.Where("short_id = ?", shortID).First(&shortRow).Error; err != nil {
		return Short{}
	}
	return shortRow
}

func (s *Short) IncrVisitCount() int {
	s.VisitCount++
	DB.Save(&s)
	return s.VisitCount
}
