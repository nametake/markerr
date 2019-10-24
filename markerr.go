package markerr

import "fmt"

func Mark(err error, marker string) error {
	if err == nil {
		return nil
	}

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

func (e *markErr) Unwrap() error {
	return e.err
}

func TakeMarker(err error) (string, error) {
	for err != nil {
		m, ok := err.(*markErr)
		if ok {
			return m.marker, m.err
		}

		unwrap, ok := err.(unwrapper)
		if !ok {
			break
		}

		err = unwrap.Unwrap()
	}

	return "", nil
}

type unwrapper interface {
	Unwrap() error
}
