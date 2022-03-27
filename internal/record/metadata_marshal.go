package record

import (
	"bytes"
	"encoding"
	"github.com/flyingdice/proton-pack/internal/serde/binary"
	"github.com/pkg/errors"
	"io"
)

var _ encoding.BinaryMarshaler = (*Metadata)(nil)
var _ encoding.BinaryUnmarshaler = (*Metadata)(nil)

// MarshalBinary converts the Metadata instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (m Metadata) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer

	if err = m.MarshalBinaryWriter(&buf); err != nil {
		return
	}

	data = buf.Bytes()
	return
}

// UnmarshalBinary converts the binary form to a Metadata instance.
//
// Interface: encoding.BinaryUnmarshaler
func (m *Metadata) UnmarshalBinary(data []byte) error {
	return m.UnmarshalBinaryReader(bytes.NewReader(data))
}

// MarshalBinaryWriter populates the io.Writer with Metadata fields
// in its binary form.
func (m Metadata) MarshalBinaryWriter(w io.Writer) error {
	if err := binary.MarshalTimestamp(w, m.Timestamp); err != nil {
		return errors.Wrap(err, "failed to marshal metadata timestamp")
	}
	if err := binary.MarshalOffset(w, m.Offset); err != nil {
		return errors.Wrap(err, "failed to marshal metadata offset")
	}
	if err := binary.MarshalTopic(w, m.Topic); err != nil {
		return errors.Wrapf(err, "failed to marshal metadata topic")
	}
	if err := binary.MarshalPartition(w, m.Partition); err != nil {
		return errors.Wrap(err, "failed to marshal metadata partition")
	}
	return nil
}

// UnmarshalBinaryReader populates Metadata fields from an io.Reader
// returning the binary form.
func (m *Metadata) UnmarshalBinaryReader(r io.Reader) error {
	if err := binary.UnmarshalTimestamp(r, &m.Timestamp); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata timestamp")
	}
	if err := binary.UnmarshalOffset(r, &m.Offset); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata offset")
	}
	if err := binary.UnmarshalTopic(r, &m.Topic); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata topic")
	}
	if err := binary.UnmarshalPartition(r, &m.Partition); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata partition")
	}
	return nil
}
