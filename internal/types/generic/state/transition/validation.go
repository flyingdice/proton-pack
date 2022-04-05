package transition

import (
	"github.com/flyingdice/proton-pack/internal/types/generic/state/state"
	"github.com/flyingdice/proton-pack/internal/validation"
)

// ErrTableMustBeSet is the validation check error returned when
// the transitions table is empty.
var ErrTableMustBeSet = validation.NewError(
	"transitions_table_must_be_set",
	"the transitions table must be set and cannot be an empty",
)

func defaultChecks[T state.State]() []validation.Check[*Transitions[T]] {
	return []validation.Check[*Transitions[T]]{
		checkTableSet[T](),
	}
}

// checkTableSet validates transitions table is set.
func checkTableSet[T state.State]() validation.Check[*Transitions[T]] {
	return func(m *Transitions[T]) *validation.Error {
		if len(m.table) == 0 {
			return ErrTableMustBeSet
		}
		return nil
	}
}
