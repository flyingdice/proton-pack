package metadata

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/types/kafka/offset"
	"github.com/flyingdice/proton-pack/internal/types/kafka/partition"
	"github.com/flyingdice/proton-pack/internal/types/kafka/timestamp"
	"github.com/flyingdice/proton-pack/internal/types/kafka/topic"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*Metadata)(nil)
var _ quick.Generator = (*Metadata)(nil)
var _ comparison.Equaler = (*Metadata)(nil)
var _ validation.Checker = (*Metadata)(nil)

type Metadata struct {
	Partition partition.Partition `json:"partition"`
	Offset    offset.Offset       `json:"offset"`
	Timestamp timestamp.Timestamp `json:"timestamp"`
	Topic     topic.Topic         `json:"topic"`
}

// NewMetadata creates and validates a new Metadata from the given fields.
func NewMetadata(
	p partition.Partition,
	o offset.Offset,
	ts timestamp.Timestamp,
	t topic.Topic,
) (Metadata, validation.ErrorGroup) {
	m := Metadata{p, o, ts, t}
	return m, m.Check()
}

// Check runs default validation checks for the Metadata.
func (m Metadata) Check() validation.ErrorGroup {
	return validation.RunChecks[Metadata](m, defaultChecks...)
}

// Equals compares two Metadata instances for equality.
//
// Interface: comparison.Equaler
func (m Metadata) Equals(v any) bool {
	switch m2 := v.(type) {
	case Metadata:
		return m.Partition.Equals(m2.Partition) &&
			m.Offset.Equals(m2.Offset) &&
			m.Topic.Equals(m2.Topic) &&
			m.Timestamp.Equals(m.Timestamp)
	default:
		return false
	}
}

// Generate random Metadata values.
//
// Interface: quick.Generator
func (Metadata) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Metadata.
//
// Interface: fmt.Stringer.
func (m Metadata) String() string {
	return fmt.Sprintf(
		"Metadata(topic=%s partition=%s offset=%s timestamp=%s)",
		m.Topic,
		m.Partition,
		m.Offset,
		m.Timestamp,
	)
}

// Generate a random Metadata value.
func Generate(rand *rand.Rand) Metadata {
	return Metadata{
		Partition: partition.Generate(rand),
		Offset:    offset.Generate(rand),
		Timestamp: timestamp.Generate(rand),
		Topic:     topic.Generate(rand),
	}
}
