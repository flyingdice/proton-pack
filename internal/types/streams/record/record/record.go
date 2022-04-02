package record

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/types/kafka/headers"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/key"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/metadata"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/val"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*Record)(nil)
var _ quick.Generator = (*Record)(nil)
var _ comparison.Equaler = (*Record)(nil)
var _ validation.Checker = (*Record)(nil)

type Record struct {
	Headers  headers.Headers   `json:"headers"`
	Metadata metadata.Metadata `json:"metadata"`
	Key      key.Key           `json:"key"`
	Val      val.Val           `json:"val"`
}

// NewRecord creates and validates new Record from the given values.
func NewRecord(
	k key.Key,
	v val.Val,
	m metadata.Metadata,
	h headers.Headers,
) (Record, validation.ErrorGroup) {
	r := Record{h, m, k, v}
	return r, r.Check()
}

// Check runs default validation checks for the Record.
func (r Record) Check() validation.ErrorGroup {
	return validation.RunChecks[Record](r, defaultChecks...)
}

// Equals compares two Record instances for equality.
//
// Interface: comparison.Equaler
func (r Record) Equals(v any) bool {
	switch r2 := v.(type) {
	case Record:
		return reflect.DeepEqual(r.Key, r2.Key) &&
			reflect.DeepEqual(r.Val, r2.Val) &&
			r.Headers.Equals(r2.Headers) &&
			r.Metadata.Equals(r2.Metadata)
	default:
		return false
	}
}

// Generate random Record values.
//
// Interface: quick.Generator
func (Record) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Record.
//
// Interface: fmt.Stringer.
func (r Record) String() string {
	return fmt.Sprintf(
		"Record(key=%v val=%v metadata=%v headers=%v)",
		r.Key,
		r.Val,
		r.Metadata,
		r.Headers,
	)
}

// Generate a random Record value.
func Generate(rand *rand.Rand) Record {
	return Record{
		Headers:  headers.Generate(rand),
		Metadata: metadata.Generate(rand),
		Key:      key.Generate(rand),
		Val:      val.Generate(rand),
	}
}
