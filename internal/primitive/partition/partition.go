package partition

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
)

var _ fmt.Stringer = (*Partition)(nil)
var _ comparison.Equaler = (*Partition)(nil)

// Partition represents the partition of a topic.
type Partition int32

// NewPartition creates and validates a new Partition from the given int32.
func NewPartition(v int32) (Partition, error) {
	p := Partition(v)
	return p, validation.Validate[Partition](p, defaultChecks...)
}

// Equals compares two Partition instances for equality.
//
// Interface: comparison.Equaler
func (p Partition) Equals(v any) bool {
	switch p2 := v.(type) {
	case Partition:
		return p == p2
	case int32:
		return int32(p) == p2
	default:
		return false
	}
}

// Generate random Partition values.
//
// Interface: quick.Generator
func (Partition) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Partition.
//
// Interface: fmt.Stringer.
func (p Partition) String() string {
	return fmt.Sprintf("%d", p)
}

// Generate a random Partition value.
func Generate(rand *rand.Rand) Partition {
	return Partition(rand.Int31())
}
