package metadata

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var defaultChecks = []validation.Check[Metadata]{
	checkFieldsAreValid(),
}

// checkFieldsAreValid validates metadata fields are valid.
func checkFieldsAreValid() validation.Check[Metadata] {
	return func(m Metadata) *validation.Error {
		if err := validation.RunCheckers(
			m.Partition,
			m.Offset,
			m.Timestamp,
			m.Topic,
		); err != nil {
			return validation.NewError(
				"metadata_must_have_valid_fields",
				err.Error(),
			)
		}
		return nil
	}
}
