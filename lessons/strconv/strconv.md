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

### ParseInt
***

* ParseInt interprets a string `s` in the given `base` (0, 2 to 36) and `bitSize` (0 to 64) and returns the corresponding value i.

* If base and bitSize is not in respective range it will return an error.

* [Go Playground](https://play.golang.org/p/bimUtt4T5YH)

```go
    func ParseUint(s string, base int, bitSize int) (uint64, error)
```