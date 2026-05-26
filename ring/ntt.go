package ring

const (
	// MinimumRingDegreeForLoopUnrolledNTT is the minimum ring degree
	// necessary for memory safe loop unrolling
	MinimumRingDegreeForLoopUnrolledNTT = 16
)

// NumberTheoreticTransformer is an interface to provide
// flexibility on what type of NTT is used by the struct Ring.
type NumberTheoreticTransformer interface {
	Forward(p1, p2 []uint64)
	ForwardLazy(p1, p2 []uint64)
	Backward(p1, p2 []uint64)
	BackwardLazy(p1, p2 []uint64)
}

type numberTheoreticTransformerBase struct {
	*NTTTable
	N            int
	Modulus      uint64
	MRedConstant uint64
	BRedConstant [2]uint64
}

// NumberTheoreticTransformerStandard computes the standard nega-cyclic NTT in the ring Z[X]/(X^N+1).
type NumberTheoreticTransformerStandard struct {
	numberTheoreticTransformerBase
}

// NTTTable store all the constants that are specifically tied to the NTT.
type NTTTable struct {
	NthRoot       uint64   // Nthroot used for the NTT
	PrimitiveRoot uint64   // 2N-th primitive root
	RootsForward  []uint64 //powers of the 2N-th primitive root in Montgomery form (in bit-reversed order)
	RootsBackward []uint64 //powers of the inverse of the 2N-th primitive root in Montgomery form (in bit-reversed order)
	NInv          uint64   //[N^-1] mod Modulus in Montgomery form
}

func NewNumberTheoreticTransformerStandard(r *SubRing, n int) NumberTheoreticTransformer {
	_ = "STUB: not implemented"
	return *new(NumberTheoreticTransformer)
}

