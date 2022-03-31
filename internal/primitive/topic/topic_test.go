package topic

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestTopic_String checks String() output is expected format.
func TestTopic_String(t *testing.T) {
	checker := func(t1 Topic) bool {
		return t1.String() == fmt.Sprintf("%s", string(t1))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTopic_EqualsTrue checks equality between two Topic instances and their
// string representations.
func TestTopic_EqualsTrue(t *testing.T) {
	checker := func(t1 Topic) bool {
		if !t1.Equals(t1) {
			return false
		}
		if !t1.Equals(string(t1)) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTopic_EqualsFalse checks inequality between two in-equal Topic instances and their
// string representations.
func TestTopic_EqualsFalse(t *testing.T) {
	checker := func(t1 Topic) bool {
		t2 := t1 + "foo"
		if t1.Equals(t2) {
			return false
		}
		if t1.Equals(string(t2)) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
