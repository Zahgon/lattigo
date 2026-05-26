package ring

func addvec(p1, p2, p3 []uint64, modulus uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func addlazyvec(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func subvec(p1, p2, p3 []uint64, modulus uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func sublazyvec(p1, p2, p3 []uint64, modulus uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func negvec(p1, p2 []uint64, modulus uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func reducevec(p1, p2 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func reducelazyvec(p1, p2 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func mulcoeffslazyvec(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffslazythenaddlazyvec(p1, p2, p3 []uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsbarrettvec(p1, p2, p3 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsbarrettlazyvec(p1, p2, p3 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsthenaddvec(p1, p2, p3 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsbarrettthenaddlazyvec(p1, p2, p3 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomeryvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerylazyvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerythenaddvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerythenaddlazyvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerylazythenaddlazyvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerythensubvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerythensublazyvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerylazythensublazyvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mulcoeffsmontgomerylazythenNegvec(p1, p2, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func addlazythenmulscalarmontgomeryvec(p1, p2 []uint64, scalarMont uint64, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func addscalarlazythenmulscalarmontgomeryvec(p1 []uint64, scalar0, scalarMont1 uint64, p2 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func addscalarvec(p1 []uint64, scalar uint64, p2 []uint64, modulus uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func addscalarlazyvec(p1 []uint64, scalar uint64, p2 []uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func addscalarlazythenNegTwoModuluslazyvec(p1 []uint64, scalar uint64, p2 []uint64, modulus uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func subscalarvec(p1 []uint64, scalar uint64, p2 []uint64, modulus uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func mulscalarmontgomeryvec(p1 []uint64, scalarMont uint64, p2 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func mulscalarmontgomerylazyvec(p1 []uint64, scalarMont uint64, p2 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func mulscalarmontgomerythenaddvec(p1 []uint64, scalarMont uint64, p2 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func mulscalarmontgomerythenaddscalarvec(p1 []uint64, scalar0, scalarMont1 uint64, p2 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func subthenmulscalarmontgomeryTwoModulusvec(p1, p2 []uint64, scalarMont uint64, p3 []uint64, modulus, mredconstant uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 */

func mformvec(p1, p2 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func mformlazyvec(p1, p2 []uint64, modulus uint64, bredconstant [2]uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

func imformvec(p1, p2 []uint64, modulus, mredconstant uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

// ZeroVec sets all values of p1 to zero.
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func ZeroVec(p1 []uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

// MaskVec evaluates p2 = vec(p1>>w) & mask
// Iteration is done with respect to len(p1).
// All input must have a size which is a multiple of 8.
func MaskVec(p1 []uint64, w int, mask uint64, p2 []uint64) { _ = "STUB: not implemented"; return }

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */
