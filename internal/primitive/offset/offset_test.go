package offset

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestOffset_String checks String() output is expected format.
func TestOffset_String(t *testing.T) {
	checker := func(o Offset) bool {
		return o.String() == fmt.Sprintf("%d", int64(o))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestOffset_EqualsTrue checks equality between two Offset instances and their
// int64 representations.
func TestOffset_EqualsTrue(t *testing.T) {
	checker := func(o Offset) bool {
		if !o.Equals(o) {
			return false
		}
		if !o.Equals(int64(o)) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestOffset_EqualsFalse checks inequality between two in-equal Offset instances and their
// int64 representations.
func TestOffset_EqualsFalse(t *testing.T) {
	checker := func(o1 Offset) bool {
		o2 := o1 + 1
		if o1.Equals(o2) {
			return false
		}
		if o1.Equals(int64(o2)) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
