package header

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var ErrKeyMustBeSet = validation.NewCheckError(
	"header_key_must_be_set",
	"the header key must be set and cannot be an empty string",
)

var defaultChecks = []validation.Check[Header]{
	checkKeySet(),
}

func checkKeySet() validation.Check[Header] {
	return func(h Header) *validation.CheckError {
		if h.Key == "" {
			return ErrKeyMustBeSet
		}
		return nil
	}
}
