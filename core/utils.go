package core

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomSlug(length uint) string {
	rBytes := make([]byte, length)
	if _, err := rand.Read(rBytes); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(rBytes)[:length]
}
