package state

import (
	"errors"
	"github.com/matryer/is"
	"testing"
)

// TestValidation_NewMachine checks that default validation checks are run.
func TestValidation_NewMachine(t *testing.T) {
	tests := []struct {
		initial     string
		transitions map[string]map[string]struct{}
	}{
		{
			initial:     "n/a",
			transitions: map[string]map[string]struct{}{},
		},
		{
			initial: "closed",
			transitions: map[string]map[string]struct{}{
				"closed": {
					"opened": struct{}{},
				},
				"opened": {
					"closed": struct{}{},
				},
			},
		},
	}

	for _, tc := range tests {
		assert := is.New(t)

		m, err := NewMachine(tc.initial, tc.transitions)
		if err != nil {
			if len(tc.transitions) == 0 {
				if !errors.Is(err, ErrTransitionsMustBeSet) {
					t.Fatalf("expected %v when no transitions, got %v", ErrTransitionsMustBeSet, err)
				}
			}
		} else {
			assert.Equal(m.state, tc.initial)
			assert.Equal(m.transitions, tc.transitions)
		}
	}
}
