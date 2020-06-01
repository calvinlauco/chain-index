package primptr

import (
	"math/big"
	"time"
)

// Shorthand functions to create primitive pointers

func Bool(value bool) *bool {
	return &value
}

func Int(value int) *int {
	return &value
}

func Int64(value int64) *int64 {
	return &value
}

func Uint8(value uint8) *uint8 {
	return &value
}

func Uint32(value uint32) *uint32 {
	return &value
}

func Uint64(value uint64) *uint64 {
	return &value
}

func String(value string) *string {
	return &value
}

func Time(value time.Time) *time.Time {
	return &value
}

func BigInt(value big.Int) *big.Int {
	return &value
}
