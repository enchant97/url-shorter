package core

import (
	"time"

	"github.com/enchant97/url-shorter/core/db"
	"github.com/thanhpk/randstr"
)

// How many characters long the short id will be
const ShortIDLength = 8

// Make a new short id
func MakeShortID() string {
	return randstr.String(ShortIDLength)
}

// Convert given time to human readable or use default string if nil
func TimeToHumanOr(inputTime *time.Time, nilDefault string) string {
	if inputTime == nil {
		return nilDefault
	}
	return inputTime.Format("2006-01-02 15:04")
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
		MaxUses:   maxUses,
	}
}
