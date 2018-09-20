package strconv

import (
	"strconv"
	"testing"
)

var res string

// UseFormatInt converts integer to string using FormatInt
func UseFormatInt(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

// UseItoa converts integer to string using Itoa
func UseItoa(i int) string {
	return strconv.Itoa(i)
}

func BenchmarkUseFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = UseFormatInt(42, 10)
	}
}

func BenchmarkUseItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = UseItoa(42)
	}
}
