package metadata

import (
	"github.com/flyingdice/proton-pack/internal/primitive/offset"
	"github.com/flyingdice/proton-pack/internal/primitive/partition"
	"github.com/flyingdice/proton-pack/internal/primitive/timestamp"
	"github.com/flyingdice/proton-pack/internal/primitive/topic"
	"testing"
	"testing/quick"
)

// TestValidation_NewMetadata checks that default validation checks are run.
func TestValidation_NewMetadata(t *testing.T) {
	checker := func(p partition.Partition, o offset.Offset, ts timestamp.Timestamp, to topic.Topic) bool {
		_, err := NewMetadata(p, o, ts, to)
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
