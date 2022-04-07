package offset

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*Offset)(nil)
var _ quick.Generator = (*Offset)(nil)
var _ comparison.Equaler = (*Offset)(nil)
var _ validation.Checker = (*Offset)(nil)

// Offset represents the position of a message within a topic+partition.
type Offset int64

// New creates and validates a new Offset from the given int64.
func New(v int64) (Offset, validation.ErrorGroup) {
	o := Offset(v)
	return o, o.Check()
}

// Check runs default validation checks for the Offset.
func (o Offset) Check() validation.ErrorGroup {
	return validation.RunChecks[Offset](o, defaultChecks...)
}

// Equals compares two Offset instances for equality.
//
// Interface: comparison.Equaler
func (o Offset) Equals(v any) bool {
	switch o2 := v.(type) {
	case Offset:
		return o == o2
	case int64:
		return int64(o) == o2
	default:
		return false
	}
}

// Generate random Offset values.
//
// Interface: quick.Generator
func (Offset) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Offset.
//
// Interface: fmt.Stringer.
func (o Offset) String() string {
	return fmt.Sprintf("%d", o)
}

// Generate a random Offset value.
func Generate(rand *rand.Rand) Offset {
	return Offset(rand.Int63n(math.MaxInt64))
}
