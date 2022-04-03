package topic_partition

import (
	"bytes"
	"encoding"
	"github.com/pkg/errors"
	"io"
)

var _ encoding.BinaryMarshaler = (*TopicPartition)(nil)
var _ encoding.BinaryUnmarshaler = (*TopicPartition)(nil)

// MarshalBinary coverts the TopicPartition instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (tp TopicPartition) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = tp.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// UnmarshalBinary converts the binary form to a TopicPartition instance.
//
// Interface: encoding.BinaryUnmarshaler
func (tp *TopicPartition) UnmarshalBinary(data []byte) error {
	return tp.UnmarshalBinaryReader(bytes.NewReader(data))
}

// MarshalBinaryWriter populates the io.Writer with TopicPartition fields
// in its binary form.
func (tp TopicPartition) MarshalBinaryWriter(w io.Writer) error {
	if err := tp.Topic.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal topic_partition topic")
	}
	if err := tp.Partition.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal topic_partition partition")
	}
	return nil
}

// UnmarshalBinaryReader populates TopicPartition fields from an io.Reader
// returning the binary form.
func (tp *TopicPartition) UnmarshalBinaryReader(r io.Reader) error {
	if err := (&tp.Topic).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal topic_partition topic")
	}
	if err := (&tp.Partition).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal topic_partition partition")
	}
	return nil
}
