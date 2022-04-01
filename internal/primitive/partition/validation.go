package partition

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

// ErrMustBePositive is the validation check error returned when
// the partition is a negative number.
var ErrMustBePositive = validation.NewError(
	"partition_must_be_positive",
	"the partition value must be a positive number",
)

var defaultChecks = []validation.Check[Partition]{
	checkPositive(),
}

// checkPositive validates partition is a positive number.
func checkPositive() validation.Check[Partition] {
	return func(p Partition) *validation.Error {
		if p < 0 {
			return ErrMustBePositive
		}
		return nil
	}
}
