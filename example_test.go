package markerr_test

import (
	"errors"
	"fmt"

	"github.com/nametake/markerr"
)

func cause() error {
	return errors.New("cause")
}

func first() error {
	err := cause()
	err = markerr.Mark(err, "warning")
	return fmt.Errorf("first: %w", err)

	// one line:
	// return fmt.Errorf("first: %w", markerr.Mark(err, "warning"))
}

func second() error {
	err := first()
	return fmt.Errorf("second: %w", err)
}

func Example() {
	err := second()

	fmt.Println(err)
	fmt.Println(markerr.Take(err))
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)

	// Output:
	// second: first: warning: cause
	// cause warning
	// first: warning: cause
	// warning: cause
	// cause
}
