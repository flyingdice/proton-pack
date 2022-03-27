package record

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

var _ fmt.Stringer = (*Metadata)(nil)
var _ quick.Generator = (*Metadata)(nil)

type Metadata struct {
	Partition int32     `json:"partition"`
	Offset    int64     `json:"offset"`
	Timestamp time.Time `json:"timestamp"`
	Topic     string    `json:"topic"`
}

// Equals compares two Metadata instances for equality.
//
// Interface: comparison.Equaler
func (m Metadata) Equals(o Metadata) bool {
	return m.Partition == o.Partition &&
		m.Offset == o.Offset &&
		m.Topic == o.Topic &&
		m.Timestamp.UnixMilli() == m.Timestamp.UnixMilli()
}

// Generate random Metadata values.
//
// Interface: quick.Generator
func (Metadata) Generate(rand *rand.Rand, size int) reflect.Value {
	faker.SetRandomSource(rand)
	m := Metadata{
		Partition: rand.Int31(),
		Offset:    rand.Int63(),
		Timestamp: time.Unix(0, 0).Add(time.Duration(rand.Int63())),
		Topic:     faker.Word(),
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
