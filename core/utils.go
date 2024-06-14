package core

import (
	"crypto/rand"
	"encoding/base64"
)

var RandomSlugEncoding = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func GenerateRandomSlug(length uint) string {
	rBytes := make([]byte, length)
	if _, err := rand.Read(rBytes); err != nil {
		panic(err)
	}
	return RandomSlugEncoding.EncodeToString(rBytes)[:length]
}
