[![Go Report Card](https://goreportcard.com/badge/github.com/mamal72/slicy)](https://goreportcard.com/report/github.com/mamal72/slicy)
[![GoDoc](https://godoc.org/github.com/mamal72/slicy?status.svg)](https://godoc.org/github.com/mamal72/slicy)
[![license](https://img.shields.io/github/license/mamal72/slicy.svg)](https://github.com/mamal72/slicy/blob/master/LICENSE)


# [ ]Slicy

Slicy is a set of *typesafe* *chainable* slice helpers that help you work with slices easier.


## Installation

```bash
go get github.com/mamal72/slicy
# or use dep, vgo, glide or anything else
```


## Usage

### Supported Types

- `byte`
- `int`
- `int8`
- `int16`
- `int32`
- `int64`
- `float32`
- `float64`
- `uint`
- `uint8`
- `uint16`
- `uint32`
- `uint64`
- `string`
- `interface{}`


### Implemented Methods

- `Filter`
- `Map`
- `Push`
- `Pop`
- `Shift`
- `Unshift`
- `Append`
- `Concat`
- `Every`
- `Some`
- `Includes`
- `Len`


### Example

Here is a simple example of string slice:

```go
s := slicy.NewFromStringSlice([]string{"my", "string", "slice", "no!"})
s.Filter(func(item string) {
    return item != "no!"
})
// You may call other methods here...

result := s.GetSlice() // returns a *[]string containing {"my", "string", "slice"}
```


## Build

This package is generated using [genny](https://github.com/cheekybits/genny). You can simply run `go generate` or `make generate` to build project.


## Ideas || Issues

Just create an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back if you want to see them here. I really appreciate that. ❤️
