package main

import "strconv"

func main() {
	UseFormatInt(42, 10)
	UseItoa(42)
}

// UseFormatInt converts integer to string using FormatInt
func UseFormatInt(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

// UseItoa converts integer to string using Itoa
func UseItoa(i int) string {
	return strconv.Itoa(i)
}
