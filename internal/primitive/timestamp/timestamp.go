package timestamp

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math"
	"math/rand"
	"reflect"
	"time"
)

var _ fmt.Stringer = (*Timestamp)(nil)
var _ comparison.Equaler = (*Timestamp)(nil)
var _ validation.Checker = (*Timestamp)(nil)

// Timestamp represents a kafka message timestamp.
type Timestamp struct {
	time.Time
}

// NewTimestamp creates and validates a new Timestamp from the given time.Time.
func NewTimestamp(t time.Time) (Timestamp, validation.ErrorGroup) {
	ts := Timestamp{t}
	return ts, ts.Check()
}

// Check runs default validation checks for the Timestamp.
func (t Timestamp) Check() validation.ErrorGroup {
	return validation.Validate[Timestamp](t, defaultChecks...)
}

// Equals compares two Timestamp instances for equality.
//
// Interface: comparison.Equaler
func (t Timestamp) Equals(v any) bool {
	switch t2 := v.(type) {
	case Timestamp:
		return t.UnixMilli() == t2.UnixMilli()
	case time.Time:
		return t.UnixMilli() == t2.UnixMilli()
	case int64:
		return t.UnixMilli() == t2
	default:
		return false
	}
}

// Generate random Timestamp values.
//
// Interface: quick.Generator
func (Timestamp) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Timestamp.
//
// Interface: fmt.Stringer.
func (t Timestamp) String() string {
	return fmt.Sprintf("%s", t.Format(time.RFC3339))
}

// Generate a random Timestamp value.
func Generate(rand *rand.Rand) Timestamp {
	t := time.Unix(0, 0).Add(time.Duration(rand.Int63n(math.MaxInt64)))
	return Timestamp{Time: t}
}
