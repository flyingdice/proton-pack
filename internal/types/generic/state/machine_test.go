package state

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestMachine_String checks String() output is expected format.
func TestMachine_String(t *testing.T) {
	checker := func(s string) bool {
		m := &Machine[string]{state: s}
		return m.String() == fmt.Sprintf("Machine[%T](state=%s)", "", s)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
