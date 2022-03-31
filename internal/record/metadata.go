package record

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/primitive/offset"
	"github.com/flyingdice/proton-pack/internal/primitive/partition"
	"github.com/flyingdice/proton-pack/internal/primitive/timestamp"
	"github.com/flyingdice/proton-pack/internal/primitive/topic"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*Metadata)(nil)
var _ quick.Generator = (*Metadata)(nil)
var _ comparison.Equaler = (*Metadata)(nil)

type Metadata struct {
	Partition partition.Partition `json:"partition"`
	Offset    offset.Offset       `json:"offset"`
	Timestamp timestamp.Timestamp `json:"timestamp"`
	Topic     topic.Topic         `json:"topic"`
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
	m := Metadata{
		Partition: partition.Generate(rand),
		Offset:    offset.Generate(rand),
		Timestamp: timestamp.Generate(rand),
		Topic:     topic.Generate(rand),
	}
	return reflect.ValueOf(m)
}

// String value of the Metadata.
//
// Interface: fmt.Stringer.
func (m Metadata) String() string {
	return fmt.Sprintf(
		"Metadata(topic=%s, partition=%d, offset=%d, timestamp=%v)",
		m.Topic,
		m.Partition,
		m.Offset,
		m.Timestamp,
	)
}
