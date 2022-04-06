package machine

import (
	"github.com/flyingdice/proton-pack/internal/types/generic/state/state"
	"github.com/flyingdice/proton-pack/internal/validation"
)

// ErrStatesMustBeSet is the validation check error returned when
// the machine states are empty.
var ErrStatesMustBeSet = validation.NewError(
	"machine_states_must_be_set",
	"the machine states must be set and cannot be an empty",
)

func defaultChecks[T state.State]() []validation.Check[*Machine[T]] {
	return []validation.Check[*Machine[T]]{
		checkTransitionsValid[T](),
		checkStatesSet[T](),
	}
}

// checkTransitionsValid validates machine transitions are set.
func checkTransitionsValid[T state.State]() validation.Check[*Machine[T]] {
	return func(m *Machine[T]) *validation.Error {
		if err := validation.RunCheckers(m.transitions); err != nil {
			return validation.NewError(
				"machine_must_have_valid_fields",
				err.Error(),
			)
		}
		return nil
	}
}

// checkStatesSet validates machine states are set.
func checkStatesSet[T state.State]() validation.Check[*Machine[T]] {
	return func(m *Machine[T]) *validation.Error {
		if len(m.states) == 0 {
			return ErrStatesMustBeSet
		}
		return nil
	}
}
