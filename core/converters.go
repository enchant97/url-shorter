package core

import "strconv"

func Base10IntToString[T int | int8 | int16 | int32 | int64](v T) string {
	return strconv.FormatInt(int64(v), 10)
}

func Base10UIntToString[T uint | uint8 | uint16 | uint32 | uint64 | uintptr](v T) string {
	return strconv.FormatUint(uint64(v), 10)
}
