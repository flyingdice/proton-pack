package machine

import (
	"github.com/flyingdice/proton-pack/internal/types/generic/state/transition"
	"github.com/matryer/is"
	"testing"
)

// TestValidation_NewMachine checks that default validation checks are run.
func TestValidation_NewMachine(t *testing.T) {
	tests := []struct {
		initial string
		states  []string
		table   transition.Table[string]
		expErr  error
	}{
		{
			initial: "",
			states:  nil,
			expErr:  transition.ErrTableMustBeSet,
		},
		{
			initial: "foo",
			states:  nil,
			table: transition.Table[string]{
				"foo": {"bar": true},
			},
			expErr: ErrStatesMustBeSet,
		},
		{
			initial: "closed",
			states:  []string{"opened", "closed"},
			table: transition.Table[string]{
				"closed": {
					"opened": true,
				},
				"opened": {
					"closed": true,
				},
			},
		},
	}

	for _, tc := range tests {
		assert := is.New(t)

		m, err := NewMachine[string](tc.initial, tc.states, tc.table)
		if err != nil {
			if tc.expErr != nil {
				assert.Equal(err.Error(), tc.expErr.Error())
			} else {
				t.Fatalf("unexpected error %v", err)
			}
		} else {
			assert.Equal(m.state, tc.initial)
			assert.Equal(m.states, tc.states)
			assert.Equal(m.transitions.Len(), len(tc.table))
		}
	}
}
