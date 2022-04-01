package header

import (
	"bytes"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
)

var _ fmt.Stringer = (*Header)(nil)
var _ comparison.Equaler = (*Header)(nil)
var _ validation.Checker = (*Header)(nil)

// Header represents the key/value pair for a Kafka message.
type Header struct {
	Key string `json:"key"`
	Val []byte `json:"val"`
}

// NewHeader creates and validates a new Header from the given key/val.
func NewHeader(key string, val []byte) (Header, validation.ErrorGroup) {
	h := Header{key, val}
	return h, h.Check()
}

// Check runs default validation checks for the Header.
func (h Header) Check() validation.ErrorGroup {
	return validation.Validate[Header](h, defaultChecks...)
}

// Equals compares two Header instances for equality.
//
// Interface: comparison.Equaler
func (h Header) Equals(v any) bool {
	switch h2 := v.(type) {
	case Header:
		return h.Key == h2.Key && bytes.Compare(h.Val, h2.Val) == 0
	default:
		return false
	}
}

// Generate random Header values.
//
// Interface: quick.Generator
func (Header) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Header.
//
// Interface: fmt.Stringer.
func (h Header) String() string {
	return fmt.Sprintf("Header(key=%s val=%s)", h.Key, h.Val)
}

// Generate a random Header value.
func Generate(rand *rand.Rand) Header {
	faker.SetRandomSource(rand)
	return Header{Key: faker.Word(), Val: []byte(faker.Word())}
}
