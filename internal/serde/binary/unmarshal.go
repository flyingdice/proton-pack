package binary

import (
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"time"
)

// UnmarshalOffset reads the int64 offset value from the given reader.
func UnmarshalOffset(r io.Reader, offset *int64) error {
	return readInt64(r, offset)
}

// UnmarshalPartition reads the int32 partition value from the given reader.
func UnmarshalPartition(r io.Reader, partition *int32) error {
	return readInt32(r, partition)
}

// UnmarshalTimestamp reads the int64 timestamp value (milliseconds) from the given reader.
func UnmarshalTimestamp(r io.Reader, ts *time.Time) error {
	var millis int64
	if err := readInt64(r, &millis); err != nil {
		return err
	}
	*ts = time.UnixMilli(millis)
	return nil
}

// UnmarshalTopic reads the string topic value from the given reader.
func UnmarshalTopic(r io.Reader, topic *string) error {
	return readString(r, topic)
}

// UnmarshalHeaders reads header key/val pairs from the given reader until EOF.
func UnmarshalHeaders(r io.Reader, headers map[string]string) error {
	for {
		var key, val string
		if err := readString(r, &key); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return errors.Wrap(err, "failed to unmarshal header key")
		}
		if err := readString(r, &val); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return errors.Wrap(err, "failed to unmarshal header val")
		}
		headers[key] = val
	}
	return nil
}

// readString reads a length prefixed string from the io.Reader.
func readString(r io.Reader, s *string) error {
	var length int32
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return err
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	*s = string(buf)
	return nil
}

// readString reads an int32 value from the io.Reader.
func readInt32(r io.Reader, v *int32) error {
	return binary.Read(r, binary.LittleEndian, v)
}

// readString reads an int64 value from the io.Reader.
func readInt64(r io.Reader, v *int64) error {
	return binary.Read(r, binary.LittleEndian, v)
}
