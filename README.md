# Overview

Welcome to the `fstr` repository!

`fstr` is a super simple package intended to make very complex string formatting
scenarios a bit more readable than the Go standard library can provide while
also keeping as close as possible to the standard lib's `Sprintf` performance.

`fstr` allows string formatting via brace-enclosed tokens, commonly seen in
other languages such as Python. The values to interpolate in place of these
tokens are provided as either a map (`fstr.Map`) or as simply groups of two in
variadic form (`fstr.Pairs`).

# Quick Start

To install `fstr` simply run:

```shell
go get github.com/illbjorn/fstr
```

# Examples

## Interpolating with Maps

Using `fstr.Map` we can format a string by replacing the brace-enclosed tokens
in our input with the matching key-value from an input map.

Example:

```go
package main

import "github.com/illbjorn/fstr"
import "fmt"

func main() {
  res := fstr.Map(
    "The quick brown {animal} {verb} over the lazy dog.",
    map[string]string{
      "animal": "fox",
      "verb": "jumped",
    },
  )

  fmt.Println(res) // The quick brown fox jumped over the lazy dog.
}
```

> [!NOTE]
> If you want an actual brace included in your output, you can escape it with a
> backslash (`\`) immediately preceding the open brace.

## Interpolating with Variadic String Pairs

Using `fstr.Pairs` we can format a string by replacing the brace-enclosed tokens
in our input with the matching-key-plus-one member in a variadic string input.

Example:

```go
package main

import "github.com/illbjorn/fstr"
import "fmt"

func main() {
  res := fstr.Pairs(
    "The quick brown {animal} {verb} over the lazy dog.",
    "animal", "fox",
    "verb", "jumped",
    },
  )

  fmt.Println(res) // The quick brown fox jumped over the lazy dog.
}
```

# Performance

For the purpose of this string formatting, I don't know it will be super
attainable to be equivalent or faster than `fmt.Sprintf`. But, `fstr` is not
terribly far behind:

```
goos: windows
goarch: amd64
pkg: github.com/illbjorn/fstr
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkMap-24        	 3386107	       373.7 ns/op	     320 B/op	       3 allocs/op
BenchmarkSprintf-24    	 3468406	       338.8 ns/op	     152 B/op	       5 allocs/op
BenchmarkPairs-24      	 3430071	       350.5 ns/op	     320 B/op	       3 allocs/op
```

This benchmark tests:

| Call          | Input                                                    | Values                                                             |
| ------------- | -------------------------------------------------------- | ------------------------------------------------------------------ |
| `fstr.Pairs`  | `"The {val1} brown {val2} jumped over the lazy {val3}."` | `[]string{"val1", "quick", "val2", "fox", "val3", "dog"}`          |
| `fstr.Map`    | `"The {val1} brown {val2} jumped over the lazy {val3}."` | `map[string]string{"val1": "quick", "val2": "fox", "val3": "dog"}` |
| `fmt.Sprintf` | `"The %s brown %s jumped over the lazy %s."`             | `[]string{"quick", "fox", "dog"}`                                  |
