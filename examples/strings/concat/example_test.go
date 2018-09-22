package main

import (
	"runtime"
	"testing"
)

func BenchmarkUsePlusSign(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		UsePlusSign("s", 1000)
	}
}

func BenchmarkUseSlice(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		UseSlice("s", 1000)
	}
}

func BenchmarkUseBytesBuffer(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		UseBytesBuffer("s", 1000)
	}
}

func BenchmarkUseStringsBuilder(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		UseStringsBuilder("s", 1000)
	}
}
