package frozen

import (
	"fmt"
	"testing"
	"testing/quick"
	"time"
)

// TestClock_String checks String() output is expected format.
func TestContext_String(t *testing.T) {
	checker := func(c Clock) bool {
		return c.String() == fmt.Sprintf("FrozenClock(now=%s)", c.Now().Format(time.RFC3339))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestClock_Now is always a fixed value.
func TestClock_Now(t *testing.T) {
	checker := func(c Clock) bool {
		t1 := c.Now()
		time.Sleep(10 * time.Millisecond)
		t2 := c.Now()
		return t1.Equal(t2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