// Forward writes the forward NTT in Z[X]/(X^N+1) of p1 on p2.
func (rntt NumberTheoreticTransformerStandard) Forward(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// ForwardLazy writes the forward NTT in Z[X]/(X^N+1) of p1 on p2.
// Returns values in the range [0, 6q-2].
func (rntt NumberTheoreticTransformerStandard) ForwardLazy(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// Backward writes the backward NTT in Z[X]/(X^N+1) of p1 on p2.
func (rntt NumberTheoreticTransformerStandard) Backward(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// BackwardLazy writes the backward NTT in Z[X]/(X^N+1) p1 on p2.
// Returns values in the range [0, 2q-1].
func (rntt NumberTheoreticTransformerStandard) BackwardLazy(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// NumberTheoreticTransformerConjugateInvariant computes the NTT in the ring Z[X+X^-1]/(X^2N+1).
// Z[X+X^-1]/(X^2N+1) is a closed sub-ring of Z[X]/(X^2N+1). Note that the input polynomial only needs to be size N
// since the right half does not provide any additional information.
// See "Approximate Homomorphic Encryption over the Conjugate-invariant Ring", https://eprint.iacr.org/2018/952.
// The implemented approach is more efficient than the one proposed in the referenced work.
// It avoids the linear map Z[X + X^-1]/(X^2N + 1) <-> Z[X]/(X^N - 1) by instead directly computing the left
// half of the NTT of Z[X + X^-1]/(X^2N + 1) since the right half provides no additional information, which
// allows to (re)use nega-cyclic NTT.
type NumberTheoreticTransformerConjugateInvariant struct {
	numberTheoreticTransformerBase
}

func NewNumberTheoreticTransformerConjugateInvariant(r *SubRing, n int) NumberTheoreticTransformer {
	_ = "STUB: not implemented"
	return *new(NumberTheoreticTransformer)
}

// Forward writes the forward NTT in Z[X+X^-1]/(X^2N+1) of p1 on p2.
func (rntt NumberTheoreticTransformerConjugateInvariant) Forward(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// ForwardLazy writes the forward NTT in Z[X+X^-1]/(X^2N+1) of p1 on p2.
// Returns values in the range [0, 2q-1].
func (rntt NumberTheoreticTransformerConjugateInvariant) ForwardLazy(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// Backward writes the backward NTT in Z[X+X^-1]/(X^2N+1) of p1 on p2.
func (rntt NumberTheoreticTransformerConjugateInvariant) Backward(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// BackwardLazy writes the backward NTT in Z[X+X^-1]/(X^2N+1) of p1 on p2.
// Returns values in the range [0, 2q-1].
func (rntt NumberTheoreticTransformerConjugateInvariant) BackwardLazy(p1, p2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// NTT evaluates p2 = NTT(P1).
func (r Ring) NTT(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// NTTLazy evaluates p2 = NTT(p1) with p2 in [0, 6*modulus-2].
func (r Ring) NTTLazy(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// INTT evaluates p2 = INTT(p1).
func (r Ring) INTT(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// INTTLazy evaluates p2 = INTT(p1) with p2 in [0, 2*modulus-1].
func (r Ring) INTTLazy(p1, p2 Poly) { _ = "STUB: not implemented"; return }

// butterfly computes X, Y = U + V*Psi, U - V*Psi mod Q.
func butterfly(U, V, Psi, twoQ, fourQ, Q, MRedConstant uint64) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// invbutterfly computes X, Y = U + V, (U - V) * Psi mod Q.
func invbutterfly(U, V, Psi, twoQ, fourQ, Q, MRedConstant uint64) (X, Y uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// At the moment it is not possible to use MRedLazy if Q > 61 bits

// NTTStandard computes the NTTStandard in the given SubRing.
func NTTStandard(p1, p2 []uint64, N int, Q, MRedConstant uint64, BRedConstant [2]uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// NTTStandardLazy computes the NTTStandard in the given SubRing with p2 in [0, 6*modulus-2].
func NTTStandardLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// INTTStandard evaluates p2 = INTTStandard(p1) in the given SubRing.
func INTTStandard(p1, p2 []uint64, N int, NInv, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// INTTStandardLazy evaluates p2 = INTT(p1) in the given SubRing with p2 in [0, 2*modulus-1].
func INTTStandardLazy(p1, p2 []uint64, N int, NInv, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// nttCoreLazy computes the NTT on the input coefficients using the input parameters with output values in the range [0, 6*modulus-2].
func nttCoreLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

func nttLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

func nttUnrolled16Lazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// Copy the result of the first round of butterflies on p2 with approximate reduction

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

// Continue the rest of the second to the n-1 butterflies on p2 with approximate reduction

/* #nosec G115 -- m cannot be negative */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%2 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%2 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%4 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%4 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/*
	for i := uint64(0); i < m; i = i + 8 {

		psi := (*[8]uint64)(unsafe.Pointer(&roots[m+i]))
		x := (*[16]uint64)(unsafe.Pointer(&p2[2*i]))

		V = MRedLazy(x[1], psi[0], Q, MRedConstant)
		x[0], x[1] = x[0]+V, x[0]+twoQ-V

		V = MRedLazy(x[3], psi[1], Q, MRedConstant)
		x[2], x[3] = x[2]+V, x[2]+twoQ-V

		V = MRedLazy(x[5], psi[2], Q, MRedConstant)
		x[4], x[5] = x[4]+V, x[4]+twoQ-V

		V = MRedLazy(x[7], psi[3], Q, MRedConstant)
		x[6], x[7] = x[6]+V, x[6]+twoQ-V

		V = MRedLazy(x[9], psi[4], Q, MRedConstant)
		x[8], x[9] = x[8]+V, x[8]+twoQ-V

		V = MRedLazy(x[11], psi[5], Q, MRedConstant)
		x[10], x[11] = x[10]+V, x[10]+twoQ-V

		V = MRedLazy(x[13], psi[6], Q, MRedConstant)
		x[12], x[13] = x[12]+V, x[12]+twoQ-V

		V = MRedLazy(x[15], psi[7], Q, MRedConstant)
		x[14], x[15] = x[14]+V, x[14]+twoQ-V
	}
*/

func inttCoreLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

func inttLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// Copy the result of the first round of butterflies on p2 with approximate reduction

func inttLazyUnrolled16(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// Copy the result of the first round of butterflies on p2 with approximate reduction

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

// Continue the rest of the second to the n-1 butterflies on p2 with approximate reduction

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%2 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%4 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

// NTTConjugateInvariant evaluates p2 = NTT(p1) in the sub-ring Z[X + X^-1]/(X^2N +1) of Z[X]/(X^2N+1).
func NTTConjugateInvariant(p1, p2 []uint64, N int, Q, MRedConstant uint64, BRedConstant [2]uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// NTTConjugateInvariantLazy evaluates p2 = NTT(p1) in the sub-ring Z[X + X^-1]/(X^2N +1) of Z[X]/(X^2N+1) with p2 in the range [0, 6*modulus-2].
func NTTConjugateInvariantLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// INTTConjugateInvariant evaluates p2 = INTT(p1) in the closed sub-ring Z[X + X^-1]/(X^2N +1) of Z[X]/(X^2N+1).
func INTTConjugateInvariant(p1, p2 []uint64, N int, NInv, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// INTTConjugateInvariantLazy evaluates p2 = INTT(p1) in the closed sub-ring Z[X + X^-1]/(X^2N +1) of Z[X]/(X^2N+1) with p2 in the range [0, 2*modulus-1].
func INTTConjugateInvariantLazy(p1, p2 []uint64, N int, NInv, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// nttCoreConjugateInvariantLazy evaluates p2 = NTT(p1) in the sub-ring Z[X + X^-1]/(X^2N +1) of Z[X]/(X^2N+1) with p2 [0, 6*modulus-2].
func nttCoreConjugateInvariantLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

func nttConjugateInvariantLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

// Continue the rest of the second to the n-1 butterflies on p2 with approximate reduction

func nttConjugateInvariantLazyUnrolled16(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// Copy the result of the first round of butterflies on p2 with approximate reduction

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

// Continue the rest of the second to the n-1 butterflies on p2 with approximate reduction

/* #nosec G115 -- m cannot be negative */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%2 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%2 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%4 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%4 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

// inttCoreConjugateInvariantLazy evaluates p2 = INTT(p1) in the sub-ring Z[X + X^-1]/(X^2N +1) of Z[X]/(X^2N+1) with p2 [0, 2*modulus-1].
func inttCoreConjugateInvariantLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

func inttConjugateInvariantLazy(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"
	return
}

func inttConjugateInvariantLazyUnrolled16(p1, p2 []uint64, N int, Q, MRedConstant uint64, roots []uint64) {
	_ = "STUB: not implemented"

	// Sanity check
	return
}

// Copy the result of the first round of butterflies on p2 with approximate reduction

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%16 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 != 0 */

// Continue the rest of the second to the n-1 butterflies on p2 with approximate reduction

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%2 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(roots)%4 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%16 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 */
