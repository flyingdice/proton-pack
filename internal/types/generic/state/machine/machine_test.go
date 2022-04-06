package machine

import (
	"fmt"
	"github.com/matryer/is"
	"testing"
	"testing/quick"
)

// TestMachine_String checks String() output is expected format.
func TestMachine_String(t *testing.T) {
	checker := func(m *Machine[string]) bool {
		return m.String() == fmt.Sprintf("Machine[%T](state=%s)", "", m.state)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestMachine_To checks machine can transition between states.
func TestMachine_To(t *testing.T) {
	checker := func(m *Machine[string]) bool {
		assert := is.New(t)

		// Attempt transition to unknown states should error.
		for _, state := range m.states {
			if !m.transitions.Valid(m.state, state) {
				err := m.To(state, NoOp)
				assert.Equal(err.Error(), (&ErrInvalidTransition[string]{m.state, state}).Error())
			}
		}

		// Transition to a valid state should succeed.
		for _, state := range m.states {
			if m.transitions.Valid(m.state, state) {
				err := m.To(state, NoOp)
				assert.NoErr(err)
				break
			}
		}

		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestMachine_MustBe checks machine is in the specified state before running an action.
func TestMachine_MustBe(t *testing.T) {
	checker := func(m *Machine[string]) bool {
		assert := is.New(t)

		// Assert initial state no error.
		err := m.MustBe(m.state, NoOp)
		assert.NoErr(err)

		// Assert incorrect state is error.
		for _, state := range m.states {
			if state != m.state {
				err := m.MustBe(state, NoOp)
				assert.Equal(err.Error(), (&ErrNotInState[string]{state}).Error())
			}
		}

		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
