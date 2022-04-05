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

func defaultChecks[T State]() []validation.Check[*Machine[T]] {
	return []validation.Check[*Machine[T]]{
		checkTransitionsValid[T](),
	}
}

// checkTransitionsValid validates machine transitions are valid.
func checkTransitionsValid[T State]() validation.Check[*Machine[T]] {
	return func(m *Machine[T]) *validation.Error {
		if len(m.transitions) == 0 {
			return ErrTransitionsMustBeSet
		}
		return nil
	}
}
