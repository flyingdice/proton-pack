package header

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestHeader_String checks String() output is expected format.
func TestHeader_String(t *testing.T) {
	checker := func(h Header) bool {
		return h.String() == fmt.Sprintf("Header(key=%s val=%s)", h.Key, h.Val)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestHeader_EqualsTrue checks equality between two Header instances.
func TestHeader_EqualsTrue(t *testing.T) {
	checker := func(h Header) bool {
		return h.Equals(h)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestHeader_EqualsFalse checks inequality between two in-equal Header instances.
func TestHeader_EqualsFalse(t *testing.T) {
	checker := func(h1 Header) bool {
		h2 := Header{Key: h1.Key + "foo", Val: h1.Val}
		return !h1.Equals(h2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
