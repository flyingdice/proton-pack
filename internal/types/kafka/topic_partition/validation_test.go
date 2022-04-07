package topic_partition

import (
	"github.com/flyingdice/proton-pack/internal/types/kafka/partition"
	"github.com/flyingdice/proton-pack/internal/types/kafka/topic"
	"testing"
	"testing/quick"
)

// TestValidation_NewTopicPartition checks that default validation checks are run.
func TestValidation_NewTopicPartition(t *testing.T) {
	checker := func(to topic.Topic, p partition.Partition) bool {
		_, err := New(to, p)
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
