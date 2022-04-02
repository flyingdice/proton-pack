package record

import (
	"bytes"
	"encoding"
	"github.com/pkg/errors"
	"io"
)

var _ encoding.BinaryMarshaler = (*Record)(nil)
var _ encoding.BinaryUnmarshaler = (*Record)(nil)

// MarshalBinary coverts the Record instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (r Record) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = r.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// UnmarshalBinary converts the binary form to a Record instance.
//
// Interface: encoding.BinaryUnmarshaler
func (r *Record) UnmarshalBinary(data []byte) error {
	return r.UnmarshalBinaryReader(bytes.NewReader(data))
}

// MarshalBinaryWriter populates the io.Writer with Record fields
// in its binary form.
func (r Record) MarshalBinaryWriter(w io.Writer) error {
	if err := r.Metadata.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal record metadata")
	}
	if err := r.Headers.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal record headers")
	}
	//if err := r.Key.MarshalBinaryWriter(w); err != nil {
	//	return errors.Wrap(err, "failed to marshal record key")
	//}
	//if err := r.Val.MarshalBinaryWriter(w); err != nil {
	//	return errors.Wrap(err, "failed to marshal record val")
	//}
	return nil
}

// UnmarshalBinaryReader populates Record fields from an io.Reader
// returning the binary form.
func (r *Record) UnmarshalBinaryReader(rd io.Reader) error {
	if err := (&r.Metadata).UnmarshalBinaryReader(rd); err != nil {
		return errors.Wrap(err, "failed to unmarshal record metadata")
	}
	if err := (&r.Headers).UnmarshalBinaryReader(rd); err != nil {
		return errors.Wrap(err, "failed to unmarshal record headers")
	}
	//if err := (&r.Key).UnmarshalBinaryReader(rd); err != nil {
	//	return errors.Wrap(err, "failed to unmarshal record key")
	//}
	//if err := (&r.Val).UnmarshalBinaryReader(rd); err != nil {
	//	return errors.Wrap(err, "failed to unmarshal record val")
	//}
	return nil
}
