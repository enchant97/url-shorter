package core

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jxskiss/base62"
)

const (
	padSep   = "-"
	padEvery = 6
)

var allowedShortIDRegex = regexp.MustCompile(`^[a-zA-Z0-9,-]+$`)

// Insert a separator every n character
func padSepString(s string, insertEvery int, sep string) string {
	if len(s) <= insertEvery {
		// Skip if separator not needed
		return s
	}
	for i := insertEvery; i < len(s); i += (insertEvery + len(sep)) {
		s = s[:i] + sep + s[i:]
	}
	return s
}

// Check whether given id is a valid ShortID (with or without padding)
func IsValidShortID(shortID string) bool {
	return allowedShortIDRegex.MatchString(shortID)
}

// Encode a db id into a short id
func EncodeID(id uint) string {
	raw := base62.FormatUint(uint64(id))
	return base62.EncodeToString(raw)
}

// Decode a short id into a db id
func DecodeID(encodedID string) (uint64, error) {
	raw, err := base62.DecodeString(encodedID)
	if err != nil {
		return 0, err
	}
	return base62.ParseUint(raw)
}

// Encode a db id into a short id with human friendly padding
func EncodeIDPadded(id uint) string {
	return padSepString(EncodeID(id), padEvery, padSep)
}

// Decode a short id with human friendly padding into a db id
func DecodeIDPadded(encodedID string) (uint64, error) {
	encodedID = strings.ReplaceAll(encodedID, padSep, "")
	return DecodeID(encodedID)
}

// Try to decode a short id from user input into a db id,
// allowing for human friendly padding
func DecodePossibleShortID(shortID string) (uint64, error) {
	if !IsValidShortID(shortID) {
		return 0, fmt.Errorf("does not match accepted regex")
	}
	decodedID, err := DecodeIDPadded(shortID)
	if err != nil {
		return 0, err
	}
	return decodedID, nil
}
