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

func Take(err error) (error, string) {
	for err != nil {
		m, ok := err.(*markErr)
		if ok {
			return m.err, m.marker
		}
		unwrap, ok := err.(unwrapper)
		if !ok {
			break
		}
		err = unwrap.Unwrap()
	}
	return nil, ""
}

type unwrapper interface {
	Unwrap() error
}
