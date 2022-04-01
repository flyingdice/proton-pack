package headers

import (
	"encoding/json"
	"testing"
	"testing/quick"
)

// TestHeaders_BinaryEncoding checks Headers equality when encoding/decoding to/from binary form.
func TestHeaders_BinaryEncoding(t *testing.T) {
	checker := func(h1 Headers) bool {
		buf, err := h1.MarshalBinary()
		if err != nil {
			t.Error(err)
			return false
		}
		var h2 Headers
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

// TestHeaders_JSONEncoding checks Headers equality when encoding/decoding to/from JSON form.
func TestHeaders_JSONEncoding(t *testing.T) {
	checker := func(h1 Headers) bool {
		buf, err := json.Marshal(h1)
		if err != nil {
			t.Error(err)
			return false
		}
		var h2 Headers
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
