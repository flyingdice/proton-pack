package state

import "fmt"

// ErrAlreadyInState returned when machine attempts to transition
// into its current state.
type ErrAlreadyInState[T State] struct {
	state T
}

func (e *ErrAlreadyInState[T]) Error() string {
	return fmt.Sprintf("machine already in state '%s'", e.state)
}

// ErrNotInState returned when caller asserts a machine to be in a
// specific state and it isn't.
type ErrNotInState[T State] struct {
	state T
}

func (e *ErrNotInState[T]) Error() string {
	return fmt.Sprintf("machine not in state '%s'", e.state)
}

// ErrInvalidTransition returned when call requests a transition
// between states that isn't registered.
type ErrInvalidTransition[T State] struct {
	current T
	next    T
}

func (e *ErrInvalidTransition[T]) Error() string {
	return fmt.Sprintf("machine cannot transition from '%s' to '%s'", e.current, e.next)
}
