package main

import (
	"runtime"
	"testing"
)

func BenchmarkParseInt(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := UseParseInt("123456", 10, 0)
		if err != nil {
			b.Fatalf("Something went wrong: %v", err)
		}
		res2 = res
	}
}

func BenchmarkAtoi(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := UseAtoi("123456")
		if err != nil {
			b.Fatalf("Something went wrong: %v", err)
		}
		res1 = res
	}
}
