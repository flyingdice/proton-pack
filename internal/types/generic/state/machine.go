package state

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/validation"
	"sync"
)

type Action func() error
type Transition func() error

// Machine represents thread-safe state machine.
type Machine[T State] struct {
	state       T
	transitions map[T]map[T]struct{}
	mu          sync.Mutex
}

// NewMachine creates and validates a new Machine.
func NewMachine[T State](initial T, transitions map[T]map[T]struct{}) (*Machine[T], validation.ErrorGroup) {
	m := &Machine[T]{
		state:       initial,
		transitions: transitions,
	}
	return m, m.Check()
}

// Check runs default validation checks for the Machine.
func (m *Machine[T]) Check() validation.ErrorGroup {
	return validation.RunChecks[*Machine[T]](m, defaultChecks[T]()...)
}

// In is true when the machine is in the given state.
func (m *Machine[T]) In(s T) bool { return m.state == s }

// To transitions the state machine to the given state.
func (m *Machine[T]) To(s T, t Transition) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.In(s) {
		return &ErrAlreadyInState[T]{s}
	}
	if _, ok := m.transitions[m.state][s]; !ok {
		return &ErrInvalidTransition[T]{m.state, s}
	}

	return t()
}

// MustBe invokes the given action if the machine is in the expected state.
//
// If not in the expected state, ErrAlreadyInState is returned.
func (m *Machine[T]) MustBe(s T, a Action) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.In(s) {
		return &ErrAlreadyInState[T]{s}
	}

	return a()
}

// String value of the Machine.
//
// Interface: fmt.Stringer.
func (m *Machine[T]) String() string {
	return fmt.Sprintf("Machine[%T](state=%s)", m.state, m.state)
}
