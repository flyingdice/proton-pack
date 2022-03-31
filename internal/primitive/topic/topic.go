package topic

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
)

var _ fmt.Stringer = (*Topic)(nil)
var _ comparison.Equaler = (*Topic)(nil)

// Topic represents a kafka topic.
type Topic string

// NewTopic creates and validates a new Topic from the given string.
func NewTopic(s string) (Topic, error) {
	t := Topic(s)
	return t, validation.Validate[Topic](t, defaultChecks...)
}

// Equals compares two Topic instances for equality.
//
// Interface: comparison.Equaler
func (t Topic) Equals(v any) bool {
	switch t2 := v.(type) {
	case Topic:
		return t == t2
	default:
		return false
	}
}

// Generate random Topic values.
//
// Interface: quick.Generator
func (Topic) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Topic.
//
// Interface: fmt.Stringer.
func (t Topic) String() string {
	return string(t)
}

// Generate a random Topic value.
func Generate(rand *rand.Rand) Topic {
	faker.SetRandomSource(rand)
	return Topic(faker.Word())
}