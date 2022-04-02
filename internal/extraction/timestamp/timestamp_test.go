package timestamp

import (
	"github.com/flyingdice/proton-pack/internal/types/streams/record/context"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/record"
	"github.com/matryer/is"
	"github.com/pkg/errors"
	"reflect"
	"testing"
	"testing/quick"
)

// TestRecord checks extractor uses the record timestamp.
func TestRecord(t *testing.T) {
	checker := func(ctx context.Context, rec record.Record) bool {
		ts, err := Record(ctx, rec)
		if err != nil {
			t.Error(err)
			return false
		}
		return ts == rec.Metadata.Timestamp
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestWallClock checks extractor uses the context clock.
func TestWallClock(t *testing.T) {
	checker := func(ctx context.Context, rec record.Record) bool {
		ts, err := WallClock(ctx, rec)
		if err != nil {
			t.Error(err)
			return false
		}
		return ts.Equals(ctx.Clock.Now())
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestFactory checks correct extractor funcs are returned by name.
func TestFactory(t *testing.T) {
	tests := []struct {
		name   string
		exp    Extractor
		expErr error
	}{
		{name: "wallclock", exp: WallClock},
		{name: "record", exp: Record},
		{name: "foobar", expErr: errors.New("unknown timestamp extractor 'foobar'")},
	}

	for _, tc := range tests {
		assert := is.New(t)

		got, err := Factory(tc.name)
		if err != nil {
			assert.Equal(err.Error(), tc.expErr.Error())
		} else {
			x := reflect.ValueOf(tc.exp).Pointer()
			y := reflect.ValueOf(got).Pointer()
			if x != y {
				t.Fatalf("expected extractor: %v, got: %v", x, y)
			}
		}
	}
}
