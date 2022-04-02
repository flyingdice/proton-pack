package window

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

var _ fmt.Stringer = (*Window)(nil)
var _ quick.Generator = (*Window)(nil)
var _ comparison.Equaler = (*Window)(nil)
var _ validation.Checker = (*Window)(nil)

type Window struct {
	lo time.Time
	hi time.Time
}

// NewWindow creates and validates a new Window from the given values.
func NewWindow(lo, hi time.Time) (Window, validation.ErrorGroup) {
	w := Window{lo, hi}
	return w, w.Check()
}

// Check runs default validation checks for the Window.
func (w Window) Check() validation.ErrorGroup {
	return validation.RunChecks[Window](w, defaultChecks...)
}

// Duration returns the time duration of the window.
func (w Window) Duration() time.Duration {
	return w.hi.Sub(w.lo)
}

// Overlaps returns true when the two time windows overlap.
func (w Window) Overlaps(w2 Window) bool {
	return w.lo.Before(w2.hi) && w2.lo.Before(w.hi)
}

// Equals compares two Window instances for equality.
//
// Interface: comparison.Equaler
func (w Window) Equals(v any) bool {
	switch w2 := v.(type) {
	case Window:
		return w.lo == w2.lo && w.hi == w2.hi
	default:
		return false
	}
}

// Generate random Window values.
//
// Interface: quick.Generator
func (Window) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Window.
//
// Interface: fmt.Stringer.
func (w Window) String() string {
	return fmt.Sprintf(
		"Window(lo=%s hi=%s)",
		w.lo.Format(time.RFC3339),
		w.hi.Format(time.RFC3339),
	)
}

// Generate a random Window value.
func Generate(rand *rand.Rand) Window {
	lo := time.Unix(0, 0).Add(time.Duration(rand.Int63n(math.MaxInt64)))
	hi := lo.Add(time.Duration(rand.Int31n(math.MaxInt32)))
	return Window{lo, hi}
}
