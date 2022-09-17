package core

import (
	"time"

	"github.com/enchant97/url-shorter/core/db"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

// How many characters long the short id will be
const ShortIDLength = 8

// Make a new short id
func MakeShortID() string {
	return randstr.String(ShortIDLength)
}

func NullableIsoStringToTime(input *string) (*time.Time, error) {
	if input == nil || *input == "" {
		return nil, nil
	}
	expiresAt, err := time.Parse("2006-01-02T15:04", *input)
	return &expiresAt, err
}

func (s *CreateShort) GenerateShort() db.Short {
	// Put expire time in correct format
	expiresAt, err := NullableIsoStringToTime(s.ExpiresAt)
	if err != nil {
		panic("time parse error")
	}
	// Ensure max use lower than 1 is represented as nil
	var maxUses *uint
	if s.MaxUses != nil && *s.MaxUses < 1 {
		maxUses = nil
	} else {
		maxUses = s.MaxUses
	}
	return db.Short{
		TargetURL: s.TargetURL,
		ShortID:   MakeShortID(),
		ExpiresAt: expiresAt,
		UsesLeft:  maxUses,
	}
}

func GetAuthenticatedUserID(c *gin.Context) *uint {
	session := sessions.Default(c)
	if userID := session.Get("authenticatedUserID"); userID != nil {
		if userID, isValid := userID.(uint); isValid {
			return &userID
		}
	}
	return nil
}

func GetAuthenticatedUser(c *gin.Context) *db.User {
	if userID := GetAuthenticatedUserID(c); userID != nil {
		if userRow := db.GetUserByID(*userID); userRow != nil {
			return userRow
		}
	}
	return nil
}

func SetAuthenticatedUserID(c *gin.Context, userID uint) error {
	session := sessions.Default(c)
	session.Set("authenticatedUserID", userID)
	return session.Save()
}

func RemoveAuthenticatedUser(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	return session.Save()
}
