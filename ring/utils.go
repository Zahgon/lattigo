package ring

type Dimensions struct {
	Rows, Cols int
}

// EvalPolyModP evaluates y = sum poly[i] * x^{i} mod p.
func EvalPolyModP(x uint64, poly []uint64, p uint64) (y uint64) {
	_ = "STUB: not implemented"
	return 0
}

// Min returns the minimum between to int
func Min(x, y int) int { _ = "STUB: not implemented"; return 0 }

// ModExp performs the modular exponentiation x^e mod p,
// x and p are required to be at most 64 bits to avoid an overflow.
func ModExp(x, e, p uint64) (result uint64) { _ = "STUB: not implemented"; return 0 }

// ModExpPow2 performs the modular exponentiation x^e mod p, where p is a power of two,
// x and p are required to be at most 64 bits to avoid an overflow.
func ModExpPow2(x, e, p uint64) (result uint64) { _ = "STUB: not implemented"; return 0 }

// ModexpMontgomery performs the modular exponentiation x^e mod p,
// where x is in Montgomery form, and returns x^e in Montgomery form.
func ModexpMontgomery(x uint64, e int, q, mredconstant uint64, bredconstant [2]uint64) (result uint64) {
	_ = "STUB: not implemented"
	return 0
}
