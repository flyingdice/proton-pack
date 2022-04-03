package topic

import (
	"github.com/flyingdice/proton-pack/internal/types/kafka/topic"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/context"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/record"
)

// Extractor represents a function that extracts a topic out of context/record
// flowing through the system.
type Extractor func(ctx context.Context, rec record.Record) (string, error)

// Record uses the topic from the current record.
func Record(ctx context.Context, rec record.Record) (topic.Topic, error) {
	return rec.Metadata.Topic, nil
}
