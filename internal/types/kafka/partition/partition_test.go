package partition

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestPartition_String checks String() output is expected format.
func TestPartition_String(t *testing.T) {
	checker := func(p Partition) bool {
		return p.String() == fmt.Sprintf("%d", int32(p))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestPartition_EqualsTrue checks equality between two Partition instances and their
// int32 representations.
func TestPartition_EqualsTrue(t *testing.T) {
	checker := func(p Partition) bool {
		if !p.Equals(p) {
			return false
		}
		if !p.Equals(int32(p)) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestPartition_EqualsFalse checks inequality between two in-equal Partition instances and their
// int32 representations.
func TestPartition_EqualsFalse(t *testing.T) {
	checker := func(p1 Partition) bool {
		p2 := p1 + 1
		if p1.Equals(p2) {
			return false
		}
		if p1.Equals(int32(p2)) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
