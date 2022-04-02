package clock

import (
	"time"
)

// Clock represents an interface around time management.
//
// Normally, one would just use the `time` package and be done with it. However, that
// makes it quite difficult to properly test functions are grabbing time values using the wallclock. The
// goal of this interface is to allow time-dependent code to abstract the underlying clock away so we can
// freeze/modify it during testing.
type Clock interface {
	// Now returns the current time for the clock.
	Now() time.Time
}
