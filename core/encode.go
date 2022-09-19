package core

import "github.com/jxskiss/base62"

// Encode a db id into a short id
func EncodeID(id uint) string {
	raw := base62.FormatUint(uint64(id))
	return base62.EncodeToString(raw)
}

// Decode a short id into a db id
func DecodeID(encodedString string) (uint64, error) {
	raw, err := base62.DecodeString(encodedString)
	if err != nil {
		return 0, err
	}
	return base62.ParseUint(raw)
}
