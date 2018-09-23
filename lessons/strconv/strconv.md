---
layout: page
title: strconv
---
***

Package __strconv__ implements conversions to and from __string__ representations of basic data types.

&nbsp;

## String to Integer
***

* A string contains digit characters like `"123"`. We can convert this string to a number with using methods __ParseInt__, __Atoi__.

&nbsp;

### ➩ [ParseInt](https://play.golang.org/p/b7eud5F9nQ0)
***

* ParseInt interprets a string `s` in the given `base` (0, 2 to 36) and `bitSize` (0 to 64) and returns the corresponding value i.

* If base and bitSize is not in their respective range it will return an `error`.

    ```go
    func ParseInt(s string, base, bitSize int) (int64, error)
    ```

&nbsp;

### ➩ [Atoi](https://play.golang.org/p/rI_TaO5bOig)
***

* This function bears the same name as the one from the `C` standard library.

* __Atoi__ returns the result of __ParseInt(s, 10, 0)__ converted to type `int`.

    ```go
    func Atoi(s string) (int, error)
    ```
&nbsp;

### ➩ [Benchmark](https://github.com/g-kutty/go-packages/blob/master/examples/strconv/to_int/example_test.go)

* Often we need to convert a string into an integer. This benchmark compares __ParseInt__ and __Atoi__.

* This test based on go version go version `go1.11 linux/amd64`

    ```sh
    goos: linux
    goarch: amd64
    BenchmarkParseInt-4   100000000        34.9 ns/op       0 B/op       0 allocs/op
    BenchmarkAtoi-4       200000000        19.0 ns/op       0 B/op       0 allocs/op
    ```

&nbsp;

## Interger to String
***

* A Integer contains only digit like `123`. We can convert this integer to a string with using methods __FormatInt__, __Itoa__.

### ➩ [FormatInt](https://play.golang.org/p/WKsxy73aDdw)
***

* __FormatInt__ returns the string representation of `i` in the given `base`, for 2 <= base <= 36.

* The result uses the lower-case letters `a` to `z` for digit values >= 10.

    ```go
    func FormatInt(i int64, base int) string
    ```
&nbsp;

### ➩ [Itoa](https://play.golang.org/p/BlKb8Q6_xjv)
***

* __Itoa__ is shorthand for __FormatInt(int64(i), 10)__.

    ```go
    func Itoa(i int) string
    ```
&nbsp;

### ➩ [Benchmark](https://github.com/g-kutty/go-packages/blob/master/examples/strconv/to_string/example_test.go)

* This test based on go version go version `go1.11 linux/amd64`

    ```sh
    goos: linux
    goarch: amd64
    BenchmarkUseFormatInt-4   500000000        10.3 ns/op       0 B/op       0 allocs/op
    BenchmarkUseItoa-4        300000000        13.5 ns/op       0 B/op       0 allocs/op
    ```