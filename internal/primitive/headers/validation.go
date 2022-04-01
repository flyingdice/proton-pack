package headers

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

var defaultChecks = []validation.Check[Headers]{
	checkEachHeaderIsValid(),
}

// checkEachHeaderIsValid validates each header in the slice is also valid.
func checkEachHeaderIsValid() validation.Check[Headers] {
	return func(h Headers) *validation.Error {
		var errs *validation.Errors

		for _, hdr := range h {
			if err := hdr.Check(); err != nil {
				errs.Append(err.Errors()...)
			}
		}

		if errs.NilWhenEmpty() == nil {
			return nil
		}

		return validation.NewError(
			"headers_must_have_valid_headers",
			errs.Error(),
		)
	}
}
