package markerr_test

import (
	"errors"
	"fmt"

	"github.com/nametake/markerr"
)

func cause() error {
	return errors.New("cause")
}

func firstMarker() error {
	err := cause()
	err = markerr.Mark(err, "warning")

	// one line:
	// return fmt.Errorf("first: %w", markerr.Mark(err, "warning"))
	return fmt.Errorf("first: %w", err)
}

func secondMarker() error {
	err := firstMarker()
	return fmt.Errorf("second: %w", err)
}

func ExampleTakeMarker() {
	err := secondMarker()

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

func firstPair() error {
	err := cause()
	err = markerr.Pair(err, errors.New("sub"))

	// one line:
	// return fmt.Errorf("first: %w", markerr.Pair(err, errors.New("sub")))
	return fmt.Errorf("first: %w", err)
}

func secondPair() error {
	err := firstPair()
	return fmt.Errorf("second: %w", err)
}

func ExampleTakePair() {
	err := secondPair()

	fmt.Println(err)
	fmt.Println(markerr.TakePair(err))
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)
	err = errors.Unwrap(err)
	fmt.Println(err)

	// Output:
	// second: first: sub: cause
	// cause sub
	// first: sub: cause
	// sub: cause
	// cause
}
