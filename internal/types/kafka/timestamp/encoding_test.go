package timestamp

import (
	"testing"
	"testing/quick"
)

// TestTimestamp_BinaryEncoding checks Timestamp equality when encoding/decoding to/from binary form.
func TestTimestamp_BinaryEncoding(t *testing.T) {
	checker := func(ts1 Timestamp) bool {
		buf, err := ts1.MarshalBinary()
		if err != nil {
			t.Error(err)
			return false
		}
		var ts2 Timestamp
		if err := ts2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
			return false
		}
		return ts2.Equals(ts2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTimestamp_JSONEncoding checks Timestamp equality when encoding/decoding to/from JSON form.
func TestTimestamp_JSONEncoding(t *testing.T) {
	checker := func(ts1 Timestamp) bool {
		buf, err := ts1.MarshalJSON()
		if err != nil {
			t.Error(err)
			return false
		}
		var ts2 Timestamp
		if err := ts2.UnmarshalJSON(buf); err != nil {
			t.Error(err)
			return false
		}
		return ts1.Equals(ts2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
