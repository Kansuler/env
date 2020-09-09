# Kansuler/env

![License](https://img.shields.io/github/license/Kansuler/env) ![Version](https://img.shields.io/github/go-mod/go-version/Kansuler/env) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/693cfbf7964f457ead99202fc6d12679)](https://www.codacy.com/manual/Kansuler/env?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=Kansuler/env&amp;utm_campaign=Badge_Grade)

A small package that gets environment variables and convert them to their purposed basic type. If invalid, use the fallback value.

API and detailed documentation can be found at [https://godoc.org/github.com/Kansuler/env](https://godoc.org/github.com/Kansuler/env)

## Installation

`go get github.com/Kansuler/env`

## Functions

```go
Bool(key string, fallback bool) bool

String(key string, fallback string) string

Bytes(key string, fallback []byte) []byte

Runes(key string, fallback []rune) []rune

Float32(key string, fallback float32) float32

Float64(key string, fallback float64) float64

Int(key string, fallback int) int

Int8(key string, fallback int8) int8

Int16(key string, fallback int16) int16

Int32(key string, fallback int32) int32

Int64(key string, fallback int64) int64

Uint(key string, fallback uint) uint

Uint8(key string, fallback uint8) uint8

Uint16(key string, fallback uint16) uint16

Uint32(key string, fallback uint32) uint32

Uint64(key string, fallback uint64) uint64

Uintptr(key string, fallback uintptr) uintptr
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
