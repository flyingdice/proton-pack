package frozen

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/types/streams/clock"
	"math"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

var _ fmt.Stringer = (*Clock)(nil)
var _ quick.Generator = (*Clock)(nil)
var _ clock.Clock = (*Clock)(nil)

// Clock implementation that is frozen in time. Useful for tests.
type Clock struct {
	now time.Time
}

// Now returns the current time for the clock.
//
// Interface: clock.Clock
func (c Clock) Now() time.Time {
	return c.now
}

// Generate random Clock values.
//
// Interface: quick.Generator
func (Clock) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Clock.
//
// Interface: fmt.Stringer.
func (c Clock) String() string {
	return fmt.Sprintf("FrozenClock(now=%s)", c.Now().Format(time.RFC3339))
}

// Generate a new frozen clock.
func Generate(rand *rand.Rand) Clock {
	now := time.Unix(0, 0).Add(time.Duration(rand.Int63n(math.MaxInt64)))
	return Clock{now}
}
