package ckks

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/ring"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

type Float interface {
	float64 | complex128 | *big.Float | *bignum.Complex
}

// FloatSlice is an empty interface whose goal is to
// indicate that the expected input should be []Float.
// See Float for information on the type constraint.
type FloatSlice interface {
}

// Complex is an empty interface whose goal is to
// indicate that the expected input should be a complex number.
type Complex any

// GaloisGen is an integer of order N/2 modulo M and that spans Z_M with the integer -1.
// The j-th ring automorphism takes the root zeta to zeta^(5^j).
const GaloisGen uint64 = ring.GaloisGen

// Encoder is a type that implements the encoding and decoding interface for the CKKS scheme. It provides methods to encode/decode
// []complex128/[]*[bignum.Complex] and []float64/[]*[big.Float] types into/from Plaintext types.
//
// Two different encodings domains are provided:
//
//   - Coefficients: The coefficients are directly embedded on the plaintext. This encoding only allows to encode []float64/[]*[big.Float] slices,
//     but of size up to N (N being the ring degree) and does not preserve the point-wise multiplication. A ciphertext multiplication will result
//     in a negacyclic polynomial convolution in the plaintext domain. This encoding does not provide native slot cyclic rotation.
//     Other operations, like addition or constant multiplication, behave as usual.
//
//   - Slots: The coefficients are first subjected to a special Fourier transform before being embedded in the plaintext by using Coeffs encoding.
//     This encoding can embed []complex128/[]*[bignum.Complex] and []float64/[]*[big.Float] slices of size at most N/2 (N being the ring degree) and
//     leverages the convolution property of the DFT to preserve point-wise complex multiplication in the plaintext domain, i.e. a ciphertext
//     multiplication will result in an element-wise multiplication in the plaintext domain. It also enables cyclic rotations on plaintext slots.
//     Other operations, like constant multiplication, behave as usual. It is considered the default encoding method for CKKS.
//
// The figure bellow illustrates the relationship between these two encodings:
//
//	                                                    Z_Q[X]/(X^N+1)
//		Coefficients: ---------------> Real^{N} ---------> Plaintext
//	                                      |
//	                                      |
//		Slots: Complex^{N/2} -> iDFT -----┘
type Encoder struct {
	parameters Parameters

	prec uint

	m        int
	rotGroup []int

	roots interface{}

	// Pools used to recycle large objects.
	BuffBigIntPool  structs.BufferPool[*[]*big.Int]
	BuffComplexPool structs.BufferPool[Complex]
	poolQ           *ring.BufferPool
}

// NewEncoder creates a new [Encoder] from the target parameters.
// Optional field `precision` can be given. If precision is empty
// or <= 53, then float64 and complex128 types will be used to
// perform the encoding. Else *[big.Float] and *[bignum.Complex] will be used.
func NewEncoder(parameters Parameters, precision ...uint) (ecd *Encoder) {
	_ = "STUB: not implemented"

	/* #nosec G115 -- library requires 64-bit system -> int = int64 */
	return nil
}

// Prec returns the precision in bits used by the target Encoder.
// A precision <= 53 will use float64, else *[big.Float].
func (ecd Encoder) Prec() uint { _ = "STUB: not implemented"; return 0 }

func (ecd Encoder) GetParameters() Parameters { _ = "STUB: not implemented"; return *new(Parameters) }

func (ecd Encoder) GetRLWEParameters() rlwe.Parameters {
	_ = "STUB: not implemented"
	return *new(rlwe.Parameters)
}

