package record

import (
	"github.com/bxcodec/faker/v3"
	"github.com/flyingdice/proton-pack/internal/types/kafka/headers"
	"github.com/flyingdice/proton-pack/internal/types/kstreams/record/metadata"
	"testing"
	"testing/quick"
)

// TestValidation_NewRecord checks that default validation checks are run.
func TestValidation_NewRecord(t *testing.T) {
	checker := func(m metadata.Metadata, h headers.Headers) bool {
		_, err := NewRecord(faker.Word(), faker.Sentence(), m, h)
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
