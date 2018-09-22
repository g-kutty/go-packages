package main

import (
	"fmt"
	"strconv"
)

var (
	res1 int
	res2 int64
)

func main() {
	UseParseInt("42", 10, 0)
	UseAtoi("42")
}

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
