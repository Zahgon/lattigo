package bgv

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

type Integer interface {
	int64 | uint64
}

// IntegerSlice is an empty interface whose goal is to
// indicate that the expected input should be []Integer.
// See Integer for information on the type constraint.
type IntegerSlice interface {
}

// GaloisGen is an integer of order N=2^d modulo M=2N and that spans Z_M with the integer -1.
// The j-th ring automorphism takes the root zeta to zeta^(5^j).
const GaloisGen uint64 = ring.GaloisGen

// Encoder is a structure that stores the parameters to encode values on a plaintext in a SIMD (Single-Instruction Multiple-Data) fashion.
type Encoder struct {
	parameters Parameters

	indexMatrix []uint64

	// BuffBigIntPool is allocated in the case when the degree of RingT is smaller
	// than the degree of RingQ (gap > 1), hence a more involved conversion
	// between the two structures is necessary.
	// The size of an object returned from the pool MaxSlots() elements.
	BuffBigIntPool structs.BufferPool[*[]*big.Int]
	poolQ          *ring.BufferPool
	poolT          *ring.BufferPool

	paramsQP []ring.ModUpConstants
	qHalf    []*big.Int

	tInvModQ []*big.Int
}

// NewEncoder creates a new [Encoder] from the provided parameters.
func NewEncoder(parameters Parameters) *Encoder { _ = "STUB: not implemented"; return nil }

// create pools for polys in ringQ and ringT

// we use the same backing pool only if ringQ and ringT have the same dimension

func permuteMatrix(logN int) (perm []uint64) { _ = "STUB: not implemented"; return nil }

/* #nosec G115 -- library requires 64-bit system -> int = int64 */

// = (pow-1)/2

// GetRLWEParameters returns the underlying [rlwe.Parameters] of the target object.
func (ecd Encoder) GetRLWEParameters() *rlwe.Parameters { _ = "STUB: not implemented"; return nil }

// Encode encodes values on a pre-allocated plaintext. The `values` must be of type [IntegerSlice] or be a [ring.Poly] from params.RingT.
// If `values` is of type [ring.Poly], then pt.IsBatched is set to false and the polynomial is encoded directly (i.e., scaled).
// If `values` is of type [IntegerSlice], then the encoding depends respects the pt.IsBatched flag:
//   - If pt.IsBatched=false, then values are interpreted as the coefficients of a polynomial and encored as above.
//   - If pt.IsBatched=true, then values are encoded in a SIMD fashion on n slots, where n is the largest value satisfying PlaintextModulus = 1 mod 2n.
func (ecd Encoder) Encode(values interface{}, pt *rlwe.Plaintext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If IsBatched=false, the plaintext can have dimension N (dimension of the ciphertext ring)

/* #nosec G115 -- type conversion purposefully used */

/* #nosec G115 -- converted value is ensured to be positive */

// EncodeRingT encodes an [IntegerSlice] at the given scale on a polynomial pT with coefficients modulo the plaintext modulus PlaintextModulus.
func (ecd Encoder) EncodeRingT(values IntegerSlice, scale rlwe.Scale, pT ring.Poly) (err error) {
	_ = "STUB: not implemented"
	return nil
}

/* #nosec G115 -- c cannot be negative */

/* #nosec G115 -- converted value is ensured to be positive */

// Zeroes the non-mapped coefficients

// INTT on the Y = X^{N/n}

// EmbedScale is a generic method to encode an IntegerSlice on [ringqp.Poly] or *[ring.Poly].
// If scaleUp is true, then the values will to be multiplied by PlaintextModulus^{-1} mod Q after being encoded on the polynomial.
// Encoding is done according to the metadata.
// Accepted polyOut.(type) are a ringqp.Poly and *ring.Poly
func (ecd Encoder) EmbedScale(values IntegerSlice, scaleUp bool, metadata *rlwe.MetaData, polyOut interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Maps Y = X^{N/n} -> X and quantizes.

func (ecd Encoder) Embed(values interface{}, metadata *rlwe.MetaData, polyOut interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// DecodeRingT decodes a polynomial pT with coefficients modulo the plaintext modulu PlaintextModulus on an InterSlice at the given scale.
func (ecd Encoder) DecodeRingT(pT ring.Poly, scale rlwe.Scale, values IntegerSlice) (err error) {
	_ = "STUB: not implemented"
	return nil
}

/* #nosec G115 -- PlaintextModulus <= 61 bits */

/* #nosec G115 -- values <= 61 bits */

// RingT2Q takes pT in base PlaintextModulus and writes it in base Q[level] on pQ.
// If scaleUp is true, multiplies the values of pQ by PlaintextModulus^{-1} mod Q[level].
func (ecd Encoder) RingT2Q(level int, scaleUp bool, pT, pQ ring.Poly) {
	_ = "STUB: not implemented"
	return
}

// RingQ2T takes pQ in base Q[level] and writes it in base PlaintextModulus on pT.
// If scaleUp is true, the values of pQ are multiplied by PlaintextModulus mod Q[level]
// before being converted into the base PlaintextModulus.
func (ecd Encoder) RingQ2T(level int, scaleDown bool, pQ, pT ring.Poly) {
	_ = "STUB: not implemented"
	return
}

// buffB only used in this block, can be put back in the pool:

// Decode decodes a [Plaintext] into values of type [IntegerSlice] or [ring.Poly].
// If pt.IsBatched=true, then values must be a [IntegerSlice] and the plaintext is decoded in a SIMD fashion from n slots, where n is the largest value satisfying PlaintextModulus = 1 mod 2n.
// If pt.IsBatched=false, then values can be either a [ring.Poly] from Parameters.RingT or an [IntegerSlice] and the plaintext is decoded as the coefficients of a polynomial
func (ecd Encoder) Decode(pt *rlwe.Plaintext, values interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

/* #nosec G115 -- PlaintextModulus <= 61 bits */

/* #nosec G115 -- values <= 61 bits */
