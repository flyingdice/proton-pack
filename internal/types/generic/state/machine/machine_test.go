package machine

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/testing/assertion"
	"testing"
	"testing/quick"
)

// TestMachine_String checks String() output is expected format.
func TestMachine_String(t *testing.T) {
	assert := assertion.Error(t)
	checker := func(m *Machine[string]) bool {
		got := m.String()
		want := fmt.Sprintf("Machine[%T](state=%s)", "", m.state)
		return assert.Equal(got, want)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestMachine_To checks machine can transition between states.
func TestMachine_To(t *testing.T) {
	assert := assertion.Error(t)
	checker := func(m *Machine[string]) bool {
		// Attempt transition to unknown states should error.
		for _, state := range m.states {
			if !m.transitions.Valid(m.state, state) {
				err := m.To(state, NoOp)
				if !assert.Equal(err.Error(), (&ErrInvalidTransition[string]{m.state, state}).Error()) {
					return false
				}
			}
		}

		// Transition to a valid state should succeed.
		for _, state := range m.states {
			if m.transitions.Valid(m.state, state) {
				err := m.To(state, NoOp)
				if !assert.OK(err) {
					return false
				}
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
	assert := assertion.Error(t)
	checker := func(m *Machine[string]) bool {
		// Assert initial state no error.
		err := m.MustBe(m.state, NoOp)
		if !assert.OK(err) {
			return false
		}

		// Assert incorrect state is error.
		for _, state := range m.states {
			if state != m.state {
				err := m.MustBe(state, NoOp)
				if !assert.Equal(err.Error(), (&ErrNotInState[string]{state}).Error()) {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
