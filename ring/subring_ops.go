package ring

// Add evaluates p3 = p1 + p2 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) Add(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// AddLazy evaluates p3 = p1 + p2.
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) AddLazy(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// Sub evaluates p3 = p1 - p2 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) Sub(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// SubLazy evaluates p3 = p1 - p2.
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) SubLazy(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// Neg evaluates p2 = -p1 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) Neg(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// Reduce evaluates p2 = p1 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) Reduce(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// ReduceLazy evaluates p2 = p1 (mod modulus) with p2 in range [0, 2*modulus-1].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) ReduceLazy(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsLazy evaluates p3 = p1*p2.
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsLazy(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsLazyThenAddLazy evaluates p3 = p3 + p1*p2.
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsLazyThenAddLazy(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrett evaluates p3 = p1*p2 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsBarrett(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrettLazy evaluates p3 = p1*p2 (mod modulus) with p3 in [0, 2*modulus-1].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsBarrettLazy(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrettThenAdd evaluates p3 = p3 + (p1*p2) (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsBarrettThenAdd(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsBarrettThenAddLazy evaluates p3 = p3 + p1*p2 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsBarrettThenAddLazy(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomery evaluates p3 = p1*p2 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomery(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryLazy evaluates p3 = p1*p2 (mod modulus) with p3 in range [0, 2*modulus-1].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryLazy(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

// MulCoeffsMontgomeryThenAdd evaluates p3 = p3 + (p1*p2) (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryThenAdd(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryThenAddLazy evaluates p3 = p3 + (p1*p2 (mod modulus)).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryThenAddLazy(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryLazyThenAddLazy evaluates p3 = p3 + p1*p2 (mod modulus) with p3 in range [0, 3modulus-2].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryLazyThenAddLazy(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryThenSub evaluates p3 = p3 - p1*p2 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryThenSub(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryThenSubLazy evaluates p3 = p3 - p1*p2 (mod modulus) with p3 in range [0, 2*modulus-2].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryThenSubLazy(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryLazyThenSubLazy evaluates p3 = p3 - p1*p2 (mod modulus) with p3 in range [0, 3*modulus-2].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryLazyThenSubLazy(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulCoeffsMontgomeryLazyThenNeg evaluates p3 = - p1*p2 (mod modulus) with p3 in range [0, 2*modulus-2].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulCoeffsMontgomeryLazyThenNeg(p1, p2, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// AddLazyThenMulScalarMontgomery evaluates p3 = (p1+p2)*scalarMont (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) AddLazyThenMulScalarMontgomery(p1, p2 []uint64, scalarMont uint64, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// AddScalarLazyThenMulScalarMontgomery evaluates p3 = (scalarMont0+p2)*scalarMont1 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) AddScalarLazyThenMulScalarMontgomery(p1 []uint64, scalar0, scalarMont1 uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// AddScalar evaluates p2 = p1 + scalar (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) AddScalar(p1 []uint64, scalar uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// AddScalarLazy evaluates p2 = p1 + scalar.
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) AddScalarLazy(p1 []uint64, scalar uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// AddScalarLazyThenNegTwoModulusLazy evaluates p2 = 2*modulus - p1 + scalar.
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) AddScalarLazyThenNegTwoModulusLazy(p1 []uint64, scalar uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// SubScalar evaluates p2 = p1 - scalar (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) SubScalar(p1 []uint64, scalar uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulScalarMontgomery evaluates p2 = p1*scalarMont (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulScalarMontgomery(p1 []uint64, scalarMont uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulScalarMontgomeryLazy evaluates p2 = p1*scalarMont (mod modulus) with p2 in range [0, 2*modulus-1].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulScalarMontgomeryLazy(p1 []uint64, scalarMont uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulScalarMontgomeryThenAdd evaluates p2 = p2 + p1*scalarMont (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulScalarMontgomeryThenAdd(p1 []uint64, scalarMont uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// MulScalarMontgomeryThenAddScalar evaluates p2 = scalar + p1*scalarMont (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MulScalarMontgomeryThenAddScalar(p1 []uint64, scalar0, scalarMont1 uint64, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// SubThenMulScalarMontgomeryTwoModulus evaluates p3 = (p1 + twomodulus - p2) * scalarMont (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) SubThenMulScalarMontgomeryTwoModulus(p1, p2 []uint64, scalarMont uint64, p3 []uint64) {
	_ = "STUB: not implemented"
	return
}

// NTT evaluates p2 = NTT(p1).
func (s *SubRing) NTT(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// NTTLazy evaluates p2 = NTT(p1) with p2 in [0, 6*modulus-2].
func (s *SubRing) NTTLazy(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// INTT evaluates p2 = INTT(p1).
func (s *SubRing) INTT(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// INTTLazy evaluates p2 = INTT(p1) with p2 in [0, 2*modulus-1].
func (s *SubRing) INTTLazy(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// MForm evaluates p2 = p1 * 2^64 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MForm(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// MFormLazy evaluates p2 = p1 * 2^64 (mod modulus) with p2 in the range [0, 2*modulus-1].
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) MFormLazy(p1, p2 []uint64) { _ = "STUB: not implemented"; return }

// IMForm evaluates p2 = p1 * (2^64)^-1 (mod modulus).
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func (s *SubRing) IMForm(p1, p2 []uint64) { _ = "STUB: not implemented"; return }
