markerr
=======

[![GoDoc](https://godoc.org/github.com/nametake/suberr?status.svg)](https://godoc.org/github.com/nametake/suberr)

Install
-------

`go get github.com/nametake/markerr`

Usage
-----

```go
package main

import (
	"errors"
	"fmt"

	"github.com/nametake/markerr"
)

func main() {
	err := errors.New("cause")
	err = markerr.Mark(err, "warning")
	err = fmt.Errorf("first: %w", err)
	err = fmt.Errorf("second: %w", err)

	fmt.Println(err)               // second: first: warning: cause
	fmt.Println(markerr.Take(err)) // cause warning
	err = errors.Unwrap(err)
	fmt.Println(err) // first: warning: cause
	err = errors.Unwrap(err)
	fmt.Println(err) // warning: cause
	err = errors.Unwrap(err)
	fmt.Println(err) // cause
}
```
