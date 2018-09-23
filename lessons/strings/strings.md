---
layout: page
title: strings
---
***

Package __strings__ implements simple functions to manipulate UTF-8 encoded strings.

&nbsp;

### [String Concatenation](https://play.golang.org/p/JnWlpMsluLO)
***

1 ) __Add__ them together

* Concatenating set of strings by manually add them to a string.

    ```go
    var ret string
    for i := 0; i < count; i++ {
        ret += str
    }
    ```

2 ) __strings.Join__

* We can join all strings in the slice together using `strings.Join` method.

    ```go
    var ret []string
    for i := 0; i < count; i++ {
        ret = append(ret, str)
    }
    res = strings.Join(ret,"")
    ```

3 ) __bytes.Buffer__

* By using `bytes.Buffer` type along with its `WriteString` and `String` method we can easily join strings together.

    ```go
    var buffer bytes.Buffer
    for i := 0; i < count; i++ {
        buffer.WriteString(str)
    }
    res = buffer.String()
    ```

4 ) __strings.Builder__

* Another most efficient way is, by using `strings.Builder` type along with its `WriteString` and `String` method.

* Its avialable from `go 1.10`.

    ```go
    var sb strings.Builder
    for i := 0; i < count; i++ {
        sb.WriteString(str)
    }
    res = sb.String()
    ```

&nbsp;

* First two are work perfectly fine for a simple program, It is a little inefficient.

* Every time we add or append string, a brand new string needs to be allocated in memory.

* This happens because strings in Go are immutable. so if we want to change a string or add content to it we need to create an entirely new string.

* To avoid creating new strings as we build our final string, we can now use the `strings.Builder` or `bytes.Buffer` type along with its `WriteString` method.

### Benchmark
***

* Most efficient way of concatenating strings together is using `strings.Builder` type.

* This test based on go version go version `go1.11 linux/amd64`

    ```sh
    goos: linux
    goarch: amd64
    BenchmarkUsePlusSign-4              300  75562500 ns/op    255076616 B/op    1000 allocs/op
    BenchmarkUseSlice-4              200000     31301 ns/op    18432 B/op       3 allocs/op
    BenchmarkUseBytesBuffer-4        200000     20065 ns/op    3296 B/op       6 allocs/op
    BenchmarkUseStringsBuilder-4     500000      9389 ns/op    2040 B/op       8 allocs/op
    ```