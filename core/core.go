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

func NullableIsoStringToTime(input *string) (*time.Time, error) {
	if input == nil || *input == "" {
		return nil, nil
	}
	expiresAt, err := time.Parse("2006-01-02T15:04", *input)
	return &expiresAt, err
}

func (s *CreateShort) GenerateShort() db.Short {
	expiresAt, err := NullableIsoStringToTime(s.ExpiresAt)
	if err != nil {
		panic("time parse error")
	}
	return db.Short{
		TargetURL: s.TargetURL,
		ShortID:   MakeShortID(),
		ExpiresAt: expiresAt,
	}
}
