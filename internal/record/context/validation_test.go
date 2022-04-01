package context

import (
	"github.com/flyingdice/proton-pack/internal/primitive/headers"
	"github.com/flyingdice/proton-pack/internal/record/metadata"
	"testing"
	"testing/quick"
)

// TestValidation_NewContext checks that default validation checks are run.
func TestValidation_NewContext(t *testing.T) {
	checker := func(m metadata.Metadata, h headers.Headers) bool {
		_, err := NewContext(m, h)
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
