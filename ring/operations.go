package ring

import (
	"math/big"
)

// Add evaluates p3 = p1 + p2 coefficient-wise in the ring.
func (r Ring) Add(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// AddLazy evaluates p3 = p1 + p2 coefficient-wise in the ring, with p3 in [0, 2*modulus-1].
func (r Ring) AddLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// Sub evaluates p3 = p1 - p2 coefficient-wise in the ring.
func (r Ring) Sub(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// SubLazy evaluates p3 = p1 - p2 coefficient-wise in the ring, with p3 in [0, 2*modulus-1].
func (r Ring) SubLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// Neg evaluates p2 = -p1 coefficient-wise in the ring.
func (r Ring) Neg(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// Reduce evaluates p2 = p1 coefficient-wise mod modulus in the ring.
func (r Ring) Reduce(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// ReduceLazy evaluates p2 = p1 coefficient-wise mod modulus in the ring, with p2 in [0, 2*modulus-1].
func (r Ring) ReduceLazy(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrett evaluates p3 = p1 * p2 coefficient-wise in the ring, with Barrett reduction.
func (r Ring) MulCoeffsBarrett(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrettLazy evaluates p3 = p1 * p2 coefficient-wise in the ring, with Barrett reduction, with p3 in [0, 2*modulus-1].
func (r Ring) MulCoeffsBarrettLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrettThenAdd evaluates p3 = p3 + p1 * p2 coefficient-wise in the ring, with Barrett reduction.
func (r Ring) MulCoeffsBarrettThenAdd(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrettThenAddLazy evaluates p3 = p1 * p2 coefficient-wise in the ring, with Barrett reduction, with p3 in [0, 2*modulus-1].
func (r Ring) MulCoeffsBarrettThenAddLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomery evaluates p3 = p1 * p2 coefficient-wise in the ring, with Montgomery reduction.
func (r Ring) MulCoeffsMontgomery(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazy evaluates p3 = p1 * p2 coefficient-wise in the ring, with Montgomery reduction, with p3 in [0, 2*modulus-1].
func (r Ring) MulCoeffsMontgomeryLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazyThenNeg evaluates p3 = -p1 * p2 coefficient-wise in the ring, with Montgomery reduction, with p3 in [0, 2*modulus-1].
func (r Ring) MulCoeffsMontgomeryLazyThenNeg(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryThenAdd evaluates p3 = p3 + p1 * p2 coefficient-wise in the ring, with Montgomery reduction, with p3 in [0, 2*modulus-1].
func (r Ring) MulCoeffsMontgomeryThenAdd(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryThenAddLazy evaluates p3 = p3 + p1 * p2 coefficient-wise in the ring, with Montgomery reduction, with p3 in [0, 2*modulus-1].
func (r Ring) MulCoeffsMontgomeryThenAddLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazyThenAddLazy evaluates p3 = p3 + p1 * p2 coefficient-wise in the ring, with Montgomery reduction, with p3 in [0, 3*modulus-2].
func (r Ring) MulCoeffsMontgomeryLazyThenAddLazy(p1, p2, p3 Poly) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryThenSub evaluates p3 = p3 - p1 * p2 coefficient-wise in the ring, with Montgomery reduction.
func (r Ring) MulCoeffsMontgomeryThenSub(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryThenSubLazy evaluates p3 = p3 - p1 * p2 coefficient-wise in the ring, with Montgomery reduction, with p3 in [0, 2*modulus-1].
func (r Ring) MulCoeffsMontgomeryThenSubLazy(p1, p2, p3 Poly) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazyThenSubLazy evaluates p3 = p3 - p1 * p2 coefficient-wise in the ring, with Montgomery reduction, with p3 in [0, 3*modulus-2].
func (r Ring) MulCoeffsMontgomeryLazyThenSubLazy(p1, p2, p3 Poly) {
	_ = "STUB: not implemented"
	return
}

// AddScalar evaluates p2 = p1 + scalar coefficient-wise in the ring.
func (r Ring) AddScalar(p1 Poly, scalar uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// AddScalarBigint evaluates p2 = p1 + scalar coefficient-wise in the ring.
func (r Ring) AddScalarBigint(p1 Poly, scalar *big.Int, p2 Poly) { _ = "STUB: not implemented"; return }

// AddDoubleRNSScalar evaluates p2 = p1[:N/2] + scalar0 || p1[N/2] + scalar1 coefficient-wise in the ring,
// with the scalar values expressed in the CRT decomposition at a given level.
func (r Ring) AddDoubleRNSScalar(p1 Poly, scalar0, scalar1 RNSScalar, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// SubDoubleRNSScalar evaluates p2 = p1[:N/2] - scalar0 || p1[N/2] - scalar1 coefficient-wise in the ring,
// with the scalar values expressed in the CRT decomposition at a given level.
func (r Ring) SubDoubleRNSScalar(p1 Poly, scalar0, scalar1 RNSScalar, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// SubScalar evaluates p2 = p1 - scalar coefficient-wise in the ring.
func (r Ring) SubScalar(p1 Poly, scalar uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// SubScalarBigint evaluates p2 = p1 - scalar coefficient-wise in the ring.
func (r Ring) SubScalarBigint(p1 Poly, scalar *big.Int, p2 Poly) { _ = "STUB: not implemented"; return }

// MulScalar evaluates p2 = p1 * scalar coefficient-wise in the ring.
func (r Ring) MulScalar(p1 Poly, scalar uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// MulScalarThenAdd evaluates p2 = p2 + p1 * scalar coefficient-wise in the ring.
func (r Ring) MulScalarThenAdd(p1 Poly, scalar uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// MulRNSScalarMontgomery evaluates p2 = p1 * scalar coefficient-wise in the ring, with a scalar value expressed in the CRT decomposition at a given level.
// It assumes the scalar decomposition to be in Montgomery form.
func (r Ring) MulRNSScalarMontgomery(p1 Poly, scalar RNSScalar, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// MulScalarThenSub evaluates p2 = p2 - p1 * scalar coefficient-wise in the ring.
func (r Ring) MulScalarThenSub(p1 Poly, scalar uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// MulScalarBigint evaluates p2 = p1 * scalar coefficient-wise in the ring.
func (r Ring) MulScalarBigint(p1 Poly, scalar *big.Int, p2 Poly) { _ = "STUB: not implemented"; return }

// MulScalarBigintThenAdd evaluates p2 = p1 * scalar coefficient-wise in the ring.
func (r Ring) MulScalarBigintThenAdd(p1 Poly, scalar *big.Int, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// MulDoubleRNSScalar evaluates p2 = p1[:N/2] * scalar0 || p1[N/2] * scalar1 coefficient-wise in the ring,
// with the scalar values expressed in the CRT decomposition at a given level.
func (r Ring) MulDoubleRNSScalar(p1 Poly, scalar0, scalar1 RNSScalar, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// MulDoubleRNSScalarThenAdd evaluates p2 = p2 + p1[:N/2] * scalar0 || p1[N/2] * scalar1 coefficient-wise in the ring,
// with the scalar values expressed in the CRT decomposition at a given level.
func (r Ring) MulDoubleRNSScalarThenAdd(p1 Poly, scalar0, scalar1 RNSScalar, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// EvalPolyScalar evaluate p2 = p1(scalar) coefficient-wise in the ring.
func (r Ring) EvalPolyScalar(p1 []Poly, scalar uint64, p2 Poly) { _ = "STUB: not implemented"; return }

// Shift evaluates p2 = p2<<<k coefficient-wise in the ring.
func (r Ring) Shift(p1 Poly, k int, p2 Poly) { _ = "STUB: not implemented"; return }

// MForm evaluates p2 = p1 * (2^64)^-1 coefficient-wise in the ring.
func (r Ring) MForm(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// MFormLazy evaluates p2 = p1 * (2^64)^-1 coefficient-wise in the ring with p2 in [0, 2*modulus-1].
func (r Ring) MFormLazy(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// IMForm evaluates p2 = p1 * 2^64 coefficient-wise in the ring.
func (r Ring) IMForm(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// MultByMonomial evaluates p2 = p1 * X^k coefficient-wise in the ring.
func (r Ring) MultByMonomial(p1 Poly, k int, p2 Poly) { _ = "STUB: not implemented"; return }

// MulByVectorMontgomery evaluates p2 = p1 * vector coefficient-wise in the ring.
func (r Ring) MulByVectorMontgomery(p1 Poly, vector []uint64, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// MulByVectorMontgomeryThenAddLazy evaluates p2 = p2 + p1 * vector coefficient-wise in the ring.
func (r Ring) MulByVectorMontgomeryThenAddLazy(p1 Poly, vector []uint64, p2 Poly) {
	_ = "STUB: not implemented"
	return
}

// MapSmallDimensionToLargerDimensionNTT maps Y = X^{N/n} -> X directly in the NTT domain
func MapSmallDimensionToLargerDimensionNTT(polSmall, polLarge Poly) {
	_ = "STUB: not implemented"
	return
}
