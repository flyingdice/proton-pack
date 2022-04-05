package machine

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/types/generic/state/state"
	"github.com/flyingdice/proton-pack/internal/types/generic/state/transition"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
	"sync"
	"testing/quick"
)

var _ fmt.Stringer = (*Machine[string])(nil)
var _ quick.Generator = (*Machine[string])(nil)
var _ validation.Checker = (*Machine[string])(nil)

type Action func() error
type Transition func() error

// Machine represents thread-safe state machine.
type Machine[T state.State] struct {
	state       T
	states      []T
	transitions *transition.Transitions[T]
	mu          sync.Mutex
}

// NewMachine creates and validates a new Machine.
func NewMachine[T state.State](
	initial T,
	states []T,
	table transition.Table[T],
) (*Machine[T], validation.ErrorGroup) {
	transitions, err := transition.NewTransitions[T](table)
	if err != nil {
		return nil, err
	}
	m := &Machine[T]{
		state:       initial,
		states:      states,
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
	if !m.transitions.Valid(m.state, s) {
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

// Generate random Machine values.
//
// Interface: quick.Generator
func (*Machine[T]) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate[T](rand))
}

// String value of the Machine.
//
// Interface: fmt.Stringer.
func (m *Machine[T]) String() string {
	return fmt.Sprintf("Machine[%T](state=%s)", m.state, m.state)
}

// Generate a random Machine value.
func Generate[T state.State](rand *rand.Rand) *Machine[T] {
	states := state.Generate[T](rand)
	transitions := transition.Generate[T](rand, states)

	return &Machine[T]{
		state:       states[rand.Intn(len(states))],
		states:      states,
		transitions: transitions,
	}
}
