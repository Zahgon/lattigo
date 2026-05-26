package ring

// BasisExtender stores the necessary parameters for RNS basis extension.
// The used algorithm is from https://eprint.iacr.org/2018/117.pdf.
type BasisExtender struct {
	ringQ                *Ring
	ringP                *Ring
	constantsQtoP        []ModUpConstants
	constantsPtoQ        []ModUpConstants
	modDownConstantsPtoQ [][]uint64
	modDownConstantsQtoP [][]uint64
	poolQ                *BufferPool
	poolP                *BufferPool
}

func genmodDownConstants(ringQ, ringP *Ring) (constants [][]uint64) {
	_ = "STUB: not implemented"
	return nil
}

// NewBasisExtender creates a new BasisExtender, enabling RNS basis extension from Q to P and P to Q.
func NewBasisExtender(ringQ, ringP *Ring) (be *BasisExtender) {
	_ = "STUB: not implemented"
	return nil
}

// we use the same backing pool only if ringQ and ringP have the same dimension

// ModUpConstants stores the necessary parameters for RNS basis extension.
type ModUpConstants struct {
	// Parameters for basis extension from Q to P
	// (Q/Qi)^-1) (mod each Qi) (in Montgomery form)
	qoverqiinvqi []uint64
	// Q/qi (mod each Pj) (in Montgomery form)
	qoverqimodp [][]uint64
	// Q*v (mod each Pj) for v in [1,...,k] where k is the number of Pj moduli
	vtimesqmodp [][]uint64
}

// GenModUpConstants generates the ModUpConstants for basis extension from Q to P and P to Q.
func GenModUpConstants(Q, P []uint64) ModUpConstants {
	_ = "STUB: not implemented"
	return *new(ModUpConstants)
}

// (Q/Qi)^-1) * r (mod Qi) (in Montgomery form)
/* #nosec G115 -- library requires 64-bit system -> int = int64 */

// (Q/qi * r) (mod Pj) (in Montgomery form)

// Correction Term (v*Q) mod each Pj

// ModUpQtoP extends the RNS basis of a polynomial from Q to QP.
// Given a polynomial with coefficients in basis {Q0,Q1....Qlevel},
// it extends its basis from {Q0,Q1....Qlevel} to {Q0,Q1....Qlevel,P0,P1...Pj}
func (be *BasisExtender) ModUpQtoP(levelQ, levelP int, polQ, polP Poly) {
	_ = "STUB: not implemented"
	return
}

// ModUpPtoQ extends the RNS basis of a polynomial from P to PQ.
// Given a polynomial with coefficients in basis {P0,P1....Plevel},
// it extends its basis from {P0,P1....Plevel} to {Q0,Q1...Qj}
func (be *BasisExtender) ModUpPtoQ(levelP, levelQ int, polP, polQ Poly) {
	_ = "STUB: not implemented"
	return
}

// ModDownQPtoQ reduces the basis of a polynomial.
// Given a polynomial with coefficients in basis {Q0,Q1....Qlevel} and {P0,P1...Pj},
// it reduces its basis from {Q0,Q1....Qlevel} and {P0,P1...Pj} to {Q0,Q1....Qlevel}
// and does a rounded integer division of the result by P.
func (be *BasisExtender) ModDownQPtoQ(levelQ, levelP int, p1Q, p1P, p2Q Poly) {
	_ = "STUB: not implemented"
	return
}

// ModDownQPtoQNTT reduces the basis of a polynomial.
// Given a polynomial with coefficients in basis {Q0,Q1....Qi} and {P0,P1...Pj},
// it reduces its basis from {Q0,Q1....Qi} and {P0,P1...Pj} to {Q0,Q1....Qi}
// and does a rounded integer division of the result by P.
// Inputs must be in the NTT domain.
func (be *BasisExtender) ModDownQPtoQNTT(levelQ, levelP int, p1Q, p1P, p2Q Poly) {
	_ = "STUB: not implemented"
	return
}

// Finally, for each level of p1 (and the buffer since they now share the same basis) we compute p2 = (P^-1) * (p1 - buff) mod Q

