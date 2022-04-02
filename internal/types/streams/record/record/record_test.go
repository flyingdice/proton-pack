package record

import (
	"fmt"
	"testing"
	"testing/quick"
)

// v checks String() output is expected format.
func TestRecord_String(t *testing.T) {
	checker := func(r Record) bool {
		return r.String() == fmt.Sprintf(
			"Record(key=%v val=%v metadata=%v headers=%v)",
			r.Key,
			r.Val,
			r.Metadata,
			r.Headers,
		)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestRecord_EqualsTrue checks equality between two Record instances.
func TestRecord_Check_EqualsTrue(t *testing.T) {
	checker := func(r Record) bool {
		return r.Equals(r)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
