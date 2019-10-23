package markerr

import "fmt"

func Mark(err error, marker string) error {
	return &markErr{
		err:    err,
		marker: marker,
	}
}

type markErr struct {
	err    error
	marker string
}

func (e *markErr) Error() string {
	return fmt.Sprintf("%s: %v", e.marker, e.err)
}

func Take(err error) (string, error) {
	panic("not implemented")
}
