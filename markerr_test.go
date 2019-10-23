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
func TestTake(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		want   string
		marker string
	}{
		{
			name: "test name",
			err: fmt.Errorf("second: %w",
				fmt.Errorf("first: %w",
					Mark(errors.New("cause"), "marker"),
				),
			),
			want:   "cause",
			marker: "marker",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err, marker := Take(test.err)
			if err.Error() != test.want {
				t.Errorf("got: %v, want %v", err.Error(), test.want)
			}
			if marker != test.marker {
				t.Errorf("got: %v, want %v", marker, test.marker)
			}
		})
	}
}
