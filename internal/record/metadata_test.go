package record

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestMetadata_String checks String() output is expected format.
func TestMetadata_String(t *testing.T) {
	checker := func(m Metadata) bool {
		return m.String() == fmt.Sprintf(
			"Metadata(topic=%s, partition=%d, offset=%d, timestamp=%v)",
			m.Topic,
			m.Partition,
			m.Offset,
			m.Timestamp,
		)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
