package topic_partition

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestTopicPartition_String checks String() output is expected format.
func TestTopicPartition_String(t *testing.T) {
	checker := func(t1 TopicPartition) bool {
		return t1.String() == fmt.Sprintf("%s-%s", t1.Topic, t1.Partition)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTopicPartition_EqualsTrue checks equality between two TopicPartition instances and their
// string representations.
func TestTopicPartition_EqualsTrue(t *testing.T) {
	checker := func(t1 TopicPartition) bool {
		return t1.Equals(t1)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTopicPartition_EqualsFalse checks inequality between two in-equal TopicPartition instances and their
// string representations.
func TestTopicPartition_EqualsFalse(t *testing.T) {
	checker := func(t1 TopicPartition) bool {
		t2 := TopicPartition{
			Topic:     t1.Topic + "foo",
			Partition: t1.Partition,
		}
		return !t1.Equals(t2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
