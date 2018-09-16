package strings

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

// UseParseInt convert string to int using parseInt
func UseParseInt(s string) (int64, error) {
	res, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func BenchmarkParseInt(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := UseParseInt("123456")
		if err != nil {
			b.Fatalf("Something went wrong: %v", err)
		}
	}
}

// UseAtoi convert string to int using Atoi
func UseAtoi(s string) (int, error) {
	res, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}

func BenchmarkAtoi(b *testing.B) {
	// reset befor start
	runtime.GC()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := UseAtoi("123456")
		if err != nil {
			b.Fatalf("Something went wrong: %v", err)
		}
	}
}
