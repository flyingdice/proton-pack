package context

import (
	"encoding/json"
	"testing"
	"testing/quick"
)

// TestContext_BinaryEncoding checks Context equality when encoding/decoding to/from binary form.
func TestContext_BinaryEncoding(t *testing.T) {
	checker := func(c1 Context) bool {
		buf, err := c1.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		c2 := &Context{}
		if err := c2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
		}
		return c1.Equals(*c2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestContext_JSONEncoding checks equality when encoding/decoding to/from json form.
func TestContext_JSONEncoding(t *testing.T) {
	checker := func(c1 Context) bool {
		buf, err := json.Marshal(c1)
		if err != nil {
			t.Error(err)
		}
		c2 := Context{}
		if err := json.Unmarshal(buf, &c2); err != nil {
			t.Error(err)
		}
		return c1.Equals(c2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
