package offset

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var ErrMustBePositive = validation.NewCheckError(
	"offset_must_be_positive",
	"the offset value must be a positive number",
)

var defaultChecks = []validation.Check[Offset]{
	checkPositive(),
}

func checkPositive() validation.Check[Offset] {
	return func(o Offset) *validation.CheckError {
		if o < 0 {
			return ErrMustBePositive
		}
		return nil
	}
}
