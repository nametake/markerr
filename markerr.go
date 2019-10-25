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

func Pair(main, sub error) error {
	return &pairErr{
		main: main,
		sub:  sub,
	}
}

func TakePair(err error) (error, error) {
	for err != nil {
		p, ok := err.(*pairErr)
		if ok {
			return p.main, p.sub
		}

		unwrap, ok := err.(unwrapper)
		if !ok {
			break
		}

		err = unwrap.Unwrap()
	}

	return nil, nil
}

type unwrapper interface {
	Unwrap() error
}

var (
	_ unwrapper = (*markErr)(nil)
	_ unwrapper = (*pairErr)(nil)
)

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

type pairErr struct {
	main, sub error
}

func (e *pairErr) Error() string {
	return fmt.Sprintf("%v: %v", e.sub, e.main)
}

func (e *pairErr) Unwrap() error {
	return e.main
}
