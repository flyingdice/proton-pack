package timestamp

import (
	"github.com/flyingdice/proton-pack/internal/types/kafka/timestamp"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/context"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/record"
	"github.com/pkg/errors"
)

// Extractor represents a function that extracts a timestamp out of context/record
// flowing through the system.
type Extractor func(ctx context.Context, rec record.Record) (timestamp.Timestamp, error)

// WallClock uses the current wall clock timestamp for "processing time" semantics.
//
// If you need "event time" semantics, use Record.
func WallClock(ctx context.Context, rec record.Record) (timestamp.Timestamp, error) {
	return timestamp.NewTimestamp(ctx.Clock.Now())
}

// Record uses the embedded timestamp of the record for "event time" semantics.
//
// If you need "processing time" semantics, use WallClock.
func Record(ctx context.Context, rec record.Record) (timestamp.Timestamp, error) {
	return rec.Metadata.Timestamp, nil
}

// Factory returns an Extractor for the given name.
func Factory(name string) (Extractor, error) {
	switch name {
	case "wallclock":
		return WallClock, nil
	case "record":
		return Record, nil
	default:
		return nil, errors.Errorf("unknown timestamp extractor '%s'", name)
	}
}
