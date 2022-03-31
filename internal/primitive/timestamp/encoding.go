package timestamp

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"encoding/json"
	"io"
	"time"
)

var _ encoding.BinaryMarshaler = (*Timestamp)(nil)
var _ encoding.BinaryUnmarshaler = (*Timestamp)(nil)
var _ json.Marshaler = (*Timestamp)(nil)
var _ json.Unmarshaler = (*Timestamp)(nil)

// MarshalBinary converts the Timestamp instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (t Timestamp) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = t.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// MarshalJSON converts the Timestamp instance to JSON form.
//
// Interface: json.Marshaler
func (t Timestamp) MarshalJSON() (data []byte, err error) {
	data, err = json.Marshal(t.Time)
	return
}

// UnmarshalBinary converts the binary form to a Timestamp instance.
//s
// Interface: encoding.BinaryUnmarshaler
func (t *Timestamp) UnmarshalBinary(data []byte) (err error) {
	return t.UnmarshalBinaryReader(bytes.NewReader(data))
}

// UnmarshalJSON converts the JSON form to a Timestamp instance.
//
// Interface: json.Unmarshal
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var timestamp time.Time
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}
	*t = Timestamp{
		Time: timestamp,
	}
	return nil
}

// MarshalBinaryWriter populates the io.Writer with Timestamp data
// in its binary form.
func (t Timestamp) MarshalBinaryWriter(w io.Writer) (err error) {
	return binary.Write(w, binary.LittleEndian, t.UnixMilli())
}

// UnmarshalBinaryReader populates Timestamp from an io.Reader
// returning the binary form.
func (t *Timestamp) UnmarshalBinaryReader(r io.Reader) error {
	var timestampMilli int64
	if err := binary.Read(r, binary.LittleEndian, &timestampMilli); err != nil {
		return err
	}
	*t = Timestamp{
		Time: time.UnixMilli(timestampMilli),
	}
	return nil
}
