package db

func CreateNewShort(short Short) (Short, error) {
	db := DB.Create(&short)
	return short, db.Error
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
