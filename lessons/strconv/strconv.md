---
layout: page
title: strconv
---
***

Package __strconv__ implements conversions to and from string representations of basic data types.

&nbsp;

## String to Int
***

* A string contains digit characters like `123`. We can convert this string to a number with using methods __ParseInt__, __Atoi__.

&nbsp;

### ➩ [ParseInt](https://play.golang.org/p/b7eud5F9nQ0)
***

* ParseInt interprets a string `s` in the given `base` (0, 2 to 36) and `bitSize` (0 to 64) and returns the corresponding value i.

* If base and bitSize is not in respective range it will return an error.

    ```go
        func ParseInt(s string, base int, bitSize int) (i int64, err error)
    ```

&nbsp;

### ➩ [Atoi](https://play.golang.org/p/rI_TaO5bOig)
***

* This function bears the same name as the one from the `C` standard library. 

* Atoi returns the result of `ParseInt(s, 10, 0)` converted to type `int`.

    ```go
        func Atoi(s string) (int, error)
    ```
&nbsp;

### ➩ [Benchmark test](https://play.golang.org/p/gRZE4VYq-jv)

* Often we need to convert a string into an integer. This benchmark compares __ParseInt__ and __Atoi__.

* This test based on go version `go version go1.11 linux/amd64`

    ```sh
        goos: linux
        goarch: amd64
        pkg: github.com/g-kutty/go-packages/examples/strconv
        BenchmarkParseInt-4   	100000000	        34.9 ns/op	       0 B/op	       0 allocs/op
        BenchmarkAtoi-4       	200000000	        19.0 ns/op	       0 B/op	       0 allocs/op
        PASS
        ok  	github.com/g-kutty/go-packages/examples/strconv	9.273s

    ```