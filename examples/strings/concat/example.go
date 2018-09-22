package main

import (
	"bytes"
	"strings"
)

var finalRes string

func main() {
	UsePlusSign("s", 1000)
	UseSlice("s", 1000)
	UseBytesBuffer("s", 1000)
	UseStringsBuilder("s", 1000)
}

// UsePlusSign perform concatenation of string using + sign
func UsePlusSign(s string, c int) {
	for i := 0; i < c; i++ {
		finalRes += s
	}
}

// UseSlice perfrom concatenation of string using slice
func UseSlice(s string, c int) {
	res := make([]string, 0, c)
	for i := 0; i < c; i++ {
		res = append(res, s)
	}
	finalRes = strings.Join(res, "")
}

// UseBytesBuffer perfrom concatenation of string using buffer
func UseBytesBuffer(s string, c int) {
	var buffer bytes.Buffer
	for i := 0; i < c; i++ {
		buffer.WriteString(s)
	}
	finalRes = buffer.String()
}

// UseStringsBuilder perform concatenation of string using builder
func UseStringsBuilder(s string, c int) {
	var sb strings.Builder

	for i := 0; i < c; i++ {
		sb.WriteString(s)
	}
	finalRes = sb.String()
}
