package metadata

import (
	"bytes"
	"encoding"
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
func (m Metadata) MarshalBinaryWriter(w io.Writer) (err error) {
	if err := m.Timestamp.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal metadata timestamp")
	}
	if err := m.Offset.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal metadata offset")
	}
	if err := m.Topic.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal metadata topic")
	}
	if err := m.Partition.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal metadata partition")
	}
	return nil
}

// UnmarshalBinaryReader populates Metadata fields from an io.Reader
// returning the binary form.
func (m *Metadata) UnmarshalBinaryReader(r io.Reader) error {
	if err := (&m.Timestamp).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata timestamp")
	}
	if err := (&m.Offset).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata offset")
	}
	if err := (&m.Topic).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata topic")
	}
	if err := (&m.Partition).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal metadata partition")
	}
	return nil
}
