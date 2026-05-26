package rlwe

import (
	"io"

	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

// ElementInterface is a common interface for [Ciphertext] and [Plaintext] types.
type ElementInterface[T ring.Poly | ringqp.Poly] interface {
	N() int
	LogN() int
	El() *Element[T]
	Degree() int
	Level() int
}

// Element is a generic struct to store a vector of T along with some metadata.
type Element[T ring.Poly | ringqp.Poly] struct {
	*MetaData
	Value structs.Vector[T]
}

// NewElement allocates a new Element[ring.Poly].
func NewElement(params ParameterProvider, degree int, levelQ ...int) *Element[ring.Poly] {
	_ = "STUB: not implemented"
	return nil
}

// NewElementExtended allocates a new Element[ringqp.Poly].
func NewElementExtended(params ParameterProvider, degree, levelQ, levelP int) *Element[ringqp.Poly] {
	_ = "STUB: not implemented"
	return nil
}

// NewElementAtLevelFromPoly constructs a new [Element] at a specific level
// where the message is set to the passed poly. No checks are performed on poly and
// the returned [Element] will share its backing array of coefficients.
// Returned [Element]'s [MetaData] is nil.
func NewElementAtLevelFromPoly(level int, poly []ring.Poly) (*Element[ring.Poly], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// N returns the ring degree used by the target element.
func (op Element[T]) N() int { _ = "STUB: not implemented"; return 0 }

// Sanity check

// LogN returns the log2 of the ring degree used by the target element.
func (op Element[T]) LogN() int {
	_ = "STUB: not implemented"
	/* #nosec G115 -- N is ensured to be greater than 0 */ return 0
}

// Equal performs a deep equal.
func (op Element[T]) Equal(other *Element[T]) bool { _ = "STUB: not implemented"; return false }

// Degree returns the degree of the target element.
func (op Element[T]) Degree() int { _ = "STUB: not implemented"; return 0 }

// Level returns the level of the target element.
func (op Element[T]) Level() int { _ = "STUB: not implemented"; return 0 }

func (op Element[T]) LevelQ() int { _ = "STUB: not implemented"; return 0 }

// Sanity check

func (op Element[T]) LevelP() int { _ = "STUB: not implemented"; return 0 }

// Sanity check

func (op *Element[T]) El() *Element[T] {
	_ = "STUB: not implemented"

	// Resize resizes the degree of the target element.
	// Sets the NTT flag of the added poly equal to the NTT flag
	// to the poly at degree zero.
	return nil
}

func (op *Element[T]) Resize(degree, level int) { _ = "STUB: not implemented"; return }

// Sanity check

// CopyNew creates a deep copy of the object and returns it.
func (op Element[T]) CopyNew() *Element[T] { _ = "STUB: not implemented"; return nil }

// Copy copies opCopy on op, up to the capacity of op (similarely to copy([]byte, []byte)).
func (op *Element[T]) Copy(opCopy *Element[T]) { _ = "STUB: not implemented"; return }

// GetSmallestLargest returns the provided element that has the smallest degree as a first
// returned value and the largest degree as second return value. If the degree match, the
// order is the same as for the input.
func GetSmallestLargest[T ring.Poly | ringqp.Poly](el0, el1 *Element[T]) (smallest, largest *Element[T], sameDegree bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// PopulateElementRandom creates a new [Element] with random coefficients.
func PopulateElementRandom(prng sampling.PRNG, params ParameterProvider, ct *Element[ring.Poly]) {
	_ = "STUB: not implemented"
	return
}

// SwitchCiphertextRingDegreeNTT changes the ring degree of ctIn to the one of opOut.
// Maps Y^{N/n} -> X^{N} or X^{N} -> Y^{N/n}.
// If the ring degree of opOut is larger than the one of ctIn, then the ringQ of opOut
// must be provided (otherwise, a nil pointer).
// The ctIn must be in the NTT domain and opOut will be in the NTT domain.
func SwitchCiphertextRingDegreeNTT(ctIn *Element[ring.Poly], ringQLargeDim *ring.Ring, opOut *Element[ring.Poly]) {
	_ = "STUB: not implemented"
	return
}

// SwitchCiphertextRingDegree changes the ring degree of ctIn to the one of opOut.
// Maps Y^{N/n} -> X^{N} or X^{N} -> Y^{N/n}.
// If the ring degree of opOut is larger than the one of ctIn, then the ringQ of ctIn
// must be provided (otherwise, a nil pointer).
func SwitchCiphertextRingDegree(ctIn, opOut *Element[ring.Poly]) { _ = "STUB: not implemented"; return }

// BinarySize returns the serialized size of the object in bytes.
func (op Element[T]) BinarySize() (size int) {
	_ = "STUB: not implemented"
	// Whether or not there is metadata
	return 0
}

// WriteTo writes the object on an [io.Writer]. It implements the [io.WriterTo]
// interface, and will write exactly object.BinarySize() bytes on w.
//
// Unless w implements the [buffer.Writer] interface (see lattigo/utils/buffer/writer.go),
// it will be wrapped into a [bufio.Writer]. Since this requires allocations, it
// is preferable to pass a [buffer.Writer] directly:
//
//   - When writing multiple times to a [io.Writer], it is preferable to first wrap the
//     [io.Writer] in a pre-allocated [bufio.Writer].
//   - When writing to a pre-allocated var b []byte, it is preferable to pass
//     buffer.NewBuffer(b) as w (see lattigo/utils/buffer/buffer.go).
func (op Element[T]) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom reads on the object from an [io.Writer]. It implements the
// io.ReaderFrom interface.
//
// Unless r implements the [buffer.Reader] interface (see see lattigo/utils/buffer/reader.go),
// it will be wrapped into a [bufio.Reader]. Since this requires allocation, it
// is preferable to pass a buffer.Reader directly:
//
//   - When reading multiple values from a [io.Reader], it is preferable to first
//     first wrap [io.Reader] in a pre-allocated [bufio.Reader].
//   - When reading from a var b []byte, it is preferable to pass a buffer.NewBuffer(b)
//     as w (see lattigo/utils/buffer/buffer.go).
func (op *Element[T]) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MarshalBinary encodes the object into a binary form on a newly allocated slice of bytes.
func (op Element[T]) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalBinary decodes a slice of bytes generated by
// [Element.MarshalBinary] or [Element.WriteTo] on the object.
func (op *Element[T]) UnmarshalBinary(p []byte) (err error) { _ = "STUB: not implemented"; return nil }
