package standard

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/types/streams/clock"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

var _ fmt.Stringer = (*Clock)(nil)
var _ quick.Generator = (*Clock)(nil)
var _ clock.Clock = (*Clock)(nil)

// Clock is the standard clock implementation that uses stdlib time.
type Clock struct{}

// Now returns the current time for the clock.
//
// Interface: clock.Clock
func (c Clock) Now() time.Time {
	return time.Now()
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
	return fmt.Sprintf("StandardClock(now=%s)", c.Now().Format(time.RFC3339))
}

// Generate a new standard clock.
func Generate(rand *rand.Rand) Clock {
	return Clock{}
}
