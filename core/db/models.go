package db

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	BaseModel
	Username       string  `gorm:"unique;not null"`
	HashedPassword []byte  `gorm:"not null" json:"-"`
	OwnedShorts    []Short `gorm:"foreignkey:OwnerID;references:id" json:"-"`
}

// Set a new password (hashing it)
func (u *User) SetPassword(newPlainPassword string) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(newPlainPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.HashedPassword = hashedPw
}

// Check if password matches the hashed stored one
func (u *User) IsPasswordMatch(plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(plainPassword)); err == nil {
		return true
	}
	return false
}

type Short struct {
	BaseModel
	TargetURL  string     `gorm:"not null" json:"targetUrl"`
	VisitCount uint       `gorm:"default:0;not null" json:"visitCount,omitempty"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty"`
	MaxUses    *uint      `json:"maxUses,omitempty"`
	OwnerID    *uint      `json:"ownerId,omitempty"`
}

// Whether the short expiry has elapsed
func (s *Short) IsExpired() bool {
	if s.ExpiresAt == nil {
		return false
	}
	return s.ExpiresAt.Before(time.Now())
}

// Whether the short can be used (a redirect)
func (s *Short) IsUsable() bool {
	if s.MaxUses == nil {
		return true
	}
	return *s.MaxUses > s.VisitCount
}
