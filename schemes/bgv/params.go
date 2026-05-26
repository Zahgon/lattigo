package bgv

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
)

const (
	NTTFlag = true
)

// ParametersLiteral is a literal representation of BGV parameters.  It has public
// fields and is used to express unchecked user-defined parameters literally into
// Go programs. The [NewParametersFromLiteral] function is used to generate the actual
// checked parameters from the literal representation.
//
// Users must set the polynomial degree (LogN) and the coefficient modulus, by either setting
// the Q and P fields to the desired moduli chain, or by setting the LogQ and LogP fields to
// the desired moduli sizes.
//
// Users must also specify the coefficient modulus in plaintext-space (T).
// If slot encoding is used (default), the number of slots available in the plaintext will be equal to the largest n s.t. T = 1 mod 2n.
// Otherwise, if coefficient encoding is used, the maximal number of slots available is equal to N.
//
// Optionally, users may specify the error variance (Sigma) and secrets' density (H). If left
// unset, standard default values for these field are substituted at parameter creation (see
// NewParametersFromLiteral).
type ParametersLiteral struct {
	LogN             int
	LogNthRoot       int
	Q                []uint64
	P                []uint64
	LogQ             []int `json:",omitempty"`
	LogP             []int `json:",omitempty"`
	Xe               ring.DistributionParameters
	Xs               ring.DistributionParameters
	PlaintextModulus uint64 // Plaintext modulus
}

// GetRLWEParametersLiteral returns the [rlwe.ParametersLiteral] from the target [bgv.ParametersLiteral].
// See the [ParametersLiteral] type for details on the BGV parameters.
func (p ParametersLiteral) GetRLWEParametersLiteral() rlwe.ParametersLiteral {
	_ = "STUB: not implemented"
	return *new(rlwe.ParametersLiteral)
}

// Parameters represents a parameter set for the BGV cryptosystem. Its fields are private and
// immutable. See [ParametersLiteral] for user-specified parameters.
type Parameters struct {
	rlwe.Parameters
	ringQMul *ring.Ring
	ringT    *ring.Ring
}

// NewParameters instantiate a set of BGV parameters from the generic RLWE parameters and the BGV-specific ones.
// It returns the empty parameters [Parameters]{} and a non-nil error if the specified parameters are invalid.
// See the [ParametersLiteral] type for more details on the BGV parameters.
func NewParameters(rlweParams rlwe.Parameters, t uint64) (p Parameters, err error) {
	_ = "STUB: not implemented"
	return *new(Parameters), nil
}

/* #nosec G115 -- NthRoot cannot be negative */

// Find the largest cyclotomic order enabled by T

/* #nosec G115 -- library requires 64-bit system -> int = int64 */

// NewParametersFromLiteral instantiate a set of BGV parameters from a [ParametersLiteral] specification.
// It returns the empty parameters [Parameters]{} and a non-nil error if the specified parameters are invalid.
//
// See [rlwe.NewParametersFromLiteral] for default values of the optional fields and other details on the BGV
// parameters.
func NewParametersFromLiteral(pl ParametersLiteral) (Parameters, error) {
	_ = "STUB: not implemented"
	return *new(Parameters), nil
}

// ParametersLiteral returns the [ParametersLiteral] of the target Parameters.
func (p Parameters) ParametersLiteral() ParametersLiteral {
	_ = "STUB: not implemented"
	return *new(ParametersLiteral)
}

// GetRLWEParameters returns a pointer to the underlying RLWE parameters.
func (p Parameters) GetRLWEParameters() *rlwe.Parameters { _ = "STUB: not implemented"; return nil }

// MaxDimensions returns the maximum dimension of the matrix that can be SIMD packed in a single plaintext polynomial.
func (p Parameters) MaxDimensions() ring.Dimensions {
	_ = "STUB: not implemented"
	return *new(ring.Dimensions)
}

// LogMaxDimensions returns the log2 of maximum dimension of the matrix that can be SIMD packed in a single plaintext polynomial.
func (p Parameters) LogMaxDimensions() ring.Dimensions {
	_ = "STUB: not implemented"
	return *new(ring.Dimensions)
}

