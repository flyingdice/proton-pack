package topic

import (
	"testing"
	"testing/quick"
)

// TestTopic_BinaryEncoding checks Topic equality when encoding/decoding to/from binary form.
func TestTopic_BinaryEncoding(t *testing.T) {
	checker := func(t1 Topic) bool {
		buf, err := t1.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		var t2 Topic
		if err := t2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
		}
		return t1.Equals(t2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTopic_JSONEncoding checks Topic equality when encoding/decoding to/from JSON form.
func TestTopic_JSONEncoding(t *testing.T) {
	checker := func(t1 Topic) bool {
		buf, err := t1.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		var t2 Topic
		if err := t2.UnmarshalJSON(buf); err != nil {
			t.Error(err)
		}
		return t1.Equals(t2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
