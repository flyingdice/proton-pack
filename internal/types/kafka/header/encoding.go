package header

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"github.com/pkg/errors"
	"io"
	"unicode/utf8"
)

var _ encoding.BinaryMarshaler = (*Header)(nil)
var _ encoding.BinaryUnmarshaler = (*Header)(nil)

// MarshalBinary converts the Header instance to binary form.
//
// Interface: encoding.BinaryMarshaler
func (h Header) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	if err = h.MarshalBinaryWriter(&buf); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

// UnmarshalBinary converts the binary form to a Header instance.
//
// Interface: encoding.BinaryUnmarshaler
func (h *Header) UnmarshalBinary(data []byte) (err error) {
	return h.UnmarshalBinaryReader(bytes.NewReader(data))
}

// MarshalBinaryWriter populates the io.Writer with Header data
// in its binary form.
func (h Header) MarshalBinaryWriter(w io.Writer) (err error) {
	keyLength := utf8.RuneCountInString(h.Key)
	valLength := len(h.Val)

	// Key.
	if err := binary.Write(w, binary.LittleEndian, int32(keyLength)); err != nil {
		return errors.Wrap(err, "failed to marshal header key length")
	}
	wrote, err := w.Write([]byte(h.Key))
	if err != nil {
		return errors.Wrap(err, "failed to marshal header key")
	}
	if wrote != keyLength {
		return errors.Errorf("failed to marshal header key. expected %d; wrote %d", keyLength, wrote)
	}

	// Val.
	// If empty/nil/undefined, store -1 for the value (without a length)
	// as the sentinel value.
	if h.Val == nil || len(h.Val) == 0 {
		if err := binary.Write(w, binary.LittleEndian, int32(-1)); err != nil {
			return errors.Wrap(err, "failed to marshal header val sentinel")
		}
	} else {
		if err := binary.Write(w, binary.LittleEndian, int32(valLength)); err != nil {
			return errors.Wrap(err, "failed to marshal header val length")
		}
		wrote, err = w.Write(h.Val)
		if err != nil {
			return errors.Wrap(err, "failed to marshal header val")
		}
		if wrote != valLength {
			return errors.Errorf("failed to marshal header val. expected %d; wrote %d", valLength, wrote)
		}
	}

	return nil
}

// UnmarshalBinaryReader populates Header from an io.Reader
// returning the binary form.
func (h *Header) UnmarshalBinaryReader(r io.Reader) (err error) {
	var keyLength, valLength int32

	// Key.
	if err := binary.Read(r, binary.LittleEndian, &keyLength); err != nil {
		return errors.Wrap(err, "failed to unmarshal header key length")
	}
	keyBuf := make([]byte, keyLength)
	if _, err := io.ReadFull(r, keyBuf); err != nil {
		return errors.Wrap(err, "failed to unmarshal header key")
	}

	// Val.
	if err := binary.Read(r, binary.LittleEndian, &valLength); err != nil {
		return errors.Wrap(err, "failed to unmarshal header val length")
	}
	valBuf := make([]byte, valLength)
	if _, err := io.ReadFull(r, valBuf); err != nil {
		return errors.Wrap(err, "failed to unmarshal header val")
	}

	*h, err = NewHeader(string(keyBuf), valBuf)
	return err
}
