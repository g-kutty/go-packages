package main

import (
	"testing"
)

var res string

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