// MaxSlots returns the total number of entries (slots) that a plaintext can store.
// This value is obtained by multiplying all dimensions from MaxDimensions.
func (p Parameters) MaxSlots() int { _ = "STUB: not implemented"; return 0 }

// LogMaxSlots returns the total number of entries (slots) that a plaintext can store.
// This value is obtained by summing all log dimensions from LogDimensions.
func (p Parameters) LogMaxSlots() int { _ = "STUB: not implemented"; return 0 }

// RingQMul returns a pointer to the ring of the extended basis for multiplication.
func (p Parameters) RingQMul() *ring.Ring {
	_ = "STUB: not implemented"

	// PlaintextModulus returns the plaintext coefficient modulus t.
	return nil
}

func (p Parameters) PlaintextModulus() uint64 { _ = "STUB: not implemented"; return 0 }

// LogT returns log2(plaintext coefficient modulus).
func (p Parameters) LogT() float64 { _ = "STUB: not implemented"; return 0 }

// RingT returns a pointer to the plaintext ring.
func (p Parameters) RingT() *ring.Ring {
	_ = "STUB: not implemented"

	// GaloisElementForColRotation returns the Galois element for generating the
	// automorphism phi(k): X -> X^{5^k mod 2N} mod (X^{N} + 1), which acts as a
	// column-wise cyclic rotation by k position to the left on batched plaintexts.
	//
	// Example:
	// Recall that batched plaintexts are 2xN/2 matrices, thus given the following
	// plaintext matrix:
	//
	// [a, b, c, d][e, f, g, h]
	//
	// a rotation by k=3 will change the plaintext to:
	//
	// [d, a, b, d][h, e, f, g]
	//
	// Providing a negative k will change direction of the cyclic rotation do the right.
	return nil
}

func (p Parameters) GaloisElementForColRotation(k int) uint64 { _ = "STUB: not implemented"; return 0 }

// GaloisElementForRowRotation returns the Galois element for generating the
// automorphism X -> X^{-1 mod NthRoot} mod (X^{N} + 1). This automorphism
// acts as a swapping the rows of the plaintext algebra when the plaintext
// is batched.
//
// Example:
// Recall that batched plaintexts are 2xN/2 matrices, thus given the following
// plaintext matrix:
//
// [a, b, c, d][e, f, g, h]
//
// a row rotation will change the plaintext to:
//
// [e, f, g, h][a, b, c, d]
func (p Parameters) GaloisElementForRowRotation() uint64 { _ = "STUB: not implemented"; return 0 }

// GaloisElementsForInnerSum returns the list of Galois elements necessary to apply the method
// InnerSum operation with parameters batch and n.
func (p Parameters) GaloisElementsForInnerSum(batch, n int) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// GaloisElementsForReplicate returns the list of Galois elements necessary to perform the
// Replicate operation with parameters batch and n.
func (p Parameters) GaloisElementsForReplicate(batch, n int) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// GaloisElementsForTrace returns the list of Galois elements required for the for the Trace operation.
// Trace maps X -> sum((-1)^i * X^{i*n+1}) for 2^{LogN} <= i < N.
func (p Parameters) GaloisElementsForTrace(logN int) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

// Equal compares two sets of parameters for equality.
func (p Parameters) Equal(other *Parameters) bool { _ = "STUB: not implemented"; return false }

// MarshalBinary returns a []byte representation of the parameter set.
// The representation corresponds to the JSON representation obtained
// from MarshalJSON.
func (p Parameters) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalBinary decodes a []byte into a parameter set struct.
		nil
}

func (p *Parameters) UnmarshalBinary(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON returns a JSON representation of this parameter set. See Marshal from the [encoding/json] package.
func (p Parameters) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON reads a JSON representation of a parameter set into the receiver Parameter. See Unmarshal from the [encoding/json] package.
func (p *Parameters) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (p *ParametersLiteral) UnmarshalJSON(b []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
