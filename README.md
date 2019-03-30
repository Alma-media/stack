# stack

[![GoDoc][godoc-badge]][godoc-link]
[![License][license-badge]][license-link]
[![Build Status][circleci-badge]][circleci-link]
[![Report Card][report-badge]][report-link]
[![GoCover][cover-badge]][cover-link]

Simple Golang stack implemented through a linked list

### Installation
```bash
go get -u github.com/tiny-go/stack
```

### Usage
```go
package main

import (
	"fmt"

	"github.com/tiny-go/stack"
)

func main() {
	var a, b, c = 1, 2, 3

	// init stack
	stk := stack.New()

	// push elements
	stk.Push(a)
	stk.Push(b)
	stk.Push(c)

	// get top element
	fmt.Println(stk.Top(), stk.Len()) // 3 3

	// pop element
	fmt.Println(stk.Pop(), stk.Pop(), stk.Pop(), stk.Len()) // 3 2 1 0
}

```

[godoc-badge]: https://godoc.org/github.com/tiny-go/stack?status.svg
[godoc-link]: https://godoc.org/github.com/tiny-go/stack
[license-badge]: https://img.shields.io/:license-MIT-green.svg
[license-link]: https://opensource.org/licenses/MIT
[circleci-badge]: https://circleci.com/gh/tiny-go/stack.svg?style=shield
[circleci-link]: https://circleci.com/gh/tiny-go/stack
[report-badge]: https://goreportcard.com/badge/github.com/tiny-go/stack
[report-link]: https://goreportcard.com/report/github.com/tiny-go/stack
[cover-badge]: https://gocover.io/_badge/github.com/tiny-go/stack
[cover-link]: https://gocover.io/github.com/tiny-go/stack
