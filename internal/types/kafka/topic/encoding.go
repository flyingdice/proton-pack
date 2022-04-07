package topic

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
)

var _ encoding.BinaryMarshaler = (*Topic)(nil)
var _ encoding.BinaryUnmarshaler = (*Topic)(nil)
var _ json.Marshaler = (*Topic)(nil)
var _ json.Unmarshaler = (*Topic)(nil)

// MarshalBinary converts the Topic instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (t Topic) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = t.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// MarshalJSON converts the Topic instance to JSON form.
//
// Interface: json.Marshaler
func (t Topic) MarshalJSON() (data []byte, err error) {
	data, err = json.Marshal(string(t))
	return
}

// UnmarshalBinary converts the binary form to a Topic instance.
//
// Interface: encoding.BinaryUnmarshaler
func (t *Topic) UnmarshalBinary(data []byte) (err error) {
	return t.UnmarshalBinaryReader(bytes.NewReader(data))
}

// UnmarshalJSON converts the JSON form to a Topic instance.
//
// Interface: json.Unmarshal
func (t *Topic) UnmarshalJSON(data []byte) (err error) {
	var topic string
	if err := json.Unmarshal(data, &topic); err != nil {
		return err
	}
	*t, err = New(topic)
	return err
}

// MarshalBinaryWriter populates the io.Writer with Topic data
// in its binary form.
func (t Topic) MarshalBinaryWriter(w io.Writer) (err error) {
	if err := binary.Write(w, binary.LittleEndian, int32(len(t))); err != nil {
		return errors.Wrap(err, "failed to marshal topic length")
	}
	wrote, err := w.Write([]byte(t))
	if err != nil {
		return errors.Wrap(err, "failed to marshal topic")
	}
	if wrote != len(t) {
		return errors.Errorf("failed to marshal topic. expected %d; wrote %d", len(t), wrote)
	}
	return nil
}

// UnmarshalBinaryReader populates Topic from an io.Reader
// returning the binary form.
func (t *Topic) UnmarshalBinaryReader(r io.Reader) (err error) {
	var length int32
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return err
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	*t, err = New(string(buf))
	return err
}
