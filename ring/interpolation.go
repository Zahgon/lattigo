package ring

// Interpolator is a struct storing the necessary
// buffer and pre-computation for polynomial interpolation
// with coefficient in finite fields.
type Interpolator struct {
	r *Ring
	x Poly
}

// NewInterpolator creates a new Interpolator. Returns an error if T is not
// prime or not congruent to 1 mod 2N, where N is the next power of two greater
// than degree+1.
func NewInterpolator(degree int, T uint64) (itp *Interpolator, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NTT(x)

// Interpolate takes a list of roots the coefficients of P(roots) = 0 mod T.
func (itp *Interpolator) Interpolate(roots []uint64) (coeffs []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// res = NTT(x-root[0])

// res = res * (x-root[i])

// Lagrange takes as input (x, y) and returns P(xi) = yi mod T.
func (itp *Interpolator) Lagrange(x, y []uint64) (coeffs []uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Powers of w are stored in bit-reversed order -> even powers of w are on the right n half
// -> map that stores all the roots of X^{N} + 1 mod T

// Computes the Lagrange basis (X-x[0]) * (X-x[1]) * ... * (X-x[i])
// but omits x[i] which are roots of X^{N} + 1 mod T.
// The roots of X^{N} + 1 mod T are the even powers of w, where w is
// is a primitive 2N-th roots of unity mod T.

// If x[i] is a root of X^{N} + 1 mod T then it is not part
// of the Lagrange basis pre-computation, so all we need is
// to add the missing roots (if any), skipping x[i].

// with the missing roots, except x[i]

// If x[i] is not a root of X^{N} + 1 mod T, then we need
// to remove it from the Lagrange basis pre-computation.
// But first we add the missing x[i], which are the
// roots of X^{N} + 1 mod T (if any).

// Continue with all the missing roots

// And then removes (X - x[i])

// TODO: unrol loop and use unsafe

/* #nosec G115 -- library requires 64-bit system -> int = int64 */

// prod(x[i] - x[j]) i != j
// TODO: make 2 iterations to avoid the if condition

// 1 / prod(x[i] - x[j])

// y[i] / prod(x[i] - x[j])

// P(X) += (y[i] / prod(x[i] - x[j])) * prod(X-x[j])

// computes p3 = (p1 - a) * p2
func subScalarMontgomeryAndMulCoeffsMontgomery(p1 []uint64, a uint64, p2, p3 []uint64, t, mredParams uint64) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p1)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p2)%8 != 0 */

/* #nosec G103 -- behavior and consequences well understood, possible buffer overflow if len(p3)%8 != 0 */
