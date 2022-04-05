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
		states      []string
		transitions map[string]map[string]struct{}
	}{
		{
			initial: "n/a",
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

		m, err := NewMachine[string](tc.initial, tc.states, Transitions[string]{tc.transitions})
		if err != nil {
			if len(tc.transitions) == 0 {
				if !errors.Is(err, ErrTransitionsMustBeSet) {
					t.Fatalf("expected %v when no transitions, got %v", ErrTransitionsMustBeSet, err)
				}
			} else if len(tc.states) == 0 {
				if !errors.Is(err, ErrStatesMustBeSet) {
					t.Fatalf("expected %v when no states, got %v", ErrStatesMustBeSet, err)
				}
			}
		} else {
			assert.Equal(m.state, tc.initial)
			assert.Equal(m.transitions, tc.transitions)
		}
	}
}
