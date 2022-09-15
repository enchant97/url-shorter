package db

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Short struct {
	BaseModel
	ShortID    string     `gorm:"unique;not null" json:"shortId"`
	TargetURL  string     `gorm:"not null" json:"targetUrl"`
	VisitCount int        `gorm:"default:0;not null" json:"visitCount,omitempty"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty"`
}

func (s *Short) IsExpired() bool {
	if s.ExpiresAt == nil {
		return false
	}
	return s.ExpiresAt.Before(time.Now())
}
