package collector

import (
	"fmt"
	"github.com/flyingdice/proton-pack/external/kafka/producer"
	"github.com/flyingdice/proton-pack/internal/codec"
	"github.com/flyingdice/proton-pack/internal/types/generic/state/machine"
	"github.com/flyingdice/proton-pack/internal/types/streams/offset/store"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/context"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/record"
	"github.com/flyingdice/proton-pack/internal/validation"
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
)

// Collector provides metrics and offset tracking while fronting a kafka producer.
type Collector struct {
	machine  *machine.Machine[State]
	offsets  *store.Store
	producer *producer.Producer
}

// NewCollector creates and validates a new Collector.
func NewCollector(producer *producer.Producer) (*Collector, validation.ErrorGroup) {
	m, err := machine.NewMachine[State](Initial, States, Transitions)
	if err != nil {
		return nil, err
	}
	o, err := store.NewStore()
	if err != nil {
		return nil, err
	}
	c := &Collector{
		machine:  m,
		offsets:  o,
		producer: producer,
	}
	return c, c.Check()
}

// Check runs default validation checks for the collector.
func (c *Collector) Check() validation.ErrorGroup {
	return validation.RunChecks[*Collector](c, defaultChecks...)
}

// Send record to Kafka.
//
// TODO (ahawker) Properly implement.
func (c *Collector) Send(ctx context.Context, rec record.Record, keyCodec, valCodec codec.Codec) error {
	// Convert key/val into bytes.
	keyBytes := []byte{}
	valBytes := []byte{}

	// Send message request on to the producer.
	res, err := c.producer.Produce(ctx.Context, producer.Request{
		Key:       keyBytes,
		Val:       valBytes,
		Topic:     string(rec.Metadata.Topic),
		Partition: int32(rec.Metadata.Partition),
	})
	if err != nil {
		return errors.Wrapf(
			err,
			"collector failed to send message for %s/%s",
			rec.Metadata.Topic,
			rec.Metadata.Partition,
		)
	}

	// Track offsets for responses.
	if err := c.offsets.PutParts(res.Topic, res.Partition, res.Offset); err != nil {
		return errors.Wrapf(
			err,
			"collector failed to store offsets for on %s/%s",
			rec.Metadata.Topic,
			rec.Metadata.Partition,
		)
	}

	return nil
}

// Open collector to begin sending messages.
func (c *Collector) Open() error {
	return c.machine.To(Opened, func() error {
		return c.producer.Open()
	})
}

// Flush buffered messages in the collector.
func (c *Collector) Flush() error {
	return c.machine.MustBe(Opened, func() error {
		return c.producer.Flush()
	})
}

// Close the collector if it is open.
func (c *Collector) Close() error {
	return c.machine.To(Closed, func() error {
		return c.producer.Close()
	})
}

// Generate random collector values.
//
// Interface: quick.Generator
func (*Collector) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the collector.
//
// Interface: fmt.Stringer.
func (c *Collector) String() string {
	return fmt.Sprintf("Collector(state=%s offsets=%s)", c.machine, c.offsets)
}

// Generate a random collector value.
func Generate(rand *rand.Rand) *Collector {
	m, err := machine.NewMachine[State](Initial, States, Transitions)
	if err != nil {
		panic(err)
	}
	return &Collector{
		machine:  m,
		offsets:  store.Generate(rand),
		producer: producer.Generate(rand),
	}
}
