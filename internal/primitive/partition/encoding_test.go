package partition

import (
	"testing"
	"testing/quick"
)

// TestPartition_BinaryEncoding checks Partition equality when encoding/decoding to/from binary form.
func TestPartition_BinaryEncoding(t *testing.T) {
	checker := func(p1 Partition) bool {
		buf, err := p1.MarshalBinary()
		if err != nil {
			t.Error(err)
		}
		var p2 Partition
		if err := p2.UnmarshalBinary(buf); err != nil {
			t.Error(err)
		}
		return p1.Equals(p2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestPartition_JSONEncoding checks Partition equality when encoding/decoding to/from JSON form.
func TestPartition_JSONEncoding(t *testing.T) {
	checker := func(p1 Partition) bool {
		buf, err := p1.MarshalJSON()
		if err != nil {
			t.Error(err)
		}
		var p2 Partition
		if err := p2.UnmarshalJSON(buf); err != nil {
			t.Error(err)
		}
		return p1.Equals(p2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
