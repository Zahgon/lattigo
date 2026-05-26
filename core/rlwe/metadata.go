package rlwe

import (
	"io"

	"github.com/tuneinsight/lattigo/v6/ring"
)

// MetaData is a struct storing metadata.
type MetaData struct {
	PlaintextMetaData
	CiphertextMetaData
}

// CopyNew returns a copy of the target.
func (m MetaData) CopyNew() *MetaData { _ = "STUB: not implemented"; return nil }

func (m *MetaData) Equal(other *MetaData) (res bool) { _ = "STUB: not implemented"; return false }

// BinarySize returns the size in bytes that the object once marshalled into a binary form.
func (m MetaData) BinarySize() int { _ = "STUB: not implemented"; return 0 }

// WriteTo writes the object on an [io.Writer]. It implements the [io.WriterTo]
// interface, and will write exactly object.BinarySize() bytes on w.
func (m MetaData) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFrom reads on the object from an [io.Writer]. It implements the
// [io.ReaderFrom] interface.
//
// Unless r implements the [buffer.Reader] interface (see see lattigo/utils/buffer/reader.go),
// it will be wrapped into a [bufio.Reader]. Since this requires allocation, it
// is preferable to pass a [buffer.Reader] directly:
//
//   - When reading multiple values from a [io.Reader], it is preferable to first
//     first wrap [io.Reader] in a pre-allocated [bufio.Reader].
//   - When reading from a var b []byte, it is preferable to pass a buffer.NewBuffer(b)
//     as w (see lattigo/utils/buffer/buffer.go).
func (m *MetaData) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (m MetaData) MarshalJSON() (p []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m MetaData) MarshalBinary() (p []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *MetaData) UnmarshalJSON(p []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (m *MetaData) UnmarshalBinary(p []byte) (err error) { _ = "STUB: not implemented"; return nil }

// PlaintextMetaData is a struct storing metadata related to the plaintext.
type PlaintextMetaData struct {
	// Scale is the scaling factor of the plaintext.
	Scale Scale

	// LogDimensions is the Log2 of the 2D plaintext matrix dimensions.
	LogDimensions ring.Dimensions

	// IsBatched is a flag indicating if the underlying plaintext is encoded
	// in such a way that product in R[X]/(X^N+1) acts as a point-wise multiplication
	// in the plaintext space.
	IsBatched bool

	// IsBitReversed is a flag indicating if the underlying plaintext is
	// bit-reversed. This can be true for both batch and non-batched plaintexts.
	IsBitReversed bool
}

// Slots returns the total number of slots that the plaintext holds.
func (m PlaintextMetaData) Slots() int { _ = "STUB: not implemented"; return 0 }

// LogSlots returns the log2 of the total number of slots that the plaintext holds.
func (m PlaintextMetaData) LogSlots() int { _ = "STUB: not implemented"; return 0 }

// LogScale returns log2(scale).
func (m PlaintextMetaData) LogScale() float64 { _ = "STUB: not implemented"; return 0 }

func (m *PlaintextMetaData) Equal(other *PlaintextMetaData) (res bool) {
	_ = "STUB: not implemented"
	return false
}

// BinarySize returns the size in bytes that the object once marshalled into a binary form.
func (m PlaintextMetaData) BinarySize() int { _ = "STUB: not implemented"; return 0 }

// WriteTo writes the object on an [io.Writer]. It implements the [io.WriterTo]
// interface, and will write exactly object.BinarySize() bytes on w.
//
// Unless w implements the [buffer.Writer] interface (see lattigo/utils/buffer/writer.go),
// it will be wrapped into a [bufio.Writer]. Since this requires allocations, it
// is preferable to pass a [buffer.Writer] directly:
//
//   - When writing multiple times to a [io.Writer], it is preferable to first wrap the
//     io.Writer in a pre-allocated [bufio.Writer].
//   - When writing to a pre-allocated var b []byte, it is preferable to pass
//     buffer.NewBuffer(b) as w (see lattigo/utils/buffer/buffer.go).
func (m PlaintextMetaData) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom reads on the object from an [io.Writer]. It implements the
// [io.ReaderFrom] interface.
//
// Unless r implements the [buffer.Reader] interface (see see lattigo/utils/buffer/reader.go),
// it will be wrapped into a [bufio.Reader]. Since this requires allocation, it
// is preferable to pass a [buffer.Reader] directly:
//
//   - When reading multiple values from a [io.Reader], it is preferable to first
//     first wrap [io.Reader] in a pre-allocated [bufio.Reader].
//   - When reading from a var b []byte, it is preferable to pass a buffer.NewBuffer(b)
//     as w (see lattigo/utils/buffer/buffer.go).
func (m *PlaintextMetaData) ReadFrom(r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m PlaintextMetaData) MarshalJSON() (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* #nosec G115 -- Rows and Cols cannot be negative if valid */

// MarshalBinary encodes the object into a binary form on a newly allocated slice of bytes.
func (m PlaintextMetaData) MarshalBinary() (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *PlaintextMetaData) UnmarshalJSON(p []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

/* #nosec G115 -- logRows and logCols are < 256 if valid */

// UnmarshalBinary decodes a slice of bytes generated by
// [PlaintextMetaData.MarshalBinary] or [PlaintextMetaData.WriteTo] on the object.
func (m *PlaintextMetaData) UnmarshalBinary(p []byte) (err error) {
	_ = "STUB: not implemented"
	return nil

	// CiphertextMetaData is a struct storing metadata related to the ciphertext.
}

type CiphertextMetaData struct {
	// IsNTT is a flag indicating if the ciphertext is in the NTT domain.
	IsNTT bool
	// IsMontgomery is a flag indicating if the ciphertext is in the Montgomery domain.
	IsMontgomery bool
}

// Equal returns true if two MetaData structs are identical.
func (m *CiphertextMetaData) Equal(other *CiphertextMetaData) (res bool) {
	_ = "STUB: not implemented"
	return false
}

// BinarySize returns the size in bytes that the object once marshalled into a binary form.
func (m *CiphertextMetaData) BinarySize() int {
	_ = "STUB: not implemented"

	// WriteTo writes the object on an [io.Writer]. It implements the [io.WriterTo]
	// interface, and will write exactly object.BinarySize() bytes on w.
	//
	// Unless w implements the [buffer.Writer] interface (see lattigo/utils/buffer/writer.go),
	// it will be wrapped into a [bufio.Writer]. Since this requires allocations, it
	// is preferable to pass a [buffer.Writer] directly:
	//
	//   - When writing multiple times to a [io.Writer], it is preferable to first wrap the
	//     io.Writer in a pre-allocated [bufio.Writer].
	//   - When writing to a pre-allocated var b []byte, it is preferable to pass
	//     buffer.NewBuffer(b) as w (see lattigo/utils/buffer/buffer.go).
	return 0
}

func (m *CiphertextMetaData) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom reads on the object from an [io.Writer]. It implements the
// [io.ReaderFrom] interface.
//
// Unless r implements the [buffer.Reader] interface (see see lattigo/utils/buffer/reader.go),
// it will be wrapped into a [bufio.Reader]. Since this requires allocation, it
// is preferable to pass a [buffer.Reader] directly:
//
//   - When reading multiple values from a [io.Reader], it is preferable to first
//     first wrap [io.Reader] in a pre-allocated [bufio.Reader].
//   - When reading from a var b []byte, it is preferable to pass a buffer.NewBuffer(b)
//     as w (see lattigo/utils/buffer/buffer.go).
func (m *CiphertextMetaData) ReadFrom(r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m CiphertextMetaData) MarshalJSON() (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalBinary encodes the object into a binary form on a newly allocated slice of bytes.
func (m CiphertextMetaData) MarshalBinary() (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *CiphertextMetaData) UnmarshalJSON(p []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalBinary decodes a slice of bytes generated by
// [CiphertextMetaData.MarshalBinary] or [CiphertextMetaData.WriteTo] on the object.
func (m *CiphertextMetaData) UnmarshalBinary(p []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func hexconv(x string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
