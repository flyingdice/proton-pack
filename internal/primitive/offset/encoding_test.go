package offset

import (
	"testing"
	"testing/quick"
)

// TestOffset_BinaryEncoding checks Offset equality when encoding/decoding to/from binary form.
func TestOffset_BinaryEncoding(t *testing.T) {
	checker := func(o1 Offset) bool {
		buf, err := o1.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		var o2 Offset
		if err := o2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
		}
		return o1.Equals(o2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestOffset_JSONEncoding checks Offset equality when encoding/decoding to/from JSON form.
func TestOffset_JSONEncoding(t *testing.T) {
	checker := func(o1 Offset) bool {
		buf, err := o1.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		var o2 Offset
		if err := o2.UnmarshalJSON(buf); err != nil {
			t.Error(err)
		}
		return o1.Equals(o2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
