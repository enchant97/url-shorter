package core

import (
	"github.com/enchant97/url-shorter/core/db"
	"github.com/thanhpk/randstr"
)

// How many characters long the short id will be
const ShortIDLength = 8

// Make a new short id
func MakeShortID() string {
	return randstr.String(ShortIDLength)
}

func (s *CreateShort) GenerateShort() db.Short {
	return db.Short{
		TargetURL: s.TargetURL,
		ShortID:   MakeShortID(),
	}
}
