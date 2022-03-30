package offset

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
)

var _ fmt.Stringer = (*Offset)(nil)
var _ comparison.Equaler = (*Offset)(nil)

// Offset represents the position of a message within a topic+partition.
type Offset int64

// NewOffset creates and validates a new Offset from the given int64.
func NewOffset(v int64) (Offset, error) {
	o := Offset(v)
	return o, validation.Validate[Offset](o, defaultChecks...)
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
	return Offset(rand.Int63())
}
