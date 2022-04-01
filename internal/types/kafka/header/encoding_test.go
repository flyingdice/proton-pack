package header

import (
	"encoding/json"
	"testing"
	"testing/quick"
)

// TestHeader_BinaryEncoding checks Header equality when encoding/decoding to/from binary form.
func TestHeader_BinaryEncoding(t *testing.T) {
	checker := func(h1 Header) bool {
		buf, err := h1.MarshalBinary()
		if err != nil {
			t.Error(err)
			return false
		}
		var h2 Header
		if err := h2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
			return false
		}

		return h1.Equals(h2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestHeader_JSONEncoding checks Header equality when encoding/decoding to/from JSON form.
func TestHeader_JSONEncoding(t *testing.T) {
	checker := func(h1 Header) bool {
		buf, err := json.Marshal(h1)
		if err != nil {
			t.Error(err)
			return false
		}
		var h2 Header
		if err := json.Unmarshal(buf, &h2); err != nil {
			t.Error(err)
			return false
		}
		return h1.Equals(h2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
