package header

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

// ErrKeyMustBeSet is the validation check error returned when
// the header key is an empty string.
var ErrKeyMustBeSet = validation.NewError(
	"header_key_must_be_set",
	"the header key must be set and cannot be an empty string",
)

var defaultChecks = []validation.Check[Header]{
	checkKeyNotEmpty(),
}

// checkKeyNotEmpty validates header key is not an empty string.
func checkKeyNotEmpty() validation.Check[Header] {
	return func(h Header) *validation.Error {
		if h.Key == "" {
			return ErrKeyMustBeSet
		}
		return nil
	}
}
