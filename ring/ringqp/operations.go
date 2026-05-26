package ringqp

import (
	"github.com/tuneinsight/lattigo/v6/ring"
)

// Add adds p1 to p2 coefficient-wise and writes the result on p3.
func (r Ring) Add(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// AddLazy adds p1 to p2 coefficient-wise and writes the result on p3 without modular reduction.
func (r Ring) AddLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// Sub subtracts p2 to p1 coefficient-wise and writes the result on p3.
func (r Ring) Sub(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// Neg negates p1 coefficient-wise and writes the result on p2.
func (r Ring) Neg(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// NewRNSScalar creates a new Scalar value (i.e., a degree-0 polynomial) in the RingQP.
func (r Ring) NewRNSScalar() ring.RNSScalar { _ = "STUB: not implemented"; return *new(ring.RNSScalar) }

// NewRNSScalarFromUInt64 creates a new Scalar in the RingQP initialized with value v.
func (r Ring) NewRNSScalarFromUInt64(v uint64) ring.RNSScalar {
	_ = "STUB: not implemented"
	return *new(ring.RNSScalar)
}

// SubRNSScalar subtracts s2 to s1 and stores the result in sout.
func (r Ring) SubRNSScalar(s1, s2, sout ring.RNSScalar) { _ = "STUB: not implemented"; return }

// MulRNSScalar multiplies s1 and s2 and stores the result in sout.
func (r Ring) MulRNSScalar(s1, s2, sout ring.RNSScalar) { _ = "STUB: not implemented"; return }

// EvalPolyScalar evaluate the polynomial pol at pt and writes the result in p3
func (r Ring) EvalPolyScalar(pol []Poly, pt uint64, p3 Poly) { _ = "STUB: not implemented"; return }

// MulScalar multiplies p1 by scalar and returns the result in p2.
func (r Ring) MulScalar(p1 Poly, scalar uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// NTT computes the NTT of p1 and returns the result on p2.
func (r Ring) NTT(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// INTT computes the inverse-NTT of p1 and returns the result on p2.
func (r Ring) INTT(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// NTTLazy computes the NTT of p1 and returns the result on p2.
// Output values are in the range [0, 2q-1].
func (r Ring) NTTLazy(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// INTTLazy computes the inverse-NTT of p1 and returns the result on p2.
// Output values are in the range [0, 2q-1].
func (r Ring) INTTLazy(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// MForm switches p1 to the Montgomery domain and writes the result on p2.
func (r Ring) MForm(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// IMForm switches back p1 from the Montgomery domain to the conventional domain and writes the result on p2.
func (r Ring) IMForm(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomery multiplies p1 by p2 coefficient-wise with a Montgomery modular reduction.
func (r Ring) MulCoeffsMontgomery(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazy multiplies p1 by p2 coefficient-wise with a constant-time Montgomery modular reduction.
// Result is within [0, 2q-1].
func (r Ring) MulCoeffsMontgomeryLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazyThenAddLazy multiplies p1 by p2 coefficient-wise with a
// constant-time Montgomery modular reduction and adds the result on p3.
// Result is within [0, 2q-1]
func (r Ring) MulCoeffsMontgomeryLazyThenAddLazy(p1, p2, p3 Poly) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryThenSub multiplies p1 by p2 coefficient-wise with
// a Montgomery modular reduction and subtracts the result from p3.
func (r Ring) MulCoeffsMontgomeryThenSub(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazyThenSubLazy multiplies p1 by p2 coefficient-wise with
// a Montgomery modular reduction and subtracts the result from p3.
func (r Ring) MulCoeffsMontgomeryLazyThenSubLazy(p1, p2, p3 Poly) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryThenAdd multiplies p1 by p2 coefficient-wise with a
// Montgomery modular reduction and adds the result to p3.
func (r Ring) MulCoeffsMontgomeryThenAdd(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulRNSScalarMontgomery multiplies p with a scalar value expressed in the CRT decomposition.
// It assumes the scalar decomposition to be in Montgomery form.
func (r Ring) MulRNSScalarMontgomery(p Poly, scalar []uint64, pOut Poly) {
	_ = "STUB: not implemented"
	return
}

// Inverse computes the modular inverse of a scalar a expressed in a CRT decomposition.
// The inversion is done in-place and assumes that a is in Montgomery form.
func (r Ring) Inverse(scalar ring.RNSScalar) { _ = "STUB: not implemented"; return }

// Reduce applies the modular reduction on the coefficients of p1 and returns the result on p2.
func (r Ring) Reduce(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// Automorphism applies the automorphism X^{i} -> X^{i*gen} on p1 and writes the result on p2.
// Method is not in place.
func (r Ring) Automorphism(p1 Poly, galEl uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// AutomorphismNTT applies the automorphism X^{i} -> X^{i*gen} on p1 and writes the result on p2.
// Method is not in place.
// Inputs are assumed to be in the NTT domain.
func (r Ring) AutomorphismNTT(p1 Poly, galEl uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// AutomorphismNTTWithIndex applies the automorphism X^{i} -> X^{i*gen} on p1 and writes the result on p2.
// Index of automorphism must be provided.
// Method is not in place.
func (r Ring) AutomorphismNTTWithIndex(p1 Poly, index []uint64, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// AutomorphismNTTWithIndexThenAddLazy applies the automorphism X^{i} -> X^{i*gen} on p1 and adds the result on p2.
// Index of automorphism must be provided.
// Method is not in place.
func (r Ring) AutomorphismNTTWithIndexThenAddLazy(p1 Poly, index []uint64, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// ExtendBasisSmallNormAndCenter extends a small-norm polynomial polQ in R_Q to a polynomial
// polQP in R_QP.
func (r Ring) ExtendBasisSmallNormAndCenter(polyInQ ring.Poly, levelP int, polyOutQ, polyOutP ring.Poly) {
	_ = "STUB: not implemented"
	return
}
