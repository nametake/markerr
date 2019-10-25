package markerr

import (
	"errors"
	"fmt"
	"testing"
)

func TestMark(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		marker string
		want   string
	}{
		{
			name:   "simple",
			err:    errors.New("cause error"),
			marker: "marker",
			want:   "marker: cause error",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Mark(test.err, test.marker).Error(); got != test.want {
				t.Errorf("Mark.Error(): got: %v, want %v", got, test.want)
			}
		})
	}
}

func TestMarkNil(t *testing.T) {
	if got := Mark(nil, "no error"); got != nil {
		t.Errorf("Mark(nil, \"no error\"): got: %v, want nil", got)
	}
}

func TestPair(t *testing.T) {
	tests := []struct {
		name string
		main error
		sub  error
		want string
	}{
		{
			name: "simple",
			main: errors.New("main"),
			sub:  errors.New("sub"),
			want: "sub: main",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Pair(test.main, test.sub).Error(); got != test.want {
				t.Errorf("Pair.Error(): got: %v, want %v", got, test.want)
			}
		})
	}
}

func TestPairNil(t *testing.T) {
	tests := []struct {
		name string
		main error
		sub  error
	}{
		{
			name: "main nil",
			main: nil,
			sub:  errors.New("sub"),
		},
		{
			name: "sub nil",
			main: errors.New("main"),
			sub:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Pair(test.main, test.sub); got != nil {
				t.Errorf("Pair.Error(): got: %v, want nil", got)
			}
		})
	}
}

func TestTakeMarker(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		want   string
		marker string
	}{
		{
			name: "simple",
			err: fmt.Errorf("second: %w",
				fmt.Errorf("first: %w",
					Mark(errors.New("cause"), "marker"),
				),
			),
			want:   "cause",
			marker: "marker",
		},
		{
			name: "while fmt.Errorf",
			err: fmt.Errorf("third: %w",
				fmt.Errorf("second: %w",
					Mark(
						fmt.Errorf("first: %w", errors.New("cause")),
						"marker",
					),
				),
			),
			want:   "first: cause",
			marker: "marker",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			marker, err := TakeMarker(test.err)
			if err.Error() != test.want {
				t.Errorf("got: %v, want %v", err.Error(), test.want)
			}
			if marker != test.marker {
				t.Errorf("got: %v, want %v", marker, test.marker)
			}
		})
	}
}

func TestTakePair(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		want1 string
		want2 string
	}{
		{
			name: "simple",
			err: fmt.Errorf("second: %w",
				fmt.Errorf("first: %w",
					Pair(errors.New("main"), errors.New("sub")),
				),
			),
			want1: "main",
			want2: "sub",
		},
		{
			name: "while fmt.Errorf",
			err: fmt.Errorf("third: %w",
				fmt.Errorf("second: %w",
					Pair(
						fmt.Errorf("first: %w", errors.New("cause")),
						errors.New("sub"),
					),
				),
			),
			want1: "first: cause",
			want2: "sub",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want1, want2 := TakePair(test.err)
			if want1.Error() != test.want1 {
				t.Errorf("got: %v, want %v", want1, test.want1)
			}
			if want2.Error() != test.want2 {
				t.Errorf("got: %v, want %v", want2.Error(), test.want2)
			}
		})
	}
}
