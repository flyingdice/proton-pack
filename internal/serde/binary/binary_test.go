package binary

import (
	"bytes"
	"testing"
	"testing/quick"
	"time"
)

var quickCheckConfig *quick.Config = nil

// TestOffset_BinaryEncoding checks equality for marshal/unmarshal of offset values.
func TestOffset_BinaryEncoding(t *testing.T) {
	checker := func(o1 int64) bool {
		var buf bytes.Buffer
		var o2 int64

		if err := MarshalOffset(&buf, o1); err != nil {
			t.Error(err)
		}
		if err := UnmarshalOffset(&buf, &o2); err != nil {
			t.Error(err)
		}

		return o1 == o2
	}
	if err := quick.Check(checker, quickCheckConfig); err != nil {
		t.Error(err)
	}
}

// TestTimestamp_BinaryEncoding checks equality for marshal/unmarshal of timestamp values.
func TestTimestamp_BinaryEncoding(t *testing.T) {
	checker := func(millis int64) bool {
		var buf bytes.Buffer
		var t1, t2 time.Time

		t1 = time.UnixMilli(millis)

		if err := MarshalTimestamp(&buf, t1); err != nil {
			t.Error(err)
		}
		if err := UnmarshalTimestamp(&buf, &t2); err != nil {
			t.Error(err)
		}

		return t1.UnixMilli() == t2.UnixMilli()
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestPartition_BinaryEncoding checks equality for marshal/unmarshal of partition values.
func TestPartition_BinaryEncoding(t *testing.T) {
	checker := func(p1 int32) bool {
		var buf bytes.Buffer
		var p2 int32

		if err := MarshalPartition(&buf, p1); err != nil {
			t.Error(err)
		}
		if err := UnmarshalPartition(&buf, &p2); err != nil {
			t.Error(err)
		}

		return p1 == p2
	}
	if err := quick.Check(checker, quickCheckConfig); err != nil {
		t.Error(err)
	}
}

// TestTopic_BinaryEncoding checks equality for marshal/unmarshal of topic values.
func TestTopic_BinaryEncoding(t *testing.T) {
	checker := func(t1 string) bool {
		var buf bytes.Buffer
		var t2 string

		if err := MarshalTopic(&buf, t1); err != nil {
			t.Error(err)
		}
		if err := UnmarshalTopic(&buf, &t2); err != nil {
			t.Error(err)
		}

		return t1 == t2
	}
	if err := quick.Check(checker, quickCheckConfig); err != nil {
		t.Error(err)
	}
}
