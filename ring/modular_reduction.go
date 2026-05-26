package ring

// MForm switches a to the Montgomery domain by computing
// a*2^64 mod q.
func MForm(a, q uint64, bredconstant [2]uint64) (r uint64) {
	_ = "STUB: not implemented"
	// R = 2^128/q
	// b = a * 2^{64}
	// mhi = (b * (R % 2^{64})) / 2^{64}
	//
	// r = b - ((b * (R/2^{64}) + mhi) * q
	//
	//	= b - (b * R)/2^{64} * q
	//	= b - n * q
	//
	// Since r is in [0, 2q] and q is in [0, 2^{63}-1]
	// we are ensured that r < 2^64, thus r = r mod 2^{64}
	// So we can work mod 2^{64} (i.e. with uint64)
	// and therefore:
	//
	// r = b - n * q mod 2^{64}
	//
	//	= a * 2^{64} - n * q mod 2^{64}
	//	= - n * q mod 2^{64}
	return 0
}

// MFormLazy switches a to the Montgomery domain by computing
// a*2^64 mod q in constant time.
// The result is between 0 and 2*q-1.
func MFormLazy(a, q uint64, bredconstant [2]uint64) (r uint64) {
	_ = "STUB: not implemented"
	// See MForm for the implementation trick.
	return 0
}

// IMForm switches a from the Montgomery domain back to the
// standard domain by computing a*(1/2^64) mod q.
func IMForm(a, q, mredconstant uint64) (r uint64) { _ = "STUB: not implemented"; return 0 }

// IMFormLazy switches a from the Montgomery domain back to the
// standard domain by computing a*(1/2^64) mod q in constant time.
// The result is between 0 and 2*q-1.
func IMFormLazy(a, q, mredconstant uint64) (r uint64) { _ = "STUB: not implemented"; return 0 }

// GenMRedConstant computes the constant mredconstant = (q^-1) mod 2^64 required for MRed.
func GenMRedConstant(q uint64) (mredconstant uint64) { _ = "STUB: not implemented"; return 0 }

// MRed computes x * y * (1/2^64) mod q.
func MRed(x, y, q, mredconstant uint64) (r uint64) { _ = "STUB: not implemented"; return 0 }

// MRedLazy computes x * y * (1/2^64) mod q in constant time.
// The result is between 0 and 2*q-1.
func MRedLazy(x, y, q, mredconstant uint64) (r uint64) { _ = "STUB: not implemented"; return 0 }

// GenBRedConstant computes the constant for the BRed algorithm.
// Returns ((2^128)/q)/(2^64) and (2^128)/q mod 2^64.
func GenBRedConstant(q uint64) [2]uint64 { _ = "STUB: not implemented"; return nil }

// BRedAdd computes a mod q.
func BRedAdd(a, q uint64, bredconstant [2]uint64) (r uint64) { _ = "STUB: not implemented"; return 0 }

// BRedAddLazy computes a mod q in constant time.
// The result is between 0 and 2*q-1.
func BRedAddLazy(x, q uint64, bredconstant [2]uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// BRed computes x*y mod q.
func BRed(x, y, q uint64, bredconstant [2]uint64) (r uint64) { _ = "STUB: not implemented"; return 0 }

// computes r = mhi * uhi + (mlo * uhi + mhi * ulo)<<64 + (mlo * ulo)) >> 128

// r = mhi * uhi

// mlo * uhi

// mlo * ulo

// mhi * ulo

// BRedLazy computes x*y mod q in constant time.
// The result is between 0 and 2*q-1.
func BRedLazy(x, y, q uint64, bredconstant [2]uint64) (r uint64) {
	_ = "STUB: not implemented"
	return 0
}

// computes r = mhi * uhi + (mlo * uhi + mhi * ulo)<<64 + (mlo * ulo)) >> 128

// r = mhi * uhi

// mlo * uhi

// mlo * ulo

// mhi * ulo

// CRed reduce returns a mod q where a is between 0 and 2*q-1.
func CRed(a, q uint64) uint64 { _ = "STUB: not implemented"; return 0 }
