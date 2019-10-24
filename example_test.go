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

	// one line:
	// return fmt.Errorf("first: %w", markerr.Mark(err, "warning"))
	return fmt.Errorf("first: %w", err)
}

func second() error {
	err := first()
	return fmt.Errorf("second: %w", err)
}

func Example() {
	err := second()

	fmt.Println(err)
	fmt.Println(markerr.TakeMarker(err))
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)

	// Output:
	// second: first: warning: cause
	// warning cause
	// first: warning: cause
	// warning: cause
	// cause
}
