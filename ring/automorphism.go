package ring

// AutomorphismNTTIndex computes the look-up table for the automorphism X^{i} -> X^{i*k mod NthRoot}.
func AutomorphismNTTIndex(N int, NthRoot, GalEl uint64) (index []uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AutomorphismNTT applies the automorphism X^{i} -> X^{i*gen} on a polynomial in the NTT domain.
// It must be noted that the result cannot be in-place.
func (r Ring) AutomorphismNTT(polIn Poly, gen uint64, polOut Poly) {
	_ = "STUB: not implemented"
	return
}

// Sanity check, this error should not happen.

// AutomorphismNTTWithIndex applies the automorphism X^{i} -> X^{i*gen} on a polynomial in the NTT domain.
// `index` is the lookup table storing the mapping of the automorphism.
// It must be noted that the result cannot be in-place.
func (r Ring) AutomorphismNTTWithIndex(polIn Poly, index []uint64, polOut Poly) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(index)%8 != 0  */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(polOut.Coeffs)%8 != 0 */

// AutomorphismNTTWithIndexThenAddLazy applies the automorphism X^{i} -> X^{i*gen} on a polynomial in the NTT domain .
// `index` is the lookup table storing the mapping of the automorphism.
// The result of the automorphism is added on polOut.
func (r Ring) AutomorphismNTTWithIndexThenAddLazy(polIn Poly, index []uint64, polOut Poly) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(index)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(polOut.Coeffs)%8 != 0 */

// Automorphism applies the automorphism X^{i} -> X^{i*gen} on a polynomial outside of the NTT domain.
// It must be noted that the result cannot be in-place.
func (r Ring) Automorphism(polIn Poly, gen uint64, polOut Poly) { _ = "STUB: not implemented"; return }

/* #nosec G115 -- N cannot be negative */

/* #nosec G115 -- bitsize cannot be negative */

// TODO: find a more efficient way to do
// the automorphism on Z[X+X^-1]

// Only consider i -> index if within [0, N-1]

// If the starting index is within [N, 2N-1]

// Wrap back between [0, N-1]
// Negate

/* #nosec G115 -- bitsize cannot be negative */
