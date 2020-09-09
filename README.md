# Kansuler/env

![License](https://img.shields.io/github/license/Kansuler/env) ![Version](https://img.shields.io/github/go-mod/go-version/Kansuler/env) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/693cfbf7964f457ead99202fc6d12679)](https://www.codacy.com/manual/Kansuler/env?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Kansuler/env&amp;utm_campaign=Badge_Grade)

A small package that gets environment variables and convert them to their purposed basic type. If invalid, use the fallback value.

API and detailed documentation can be found at [https://godoc.org/github.com/Kansuler/env](https://godoc.org/github.com/Kansuler/env)

## Installation

`go get github.com/Kansuler/env`

## Functions

```go
func Bool(key string, fallback bool) bool

func String(key string, fallback string) string

func Bytes(key string, fallback []byte) []byte

func Runes(key string, fallback []rune) []rune

func Float32(key string, fallback float32) float32

func Float64(key string, fallback float64) float64

func Int(key string, fallback int) int

func Int8(key string, fallback int8) int8

func Int16(key string, fallback int16) int16

func Int32(key string, fallback int32) int32

func Int64(key string, fallback int64) int64

func Uint(key string, fallback uint) uint

func Uint8(key string, fallback uint8) uint8

func Uint16(key string, fallback uint16) uint16

func Uint32(key string, fallback uint32) uint32

func Uint64(key string, fallback uint64) uint64

func Uintptr(key string, fallback uintptr) uintptr
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
