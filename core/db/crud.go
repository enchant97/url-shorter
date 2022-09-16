package db

func GetShortByShortID(shortID string) *Short {
	var shortRow Short
	if err := DB.Where("short_id = ?", shortID).First(&shortRow).Error; err != nil {
		return nil
	}
	return &shortRow
}

func (s *Short) Create() error {
	return DB.Create(&s).Error
}

// Records a new visitor & reduces uses left count
func (s *Short) IncrVisitCount() int {
	s.VisitCount++
	if s.UsesLeft != nil && *s.UsesLeft > 0 {
		*s.UsesLeft--
	}
	DB.Save(&s)
	return s.VisitCount
}
