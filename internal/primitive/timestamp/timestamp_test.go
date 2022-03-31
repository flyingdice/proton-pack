package timestamp

import (
	"fmt"
	"testing"
	"testing/quick"
	"time"
)

// TestTimestamp_String checks String() output is expected format.
func TestTimestamp_String(t *testing.T) {
	checker := func(ts Timestamp) bool {
		return ts.String() == fmt.Sprintf("%s", ts.Format(time.RFC3339))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTimestamp_EqualsTrue checks equality between two Timestamp instances and their
// int32 representations.
func TestTimestamp_EqualsTrue(t *testing.T) {
	checker := func(ts Timestamp) bool {
		if !ts.Equals(ts) {
			return false
		}
		if !ts.Equals(ts.Time) {
			return false
		}
		if !ts.Equals(ts.UnixMilli()) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestTimestamp_EqualsFalse checks inequality between two in-equal Timestamp instances and their
// int32 representations.
func TestTimestamp_EqualsFalse(t *testing.T) {
	checker := func(ts1 Timestamp) bool {
		ts2 := Timestamp{
			Time: ts1.Add(1 * time.Second),
		}
		if ts1.Equals(ts2) {
			return false
		}
		if ts1.Equals(ts2.Time) {
			return false
		}
		if ts1.Equals(ts2.UnixMilli()) {
			return false
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
