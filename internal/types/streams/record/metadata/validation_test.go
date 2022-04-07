package metadata

import (
	"github.com/flyingdice/proton-pack/internal/types/kafka/offset"
	"github.com/flyingdice/proton-pack/internal/types/kafka/partition"
	"github.com/flyingdice/proton-pack/internal/types/kafka/timestamp"
	"github.com/flyingdice/proton-pack/internal/types/kafka/topic"
	"testing"
	"testing/quick"
)

// TestValidation_NewMetadata checks that default validation checks are run.
func TestValidation_NewMetadata(t *testing.T) {
	checker := func(p partition.Partition, o offset.Offset, ts timestamp.Timestamp, to topic.Topic) bool {
		_, err := New(p, o, ts, to)
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
