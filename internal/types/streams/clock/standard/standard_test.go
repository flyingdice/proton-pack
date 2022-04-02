package standard

import (
	"strings"
	"testing"
	"testing/quick"
	"time"
)

// TestClock_String checks String() output is expected format.
func TestContext_String(t *testing.T) {
	checker := func(c Clock) bool {
		// Can't do exact equality since c.Now() is unstable. Prefix match fine for now.
		return strings.HasPrefix(c.String(), "StandardClock")
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestClock_Now is an increasing value using the system clock.
func TestClock_Now(t *testing.T) {
	checker := func(c Clock) bool {
		t1 := c.Now()
		time.Sleep(10 * time.Millisecond)
		t2 := c.Now()
		return t1.Before(t2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
