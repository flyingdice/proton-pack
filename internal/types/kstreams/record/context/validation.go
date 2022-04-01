package context

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var defaultChecks = []validation.Check[Context]{
	checkFieldsAreValid(),
}

// checkFieldsAreValid validates context metadata and headers are valid.
func checkFieldsAreValid() validation.Check[Context] {
	return func(c Context) *validation.Error {
		if err := validation.RunCheckers(c.Metadata, c.Headers); err != nil {
			return validation.NewError(
				"context_must_have_valid_fields",
				err.Error(),
			)
		}
		return nil
	}
}
