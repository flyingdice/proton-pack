package topic_partition

import (
	"encoding/json"
	"testing"
	"testing/quick"
)

// TestTopicPartition_BinaryEncoding checks TopicPartition equality when encoding/decoding to/from binary form.
func TestTopicPartition_BinaryEncoding(t *testing.T) {
	checker := func(tp1 TopicPartition) bool {
		buf, err := tp1.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		tp2 := &TopicPartition{}
		if err := tp2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
		}
		return tp1.Equals(*tp2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTopicPartition_JSONEncoding checks equality when encoding/decoding to/from json form.
func TestTopicPartition_JSONEncoding(t *testing.T) {
	checker := func(tp1 TopicPartition) bool {
		buf, err := json.Marshal(tp1)
		if err != nil {
			t.Error(err)
		}
		tp2 := TopicPartition{}
		if err := json.Unmarshal(buf, &tp2); err != nil {
			t.Error(err)
		}
		return tp1.Equals(tp2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
