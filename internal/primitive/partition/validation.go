package partition

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var ErrMustBePositive = validation.NewCheckError(
	"partition_must_be_positive",
	"the partition value must be a positive number",
)

var defaultChecks = []validation.Check[Partition]{
	checkPositive(),
}

func checkPositive() validation.Check[Partition] {
	return func(p Partition) validation.CheckError {
		if p < 0 {
			return ErrMustBePositive
		}
		return nil
	}
}
