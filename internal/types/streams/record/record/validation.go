package record

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var defaultChecks = []validation.Check[Record]{
	checkFieldsAreValid(),
}

// checkFieldsAreValid validates record metadata and headers are valid.
func checkFieldsAreValid() validation.Check[Record] {
	return func(r Record) *validation.Error {
		if err := validation.RunCheckers(r.Metadata, r.Headers); err != nil {
			return validation.NewError(
				"record_must_have_valid_fields",
				err.Error(),
			)
		}
		return nil
	}
}
