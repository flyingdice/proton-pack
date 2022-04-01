package headers

import (
	"github.com/flyingdice/proton-pack/internal/primitive/header"
	"testing"
	"testing/quick"
)

// TestValidation_NewHeaders checks that default validation checks are run.
func TestValidation_NewHeaders(t *testing.T) {
	checker := func(headers []header.Header) bool {
		_, err := NewHeaders(headers)
		if err != nil {
			t.Error(err)
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
