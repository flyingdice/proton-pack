package offset

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"encoding/json"
	"io"
)

var _ encoding.BinaryMarshaler = (*Offset)(nil)
var _ encoding.BinaryUnmarshaler = (*Offset)(nil)
var _ json.Marshaler = (*Offset)(nil)
var _ json.Unmarshaler = (*Offset)(nil)

// MarshalBinary converts the Offset instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (o Offset) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = o.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// MarshalJSON converts the Offset instance to JSON form.
//
// Interface: json.Marshaler
func (o Offset) MarshalJSON() (data []byte, err error) {
	data, err = json.Marshal(int64(o))
	return
}

// UnmarshalBinary converts the binary form to a Offset instance.
//
// Interface: encoding.BinaryUnmarshaler
func (o *Offset) UnmarshalBinary(data []byte) (err error) {
	return o.UnmarshalBinaryReader(bytes.NewReader(data))
}

// UnmarshalJSON converts the JSON form to a Offset instance.
//
// Interface: json.Unmarshal
func (o *Offset) UnmarshalJSON(data []byte) (err error) {
	var offset int64
	if err := json.Unmarshal(data, &offset); err != nil {
		return err
	}
	*o, err = NewOffset(offset)
	return err
}

// MarshalBinaryWriter populates the io.Writer with Offset data
// in its binary form.
func (o Offset) MarshalBinaryWriter(w io.Writer) (err error) {
	return binary.Write(w, binary.LittleEndian, int64(o))
}

// UnmarshalBinaryReader populates Offset from an io.Reader
// returning the binary form.
func (o *Offset) UnmarshalBinaryReader(r io.Reader) (err error) {
	var offset int64
	if err := binary.Read(r, binary.LittleEndian, &offset); err != nil {
		return err
	}
	*o, err = NewOffset(offset)
	return err
}
