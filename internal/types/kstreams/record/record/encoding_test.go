package record

import (
	"encoding/json"
	"testing"
	"testing/quick"
)

// TestRecord_BinaryEncoding checks Record equality when encoding/decoding to/from binary form.
func TestRecord_BinaryEncoding(t *testing.T) {
	t.Skip("expected failure until key/val serde complete")

	checker := func(c1 Record) bool {
		buf, err := c1.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		c2 := &Record{}
		if err := c2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
		}
		return c1.Equals(*c2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestRecord_JSONEncoding checks equality when encoding/decoding to/from json form.
func TestRecord_JSONEncoding(t *testing.T) {
	checker := func(c1 Record) bool {
		buf, err := json.Marshal(c1)
		if err != nil {
			t.Error(err)
		}
		c2 := Record{}
		if err := json.Unmarshal(buf, &c2); err != nil {
			t.Error(err)
		}
		return c1.Equals(c2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
