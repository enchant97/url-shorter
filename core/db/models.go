package db

import "gorm.io/gorm"

type Short struct {
	gorm.Model
	ShortID    string `gorm:"unique"`
	TargetURL  string
	VisitCount int `gorm:"default:0"`
}