// Then for each coefficient we compute (P^-1) * (p1[i][j] - buff[i][j]) mod qi

// ModDownQPtoP reduces the basis of a polynomial.
// Given a polynomial with coefficients in basis {Q0,Q1....QlevelQ} and {P0,P1...PlevelP},
// it reduces its basis from {Q0,Q1....QlevelQ} and {P0,P1...PlevelP} to {P0,P1...PlevelP}
// and does a floored integer division of the result by Q.
func (be *BasisExtender) ModDownQPtoP(levelQ, levelP int, p1Q, p1P, p2P Poly) {
	_ = "STUB: not implemented"
	return
}

// Finally, for each level of p1 (and buff since they now share the same basis) we compute p2 = (P^-1) * (p1 - buff) mod Q

// Then, for each coefficient we compute (P^-1) * (p1[i][j] - buff[i][j]) mod qi

// In total we do len(P) + len(Q) NTT, which is optimal (linear in the number of moduli of P and Q)

// ModUpExact takes p1 mod Q and switches its basis to P, returning the result on p2.
// Caution: values are not centered and returned values are in [0, 2P-1].
func ModUpExact(p1, p2 [][]uint64, ringQ, ringP *Ring, MUC ModUpConstants) {
	_ = "STUB: not implemented"
	return
}

// We loop over each coefficient and apply the basis extension

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2[j])%8 != 0*/

// Decomposer is a structure that stores the parameters of the arbitrary decomposer.
// This decomposer takes a p(x)_Q (in basis Q) and returns p(x) mod qi in basis QP, where
// qi = prod(Q_i) for 0<=i<=L, where L is the number of factors in P.
type Decomposer struct {
	ringQ, ringP   *Ring
	ModUpConstants [][][]ModUpConstants
}

// NewDecomposer creates a new Decomposer.
func NewDecomposer(ringQ, ringP *Ring) (decomposer *Decomposer) {
	_ = "STUB: not implemented"
	return nil
}

// Create ModUpConstants for each possible combination of [Qi,Pj] according to xnbPi

// DecomposeAndSplit decomposes a polynomial p(x) in basis Q, reduces it modulo qi, and returns
// the result in basis QP separately.
func (decomposer *Decomposer) DecomposeAndSplit(levelQ, levelP, nbPi, BaseRNSDecompositionVectorSize int, p0Q, p1Q, p1P Poly) {
	_ = "STUB: not implemented"
	return
}

// First we check if the vector can simply by coping and rearranging elements (the case where no reconstruction is needed)

// Otherwise, we apply a fast exact base conversion for the reconstruction

// We loop over each coefficient and apply the basis extension

// Coefficients of index smaller than the ones to be decomposed

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1Q.Coeffs[j])%8 != 0 */

// Coefficients of index greater than the ones to be decomposed

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1Q.Coeffs[j])%8 != 0 */

// Coefficients of the special primes Pi

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1P.Coeffs[j])%8 != 0 */

func reconstructRNSCentered(start, end, x int, p [][]uint64, v *[8]uint64, vi *[8]float64, y0, y1, y2, y3, y4, y5, y6, y7 *[32]uint64, QHalfModqi, Q, mredQ, qoverqiinvqi []uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p[j])%8 != 0 */

// Computation of the correction term v * Q%pi

// Index of the correction term

func reconstructRNS(start, end, x int, p [][]uint64, v *[8]uint64, y0, y1, y2, y3, y4, y5, y6, y7 *[32]uint64, Q, QInv, QbMont []uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p[i])%8 != 0 */

// Computation of the correction term v * Q%pi

// Caution, returns the values in [0, 2q-1]
func multSum(level int, res, rlo, rhi, v *[8]uint64, y0, y1, y2, y3, y4, y5, y6, y7 *[32]uint64, q, qInv uint64, vtimesqmodp, qoverqimodp []uint64) {
	_ = "STUB: not implemented"
	return
}

// Accumulates the sum on uint128 and does a lazy montgomery reduction at the end
