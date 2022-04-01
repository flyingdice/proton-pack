package metadata

import (
	"encoding/json"
	"testing"
	"testing/quick"
)

// TestMetadata_BinaryEncoding checks equality when encoding/decoding to/from binary form.
func TestMetadata_BinaryEncoding(t *testing.T) {
	checker := func(m1 Metadata) bool {
		buf, err := m1.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		m2 := &Metadata{}
		if err := m2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
		}
		return m1.Equals(*m2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestMetadata_JSONEncoding checks equality when encoding/decoding to/from json form.
func TestMetadata_JSONEncoding(t *testing.T) {
	checker := func(m1 Metadata) bool {
		buf, err := json.Marshal(m1)
		if err != nil {
			t.Error(err)
		}
		m2 := Metadata{}
		if err := json.Unmarshal(buf, &m2); err != nil {
			t.Error(err)
		}
		return m1.Equals(m2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
