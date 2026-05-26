package ring

import (
	"math/big"
)

// RNSScalar represents a scalar value in the Ring (i.e., a degree-0 polynomial) in RNS form.
type RNSScalar []uint64

// NewRNSScalar creates a new Scalar value.
func (r *Ring) NewRNSScalar() RNSScalar { _ = "STUB: not implemented"; return *new(RNSScalar) }

// NewRNSScalarFromUInt64 creates a new Scalar initialized with value v.
func (r *Ring) NewRNSScalarFromUInt64(v uint64) (rns RNSScalar) {
	_ = "STUB: not implemented"
	return *new(RNSScalar)
}

// NewRNSScalarFromBigint creates a new Scalar initialized with value v.
func (r *Ring) NewRNSScalarFromBigint(v *big.Int) (rns RNSScalar) {
	_ = "STUB: not implemented"
	return *new(RNSScalar)
}

// MFormRNSScalar switches an RNS scalar to the Montgomery domain.
// s2 = s1<<64 mod Q
func (r *Ring) MFormRNSScalar(s1, s2 RNSScalar) { _ = "STUB: not implemented"; return }

// NegRNSScalar evaluates s2 = -s1.
func (r *Ring) NegRNSScalar(s1, s2 RNSScalar) { _ = "STUB: not implemented"; return }

// SubRNSScalar subtracts s2 to s1 and stores the result in sout.
func (r *Ring) SubRNSScalar(s1, s2, sout RNSScalar) { _ = "STUB: not implemented"; return }

// MulRNSScalar multiplies s1 and s2 and stores the result in sout.
// Multiplication is operated with Montgomery.
func (r *Ring) MulRNSScalar(s1, s2, sout RNSScalar) { _ = "STUB: not implemented"; return }

// Inverse computes the modular inverse of a scalar a expressed in a CRT decomposition.
// The inversion is done in-place and assumes that a is in Montgomery form.
func (r *Ring) Inverse(a RNSScalar) { _ = "STUB: not implemented"; return }

/* #nosec G115 -- library requires 64-bit system -> int = int64 */
