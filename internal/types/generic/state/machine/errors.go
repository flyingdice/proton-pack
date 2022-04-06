package machine

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/types/generic/state/state"
)

// ErrAlreadyInState returned when machine attempts to transition
// into its current state.
type ErrAlreadyInState[T state.State] struct {
	State T
}

func (e *ErrAlreadyInState[T]) Error() string {
	return fmt.Sprintf("machine already in state '%s'", e.State)
}

// ErrNotInState returned when caller asserts a machine to be in a
// specific state and it isn't.
type ErrNotInState[T state.State] struct {
	State T
}

func (e *ErrNotInState[T]) Error() string {
	return fmt.Sprintf("machine not in state '%s'", e.State)
}

// ErrInvalidTransition returned when call requests a transition
// between states that isn't registered.
type ErrInvalidTransition[T state.State] struct {
	Current T
	Next    T
}

func (e *ErrInvalidTransition[T]) Error() string {
	return fmt.Sprintf("machine cannot transition from '%s' to '%s'", e.Current, e.Next)
}
