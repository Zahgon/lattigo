package ring

// UnfoldConjugateInvariantToStandard maps the compressed representation (N/2 coefficients)
// of Z_Q[X+X^-1]/(X^2N + 1) to full representation in Z_Q[X]/(X^2N+1).
// Requires degree(polyConjugateInvariant) = 2*degree(polyStandard).
// Requires that polyStandard and polyConjugateInvariant share the same moduli.
func (r Ring) UnfoldConjugateInvariantToStandard(polyConjugateInvariant, polyStandard Poly) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// FoldStandardToConjugateInvariant folds [X]/(X^N+1) to [X+X^-1]/(X^N+1) in compressed form (N/2 coefficients).
// Requires degree(polyConjugateInvariant) = 2*degree(polyStandard).
// Requires that polyStandard and polyConjugateInvariant share the same moduli.
func (r Ring) FoldStandardToConjugateInvariant(polyStandard Poly, permuteNTTIndexInv []uint64, polyConjugateInvariant Poly) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// PadDefaultRingToConjugateInvariant converts a polynomial in Z[X]/(X^N +1) to a polynomial in Z[X+X^-1]/(X^2N+1).
func (r Ring) PadDefaultRingToConjugateInvariant(polyStandard Poly, IsNTT bool, polyConjugateInvariant Poly) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}
