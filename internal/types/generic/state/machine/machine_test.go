package machine

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestMachine_String checks String() output is expected format.
func TestMachine_String(t *testing.T) {
	checker := func(m *Machine[string]) bool {
		return m.String() == fmt.Sprintf("Machine[%T](state=%s)", "", m.state)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
