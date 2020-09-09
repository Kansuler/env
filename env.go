// env is a small package that parse environment variables to a basic type.
package env

import (
	"os"
	"strconv"
)

// Bool takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Bool(key string, fallback bool) bool {
	result, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return fallback
	}

	return result
}

// String takes an environment key, and a fallback value. Returns environment value if it isn't unset.
func String(key string, fallback string) string {
	if os.Getenv(key) == "" {
		return fallback
	}

	return os.Getenv(key)
}

// Int takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int(key string, fallback int) int {
	result, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}

	return result
}

// Int8 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int8(key string, fallback int8) int8 {
	result, err := strconv.ParseInt(os.Getenv(key), 10, 8)
	if err != nil {
		return fallback
	}

	return int8(result)
}

// Int16 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int16(key string, fallback int16) int16 {
	result, err := strconv.ParseInt(os.Getenv(key), 10, 16)
	if err != nil {
		return fallback
	}

	return int16(result)
}

// Int32 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int32(key string, fallback int32) int32 {
	result, err := strconv.ParseInt(os.Getenv(key), 10, 32)
	if err != nil {
		return fallback
	}

	return int32(result)
}

// Int64 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Int64(key string, fallback int64) int64 {
	result, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		return fallback
	}

	return result
}

// Uint takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint(key string, fallback uint) uint {
	result, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return fallback
	}

	return uint(result)
}

// Uint8 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint8(key string, fallback uint8) uint8 {
	result, err := strconv.ParseUint(os.Getenv(key), 10, 8)
	if err != nil {
		return fallback
	}

	return uint8(result)
}

// Uint16 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint16(key string, fallback uint16) uint16 {
	result, err := strconv.ParseUint(os.Getenv(key), 10, 16)
	if err != nil {
		return fallback
	}

	return uint16(result)
}

// Uint32 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint32(key string, fallback uint32) uint32 {
	result, err := strconv.ParseUint(os.Getenv(key), 10, 32)
	if err != nil {
		return fallback
	}

	return uint32(result)
}

// Uint64 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uint64(key string, fallback uint64) uint64 {
	result, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return fallback
	}

	return result
}

// Uintptr takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Uintptr(key string, fallback uintptr) uintptr {
	result, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		return fallback
	}

	return uintptr(result)
}

// Bytes takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Bytes(key string, fallback []byte) []byte {
	if os.Getenv(key) == "" {
		return fallback
	}

	return []byte(os.Getenv(key))
}

// Runes takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Runes(key string, fallback []rune) []rune {
	if os.Getenv(key) == "" {
		return fallback
	}

	return []rune(os.Getenv(key))
}

// Float32 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Float32(key string, fallback float32) float32 {
	result, err := strconv.ParseFloat(os.Getenv(key), 32)
	if err != nil {
		return fallback
	}

	return float32(result)
}

// Float64 takes an environment key, and a fallback value. Returns environment variable with converted type, or fallback
// value if it fails.
func Float64(key string, fallback float64) float64 {
	result, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return fallback
	}

	return result
}
