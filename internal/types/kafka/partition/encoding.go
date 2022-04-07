package partition

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"encoding/json"
	"io"
)

var _ encoding.BinaryMarshaler = (*Partition)(nil)
var _ encoding.BinaryUnmarshaler = (*Partition)(nil)
var _ json.Marshaler = (*Partition)(nil)
var _ json.Unmarshaler = (*Partition)(nil)

// MarshalBinary converts the Partition instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (p Partition) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = p.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// MarshalJSON converts the Partition instance to JSON form.
//
// Interface: json.Marshaler
func (p Partition) MarshalJSON() (data []byte, err error) {
	data, err = json.Marshal(int32(p))
	return
}

// UnmarshalBinary converts the binary form to a Partition instance.
//
// Interface: encoding.BinaryUnmarshaler
func (p *Partition) UnmarshalBinary(data []byte) (err error) {
	return p.UnmarshalBinaryReader(bytes.NewReader(data))
}

// UnmarshalJSON converts the JSON form to a Partition instance.
//
// Interface: json.Unmarshal
func (p *Partition) UnmarshalJSON(data []byte) (err error) {
	var partition int32
	if err := json.Unmarshal(data, &partition); err != nil {
		return err
	}
	*p, err = New(partition)
	return err
}

// MarshalBinaryWriter populates the io.Writer with Partition data
// in its binary form.
func (p Partition) MarshalBinaryWriter(w io.Writer) (err error) {
	return binary.Write(w, binary.LittleEndian, int32(p))
}

// UnmarshalBinaryReader populates Partition from an io.Reader
// returning the binary form.
func (p *Partition) UnmarshalBinaryReader(r io.Reader) (err error) {
	var partition int32
	if err := binary.Read(r, binary.LittleEndian, &partition); err != nil {
		return err
	}
	*p, err = New(partition)
	return err
}
