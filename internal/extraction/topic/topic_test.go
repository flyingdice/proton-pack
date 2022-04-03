package topic

import (
	"github.com/flyingdice/proton-pack/internal/types/streams/record/context"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/record"
	"testing"
	"testing/quick"
)

// TestRecord checks extractor uses the record topic.
func TestRecord(t *testing.T) {
	checker := func(ctx context.Context, rec record.Record) bool {
		topic, err := Record(ctx, rec)
		if err != nil {
			return false
		}
		return topic == rec.Metadata.Topic
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
