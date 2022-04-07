package topic_partition

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/types/kafka/partition"
	"github.com/flyingdice/proton-pack/internal/types/kafka/topic"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*TopicPartition)(nil)
var _ quick.Generator = (*TopicPartition)(nil)
var _ comparison.Equaler = (*TopicPartition)(nil)
var _ validation.Checker = (*TopicPartition)(nil)

// TopicPartition represents an individual Partition of a Topic.
type TopicPartition struct {
	Topic     topic.Topic         `json:"topic"`
	Partition partition.Partition `json:"partition"`
}

// New creates and validates a new TopicPartition from the given Topic/Partition.
func New(t topic.Topic, p partition.Partition) (TopicPartition, validation.ErrorGroup) {
	tp := TopicPartition{t, p}
	return tp, tp.Check()
}

// Check runs default validation checks for the TopicPartition.
func (tp TopicPartition) Check() validation.ErrorGroup {
	return validation.RunChecks[TopicPartition](tp, defaultChecks...)
}

// Equals compares two TopicPartition instances for equality.
//
// Interface: comparison.Equaler
func (tp TopicPartition) Equals(v any) bool {
	switch tp2 := v.(type) {
	case TopicPartition:
		return tp.Topic.Equals(tp2.Topic) && tp.Partition.Equals(tp2.Partition)
	default:
		return false
	}
}

// Generate random TopicPartition values.
//
// Interface: quick.Generator
func (TopicPartition) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the TopicPartition.
//
// Interface: fmt.Stringer.
func (tp TopicPartition) String() string {
	return fmt.Sprintf("%s-%s", tp.Topic, tp.Partition)
}

// Generate a random TopicPartition value.
func Generate(rand *rand.Rand) TopicPartition {
	return TopicPartition{
		Topic:     topic.Generate(rand),
		Partition: partition.Generate(rand),
	}
}
