package state

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

// ErrTransitionsMustBeSet is the validation check error returned when
// the machine transitions are empty.
var ErrTransitionsMustBeSet = validation.NewError(
	"machine_transitions_must_be_set",
	"the machine transitions must be set and cannot be an empty",
)

// ErrStatesMustBeSet is the validation check error returned when
// the machine states are empty.
var ErrStatesMustBeSet = validation.NewError(
	"machine_states_must_be_set",
	"the machine states must be set and cannot be an empty",
)

func defaultChecks[T State]() []validation.Check[*Machine[T]] {
	return []validation.Check[*Machine[T]]{
		checkTransitionsSet[T](),
		checkStatesSet[T](),
	}
}

// checkTransitionsValid validates machine transitions are set.
func checkTransitionsSet[T State]() validation.Check[*Machine[T]] {
	return func(m *Machine[T]) *validation.Error {
		if m.transitions.Len() == 0 {
			return ErrTransitionsMustBeSet
		}
		return nil
	}
}

// checkStatesSet validates machine states are set.
func checkStatesSet[T State]() validation.Check[*Machine[T]] {
	return func(m *Machine[T]) *validation.Error {
		if len(m.states) == 0 {
			return ErrStatesMustBeSet
		}
		return nil
	}
}
