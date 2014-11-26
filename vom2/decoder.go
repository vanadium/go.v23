package vom2

import (
	"io"
	"reflect"

	"veyron.io/veyron/veyron2/vdl/valconv"
	"veyron.io/veyron/veyron2/verror"
)

var (
	errDecodeNil         = verror.BadArgf("invalid decode into nil interface{}")
	errDecodeNilRawValue = verror.BadArgf("invalid decode into nil *RawValue")
)

// Decoder manages the receipt and unmarshaling of typed values from the other
// side of a connection.
type Decoder struct {
	dec decoder
}

type decoder interface {
	Decode(target valconv.Target) error
	DecodeRaw(raw *RawValue) error
	Ignore() error
}

// NewDecoder returns a new Decoder that reads from the given reader.  The
// Decoder understands all formats generated by the Encoder.
func NewDecoder(r io.Reader) (*Decoder, error) {
	buf, types := newDecbuf(r), newDecoderTypes()
	magic, err := buf.PeekByte()
	if err != nil {
		return nil, verror.BadProtocolf("error reading magic byte %v", err)
	}
	if magic != binaryMagicByte {
		return nil, verror.BadProtocolf("bad magic byte, got %x, want %x", magic, binaryMagicByte)
	}
	buf.Skip(1)
	return &Decoder{newBinaryDecoder(buf, types)}, nil
}

// Decode reads the next value from the reader and stores it in value v.
// The type of v need not exactly match the type of the originally encoded
// value; decoding succeeds as long as the values are compatible.
//
//   Types that are special-cased, only for v:
//     *RawValue  - Store raw (uninterpreted) bytes in v.
//
//   Types that are special-cased, recursively throughout v:
//     *val.Value    - Decode into v.
//     reflect.Value - Decode into v, which must be settable.
//
// Decoding into a RawValue captures the value in a raw form, which may be
// subsequently passed to an Encoder for transcoding.
//
// Decode(nil) always returns an error.  Use Ignore() to ignore the next value.
func (d *Decoder) Decode(v interface{}) error {
	switch tv := v.(type) {
	case nil:
		return errDecodeNil
	case *RawValue:
		if tv == nil {
			return errDecodeNilRawValue
		}
		return d.dec.DecodeRaw(tv)
	}
	target, err := valconv.ReflectTarget(reflect.ValueOf(v))
	if err != nil {
		return err
	}
	return d.dec.Decode(target)
}

// Ignore ignores the next value from the reader.
func (d *Decoder) Ignore() error {
	return d.dec.Ignore()
}
