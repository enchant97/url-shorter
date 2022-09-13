package core

import "github.com/thanhpk/randstr"

// until real database is implemented
var FakeShortsDB = map[string]string{}

// How many characters long the short id will be
const ShortIDLength = 8

// Make a new short id
func MakeShortID() string {
	return randstr.String(ShortIDLength)
}
