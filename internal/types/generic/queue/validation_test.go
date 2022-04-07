package queue

import (
	"github.com/matryer/is"
	"github.com/pkg/errors"
	"testing"
)

// TestValidation_NewQueue checks that default validation checks are run.
func TestValidation_NewQueue(t *testing.T) {
	tests := []struct {
		capacity int
		expErr   error
	}{
		{capacity: 1},
		{capacity: 0},
	}

	for _, tc := range tests {
		assert := is.New(t)

		m, err := New[string](tc.capacity)
		if err != nil {
			if m.ch == nil {
				if !errors.Is(err, ErrChannelMustBeSet) {
					t.Fatalf("expected %v when channel is nil, got %v", ErrChannelMustBeSet, err)
				}
			}
		} else {
			assert.Equal(m.Cap(), tc.capacity)
		}
	}
}
