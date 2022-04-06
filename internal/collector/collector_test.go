package collector

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/types/generic/state/machine"
	"testing"
	"testing/quick"
)

// TestCollector_String checks String() output is expected format.
func TestCollector_String(t *testing.T) {
	checker := func(c *Collector) bool {
		return c.String() == fmt.Sprintf("Collector(state=%s offsets=%s)", c.machine, c.offsets)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestCollector_StateMachine checks state transitions of collector.
func TestCollector_StateMachine(t *testing.T) {
	checker := func(c *Collector) bool {
		// Initial state is Closed.
		if !wantState(t, Closed, c.machine.Current()) {
			return false
		}

		// Close() while Opened is error.
		err := c.Close()
		if err == nil {
			t.Errorf("expected error closing closed collector; got none")
			return false
		}
		got := err.Error()
		want := (&machine.ErrInvalidTransition[State]{Closed, c.machine.Current()}).Error()
		if want != got {
			t.Errorf("expected error %s; got %s", want, got)
			return false
		}

		// Open() transitions self to Opened.
		err = c.Open()
		if err != nil {
			t.Errorf("unexpected error opening collector %s", err)
			return false
		}
		if !wantState(t, Opened, c.machine.Current()) {
			return false
		}

		// Open() while Opened is error.
		err = c.Open()
		if err == nil {
			t.Errorf("expected error opening opened collector; got none")
			return false
		}
		got = err.Error()
		want = (&machine.ErrInvalidTransition[State]{Opened, c.machine.Current()}).Error()
		if want != got {
			t.Errorf("expected error %s; got %s", want, got)
			return false
		}

		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

func wantState(t *testing.T, want, got State) bool {
	t.Helper()
	if want != got {
		t.Errorf("expected state %s; got %s", want, got)
		return false
	}
	return true
}
