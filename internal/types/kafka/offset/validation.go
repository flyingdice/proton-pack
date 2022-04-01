package offset

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

// ErrMustBePositive is the validation check error returned when
// an offset is a negative number.
var ErrMustBePositive = validation.NewError(
	"offset_must_be_positive",
	"the offset value must be a positive number",
)

var defaultChecks = []validation.Check[Offset]{
	checkPositive(),
}

// checkPositive validates offset value is a positive number.
func checkPositive() validation.Check[Offset] {
	return func(o Offset) *validation.Error {
		if o < 0 {
			return ErrMustBePositive
		}
		return nil
	}
}
