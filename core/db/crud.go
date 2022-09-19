package db

func GetShortByID(id uint) *Short {
	var shortRow Short
	if err := DB.Where("id = ?", id).First(&shortRow).Error; err != nil {
		return nil
	}
	return &shortRow
}

func (s *Short) Create() error {
	return DB.Create(&s).Error
}

// Records a new visitor
func (s *Short) IncrVisitCount() (uint, error) {
	s.VisitCount++
	err := DB.Save(&s).Error
	return s.VisitCount, err
}

func GetUserByUsername(username string) *User {
	var userRow User
	if err := DB.Where("username = ?", username).First(&userRow).Error; err != nil {
		return nil
	}
	return &userRow
}

func GetUserByID(userID uint) *User {
	var userRow User
	if err := DB.Where("id = ?", userID).First(&userRow).Error; err != nil {
		return nil
	}
	return &userRow
}
