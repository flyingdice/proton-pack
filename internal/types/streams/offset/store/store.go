package store

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/types/kafka/offset"
	"github.com/flyingdice/proton-pack/internal/types/kafka/partition"
	"github.com/flyingdice/proton-pack/internal/types/kafka/topic"
	"github.com/flyingdice/proton-pack/internal/types/kafka/topic_partition"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*Store)(nil)
var _ quick.Generator = (*Store)(nil)
var _ validation.Checker = (*Store)(nil)

// Store holds mapping between topic partitions and offsets.
//
// This is used to track highest/most-recent offsets for each topic partition
// as the flow through the system.
type Store struct {
	store map[topic_partition.TopicPartition]offset.Offset
}

// NewStore creates and validates a new Store.
func NewStore() (*Store, validation.ErrorGroup) {
	s := &Store{make(map[topic_partition.TopicPartition]offset.Offset)}
	return s, s.Check()
}

// Check runs default validation checks for the store.
func (s *Store) Check() validation.ErrorGroup {
	return validation.RunChecks[*Store](s, defaultChecks...)
}

// PutParts converts and stores primitive types to topic partition -> offset.
func (s *Store) PutParts(topic_ string, partition_ int32, offset_ int64) error {
	t, err := topic.NewTopic(topic_)
	if err != nil {
		return err
	}
	p, err := partition.NewPartition(partition_)
	if err != nil {
		return err
	}
	o, err := offset.New(offset_)
	if err != nil {
		return err
	}
	tp, err := topic_partition.NewTopicPartition(t, p)
	if err != nil {
		return err
	}
	return s.Put(tp, o)
}

// Put stores topic partition -> offset.
func (s *Store) Put(tp topic_partition.TopicPartition, offset offset.Offset) error {
	s.store[tp] = offset
	return nil
}

// Generate random Store values.
//
// Interface: quick.Generator
func (*Store) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Store.
//
// Interface: fmt.Stringer.
func (s *Store) String() string {
	return fmt.Sprintf("Store(len=%d)", len(s.store))
}

// Generate a random Header value.
func Generate(rand *rand.Rand) *Store {
	store := make(map[topic_partition.TopicPartition]offset.Offset)
	for i := 0; i < rand.Intn(10-1)+1; i++ {
		store[topic_partition.Generate(rand)] = offset.Generate(rand)
	}
	return &Store{store}
}
