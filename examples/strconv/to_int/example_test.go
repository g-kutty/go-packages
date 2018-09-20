package strconv

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

var (
	res1 int
	res2 int64
)

// UseParseInt convert string to int using parseInt
func UseParseInt(s string, base, bitSize int) (int64, error) {
	res, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// UseAtoi convert string to int using Atoi
func UseAtoi(s string) (int, error) {
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

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
