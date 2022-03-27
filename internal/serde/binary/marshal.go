package binary

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"time"
)

// MarshalOffset marshals the int64 offset value to the given writer.
func MarshalOffset(w io.Writer, offset int64) error {
	return writeInt64(w, offset)
}

// MarshalPartition marshals the int32 partition value to the given writer.
func MarshalPartition(w io.Writer, partition int32) error {
	return writeInt32(w, partition)
}

// MarshalTimestamp marshals the time.Time timestamp value (as int64 milliseconds) to the given writer.
func MarshalTimestamp(w io.Writer, ts time.Time) error {
	return writeInt64(w, ts.UnixMilli())
}

// MarshalTopic marshals the string topic value to the given writer.
func MarshalTopic(w io.Writer, s string) error {
	return writeString(w, s)
}

// MarshalHeaders marshals the key/val header values to the given writer.
func MarshalHeaders(w io.Writer, headers map[string]string) error {
	for key, val := range headers {
		if err := writeString(w, key); err != nil {
			return errors.Wrapf(err, "failed to marshal header key %s", key)
		}
		if err := writeString(w, val); err != nil {
			return errors.Wrapf(err, "failed to marshal header val %s", val)
		}
	}
	return nil
}

// writeString writes a length prefixed string to the io.Writer.
func writeString(w io.Writer, s string) error {
	if err := binary.Write(w, binary.LittleEndian, int32(len(s))); err != nil {
		return errors.Wrap(err, "failed to marshal string length")
	}
	wrote, err := w.Write([]byte(s))
	if err != nil {
		return errors.Wrap(err, "failed to marshal string")
	}
	if wrote != len(s) {
		return errors.Errorf("failed to marshal string. expected %d; wrote %d", len(s), wrote)
	}
	return nil
}

// writeInt32 strings a little endian int32 to the io.Writer.
func writeInt32(w io.Writer, v int32) error {
	return binary.Write(w, binary.LittleEndian, v)
}

// writeInt64 strings a little endian int64 to the io.Writer.
func writeInt64(w io.Writer, v int64) error {
	return binary.Write(w, binary.LittleEndian, v)
}
