package queue

import (
	"github.com/flyingdice/proton-pack/internal/testing/assertion"
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
		assert := assertion.Fatal(t)

		m, err := New[string](tc.capacity)
		if err != nil {
			if m.ch == nil {
				assert.Equal(err, ErrChannelMustBeSet)
			}
		} else {
			assert.Equal(m.Cap(), tc.capacity)
		}
	}
}
