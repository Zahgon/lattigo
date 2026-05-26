package ring

// DivFloorByLastModulusNTT divides (floored) the polynomial by its last modulus.
// The input must be in the NTT domain.
// Output poly level must be equal or one less than input level.
func (r Ring) DivFloorByLastModulusNTT(p0, p1 Poly) { _ = "STUB: not implemented"; return }

// (-x[i] + x[-1]) * -InvQ

// DivFloorByLastModulus divides (floored) the polynomial by its last modulus.
// Output poly level must be equal or one less than input level.
func (r Ring) DivFloorByLastModulus(p0, p1 Poly) { _ = "STUB: not implemented"; return }

// DivFloorByLastModulusManyNTT divides (floored) sequentially nbRescales times the polynomial by its last modulus. Input must be in the NTT domain.
// Output poly level must be equal or nbRescales less than input level.
func (r Ring) DivFloorByLastModulusManyNTT(nbRescales int, p0, p1 Poly) {
	_ = "STUB: not implemented"
	return
}

// DivFloorByLastModulusMany divides (floored) sequentially nbRescales times the polynomial by its last modulus.
// Output poly level must be equal or nbRescales less than input level.
func (r Ring) DivFloorByLastModulusMany(nbRescales int, p0, buff, p1 Poly) {
	_ = "STUB: not implemented"
	return
}

// DivRoundByLastModulusNTT divides (rounded) the polynomial by its last modulus. The input must be in the NTT domain.
// Output poly level must be equal or one less than input level.
func (r Ring) DivRoundByLastModulusNTT(p0, p1 Poly) { _ = "STUB: not implemented"; return }

// Center by (p-1)/2

// DivRoundByLastModulus divides (rounded) the polynomial by its last modulus. The input must be in the NTT domain.
// Output poly level must be equal or one less than input level.
func (r Ring) DivRoundByLastModulus(p0, p1 Poly) { _ = "STUB: not implemented"; return }

// Center by (p-1)/2

// DivRoundByLastModulusManyNTT divides (rounded) sequentially nbRescales times the polynomial by its last modulus. The input must be in the NTT domain.
// Output poly level must be equal or nbRescales less than input level.
func (r Ring) DivRoundByLastModulusManyNTT(nbRescales int, p0, buff, p1 Poly) {
	_ = "STUB: not implemented"
	return
}

// DivRoundByLastModulusMany divides (rounded) sequentially nbRescales times the polynomial by its last modulus.
// Output poly level must be equal or nbRescales less than input level.
func (r Ring) DivRoundByLastModulusMany(nbRescales int, p0, buff, p1 Poly) {
	_ = "STUB: not implemented"
	return
}
