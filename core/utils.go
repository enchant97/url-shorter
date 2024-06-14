package core

import (
	"crypto/rand"
	"github.com/jxskiss/base62"
)

func GenerateRandomSlug(length uint) string {
	rBytes := make([]byte, length)
	if _, err := rand.Read(rBytes); err != nil {
		panic(err)
	}
	return base62.EncodeToString(rBytes)[:length]
}