// Encode encodes a [FloatSlice] on the target plaintext.
// Encoding is done at the level and scale of the plaintext.
// Encoding domain is done according to the metadata of the plaintext.
// User must ensure that 1 <= len(values) <= 2^pt.LogMaxDimensions < 2^logN.
// The imaginary part will be discarded if ringType == ring.ConjugateInvariant.
func (ecd Encoder) Encode(values interface{}, pt *rlwe.Plaintext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Decode decodes the input plaintext on a new FloatSlice.
func (ecd Encoder) Decode(pt *rlwe.Plaintext, values interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// DecodePublic decodes the input plaintext on a [FloatSlice].
// It adds, before the decoding step (i.e. in the Ring) noise that follows the given distribution parameters.
// If the underlying ringType is [ring.ConjugateInvariant], the imaginary part (and its related error) are zero.
func (ecd Encoder) DecodePublic(pt *rlwe.Plaintext, values FloatSlice, logprec float64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Embed is a generic method to encode a [FloatSlice] on the target polyOut.
// This method it as the core of the slot encoding.
// Values are encoded according to the provided metadata.
// Accepted polyOut.(type) are ringqp.Poly and ring.Poly.
// The imaginary part will be discarded if ringType == ring.ConjugateInvariant.
func (ecd Encoder) Embed(values interface{}, metadata *rlwe.MetaData, polyOut interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// embedDouble encode a FloatSlice on polyOut using FFT with complex128 arithmetic.
// Values are encoded according to the provided metadata.
// Accepted polyOut.(type) are [ringqp.Poly] and [ring.Poly].
func (ecd Encoder) embedDouble(values FloatSlice, metadata *rlwe.MetaData, polyOut interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Zeroes all other values

// IFFT

// Maps Y = X^{N/n} -> X and quantizes.

// embedArbitrary encode a FloatSlice on polyOut using FFT with *bignum.Complex arithmetic.
// Values are encoded according to the provided metadata.
// Accepted polyOut.(type) are [ringqp.Poly] and [ring.Poly].
func (ecd Encoder) embedArbitrary(values FloatSlice, metadata *rlwe.MetaData, polyOut interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Zeroes all other values

// Maps Y = X^{N/n} -> X and quantizes.

// plaintextToComplex maps a CRT polynomial to a complex valued [FloatSlice].
func (ecd Encoder) plaintextToComplex(level int, scale rlwe.Scale, logSlots int, p ring.Poly, values FloatSlice) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// plaintextToFloat maps a CRT polynomial to a real valued [FloatSlice].
func (ecd Encoder) plaintextToFloat(level int, scale rlwe.Scale, logSlots int, p ring.Poly, values FloatSlice) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// decodePublic decode a plaintext to a [FloatSlice].
// The method will add a flooding noise before the decoding process following the defined distribution if it is not nil.
func (ecd Encoder) decodePublic(pt *rlwe.Plaintext, values FloatSlice, logprec float64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// 2^logprec

// Adds/Subtracts 0.5

// Round = floor +/- 0.5

// Real

// Adds/Subtracts 0.5

// Round = floor +/- 0.5

// Imag

// Adds/Subtracts 0.5

// Round = floor +/- 0.5

// IFFT evaluates the special 2^{LogN}-th encoding discrete Fourier transform on [FloatSlice].
func (ecd Encoder) IFFT(values FloatSlice, logN int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// FFT evaluates the special 2^{LogN}-th decoding discrete Fourier transform on [FloatSlice].
func (ecd Encoder) FFT(values FloatSlice, logN int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// polyToComplexNoCRT decodes a single-level CRT poly on a complex valued [FloatSlice].
func polyToComplexNoCRT(coeffs []uint64, values FloatSlice, scale rlwe.Scale, logSlots int, isreal bool, ringQ *ring.Ring) (err error) {
	_ = "STUB: not implemented"
	return

	/* #nosec G115 -- library requires 64-bit system -> int = int64 */
	nil
}

// [X]/(X^N+1) to [X+X^-1]/(X^N+1)

/* #nosec G115 -- Q - c <= 61 bits */

/* #nosec G115 -- c <= 61 bits */

/* #nosec G115 -- Q - c <= 61 bits */

/* #nosec G115 -- c <= 61 bits */

// polyToComplexNoCRT decodes a multiple-level CRT poly on a complex valued [FloatSlice].
func polyToComplexCRT(poly ring.Poly, bigintCoeffs []*big.Int, values FloatSlice, scale rlwe.Scale, logSlots int, isreal bool, ringQ *ring.Ring) (err error) {
	_ = "STUB: not implemented"

	/* #nosec G115 -- library requires 64-bit system -> int = int64 */
	return nil
}

// [X]/(X^N+1) to [X+X^-1]/(X^N+1)

// [X]/(X^N+1) to [X+X^-1]/(X^N+1)

// polyToFloatCRT decodes a multiple-level CRT poly on a real valued [FloatSlice].
func (ecd *Encoder) polyToFloatCRT(p ring.Poly, values FloatSlice, scale rlwe.Scale, logSlots int, r *ring.Ring) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Centers the value around the current modulus

// polyToFloatNoCRT decodes a single-level CRT poly on a real valued [FloatSlice].
func (ecd *Encoder) polyToFloatNoCRT(coeffs []uint64, values FloatSlice, scale rlwe.Scale, logSlots int, r *ring.Ring) (err error) {
	_ = "STUB: not implemented"
	return nil
}

/* #nosec G115 -- Q - coeffs[i] <= 61 bits */

/* #nosec G115 -- coeffs[i] <= 61 bits */

/* #nosec G115 -- Q - coeffs[i] <= 61 bits */

/* #nosec G115 -- coeffs[i] <= 61 bits */
