package metadata

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestMetadata_String checks String() output is expected format.
func TestMetadata_String(t *testing.T) {
	checker := func(m Metadata) bool {
		return m.String() == fmt.Sprintf(
			"Metadata(topic=%s partition=%s offset=%s timestamp=%s)",
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

// TestMetadata_EqualsTrue checks equality between two Metadata instances.
func TestMetadata_EqualsTrue(t *testing.T) {
	checker := func(m Metadata) bool {
		return m.Equals(m)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestMetadata_EqualsFalse checks inequality between two in-equal Metadata instances.
func TestMetadata_EqualsFalse(t *testing.T) {
	checker := func(m1 Metadata) bool {
		m2 := Metadata{
			Partition: m1.Partition + 1,
			Offset:    m1.Offset + 1,
			Timestamp: m1.Timestamp,
			Topic:     m1.Topic + "foo",
		}
		return !m1.Equals(m2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
