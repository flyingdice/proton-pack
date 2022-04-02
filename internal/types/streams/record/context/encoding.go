package context

import (
	"bytes"
	"encoding"
	"github.com/pkg/errors"
	"io"
)

var _ encoding.BinaryMarshaler = (*Context)(nil)
var _ encoding.BinaryUnmarshaler = (*Context)(nil)

// MarshalBinary coverts the Context instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (c Context) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = c.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// UnmarshalBinary converts the binary form to a Context instance.
//
// Interface: encoding.BinaryUnmarshaler
func (c *Context) UnmarshalBinary(data []byte) error {
	return c.UnmarshalBinaryReader(bytes.NewReader(data))
}

// MarshalBinaryWriter populates the io.Writer with Context fields
// in its binary form.
func (c Context) MarshalBinaryWriter(w io.Writer) error {
	if err := c.Metadata.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal context metadata")
	}
	if err := c.Headers.MarshalBinaryWriter(w); err != nil {
		return errors.Wrap(err, "failed to marshal context headers")
	}
	return nil
}

// UnmarshalBinaryReader populates Context fields from an io.Reader
// returning the binary form.
func (c *Context) UnmarshalBinaryReader(r io.Reader) error {
	if err := (&c.Metadata).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal context metadata")
	}
	if err := (&c.Headers).UnmarshalBinaryReader(r); err != nil {
		return errors.Wrap(err, "failed to unmarshal context headers")
	}
	return nil
}
