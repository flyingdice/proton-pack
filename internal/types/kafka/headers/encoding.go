package headers

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"github.com/flyingdice/proton-pack/internal/types/kafka/header"
	"github.com/pkg/errors"
	"io"
)

var _ encoding.BinaryMarshaler = (*Headers)(nil)
var _ encoding.BinaryUnmarshaler = (*Headers)(nil)

// MarshalBinary converts the Headers instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (h Headers) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = h.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// UnmarshalBinary converts the binary form to a Headers instance.
//
// Interface: encoding.BinaryUnmarshaler
func (h *Headers) UnmarshalBinary(data []byte) (err error) {
	return h.UnmarshalBinaryReader(bytes.NewReader(data))
}

// MarshalBinaryWriter populates the io.Writer with Headers data
// in its binary form.
func (h Headers) MarshalBinaryWriter(w io.Writer) (err error) {
	// Length.
	if err := binary.Write(w, binary.LittleEndian, int32(len(h))); err != nil {
		return errors.Wrap(err, "failed to marshal headers length")
	}

	// Individual headers.
	for _, hdr := range h {
		if err := hdr.MarshalBinaryWriter(w); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalBinaryReader populates Headers from an io.Reader
// returning the binary form.
func (h *Headers) UnmarshalBinaryReader(r io.Reader) (err error) {
	var headers []header.Header
	var length int32

	// Length.
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return errors.Wrap(err, "failed to unmarshal headers length")
	}

	// Individual headers.
	for i := 0; i < int(length); i++ {
		var hdr header.Header

		if err := (&hdr).UnmarshalBinaryReader(r); err != nil {
			return err
		}

		headers = append(headers, hdr)
	}

	*h, err = NewHeaders(headers)
	return err
}
