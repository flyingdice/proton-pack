package context

import (
	"context"
	"github.com/flyingdice/proton-pack/internal/types/kafka/headers"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/metadata"
	"testing"
	"testing/quick"
)

// TestValidation_NewContext checks that default validation checks are run.
func TestValidation_NewContext(t *testing.T) {
	checker := func(m metadata.Metadata, h headers.Headers) bool {
		_, err := New(context.TODO(), m, h)
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
