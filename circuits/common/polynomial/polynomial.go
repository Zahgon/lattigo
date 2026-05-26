// Package polynomial bundles generic parts of the homomorphic polynomial evaluation circuit.
package polynomial

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// Polynomial is a struct for representing plaintext polynomials
// for their homomorphic evaluation in an encrypted point. The
// type wraps a [bignum.Polynomial] along with several evaluation-
// related parameters.
type Polynomial struct {
	bignum.Polynomial
	MaxDeg int        // Always set to len(Coeffs)-1
	Lead   bool       // Always set to true
	Lazy   bool       // Flag for lazy-relinearization
	Level  int        // Metadata for BSGS polynomial evaluation
	Scale  rlwe.Scale // Metadata for BSGS polynomial evaluation
}

// NewPolynomial returns an instantiated Polynomial for the
// provided [bignum.Polynomial].
func NewPolynomial(poly bignum.Polynomial) Polynomial {
	_ = "STUB: not implemented"
	return *new(Polynomial)
}

// Factorize factorizes p as X^{n} * pq + pr.
func (p Polynomial) Factorize(n int) (pq, pr Polynomial) {
	_ = "STUB: not implemented"
	return *new(Polynomial), *new(Polynomial)
}

// PatersonStockmeyerPolynomial is a struct that stores
// the Paterson Stockmeyer decomposition of a polynomial.
// The decomposition of P(X) is given as sum pi(X) * X^{2^{n}}
// where degree(pi(X)) =~ sqrt(degree(P(X)))
type PatersonStockmeyerPolynomial struct {
	Degree int
	Base   int
	Level  int
	Scale  rlwe.Scale
	Value  []Polynomial
}

// PatersonStockmeyerPolynomial returns the Paterson Stockmeyer polynomial decomposition of the target polynomial.
// The decomposition is done with the power of two basis.
func (p Polynomial) PatersonStockmeyerPolynomial(params rlwe.ParameterProvider, inputLevel int, inputScale, outputScale rlwe.Scale, eval SimEvaluator) PatersonStockmeyerPolynomial {
	_ = "STUB: not implemented"

	// ceil(log2(degree))
	/* #nosec G115 -- degree cannot be negative */
	return *new(PatersonStockmeyerPolynomial)
}

// optimal ratio between degree(pi(X)) et degree(P(X))

// Initializes the simulated polynomial evaluation

// Generates the simulated powers (to get the scaling factors)

// Simulates the homomorphic evaluation with levels and scaling factors to retrieve the scaling factor of each pi(X).

// recursePS is a recursive implementation of a polynomial evaluation via the Paterson Stockmeyer algorithm with a power of two decomposition.
func recursePS(params rlwe.ParameterProvider, logSplit, targetLevel int, p Polynomial, pb SimPowerBasis, outputScale rlwe.Scale, eval SimEvaluator) ([]Polynomial, *SimOperand) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* #nosec G115 -- MaxDeg cannot be negative */

/* #nosec G115 -- Degree cannot be negative */

// This checks that the underlying algorithm behaves as expected, which will always be
// the case, unless the user provides an incorrect custom implementation.

// PolynomialVector is a struct storing a set of polynomials and a mapping that
// indicates on which slot each polynomial has to be independently evaluated.
// For example, if we are given two polynomials P0(X) and P1(X) and the folling mapping: map[int][]int{0:[0, 1, 2], 1:[3, 4, 5]},
// then the polynomial evaluation on a vector [a, b, c, d, e, f, g, h] will evaluate to [P0(a), P0(b), P0(c), P1(d), P1(e), P1(f), 0, 0]
type PolynomialVector struct {
	Value   []Polynomial
	Mapping map[int][]int
}

// NewPolynomialVector instantiates a new [PolynomialVector] from a set of [bignum.Polynomial] and a mapping indicating
// which polynomial has to be evaluated on which slot.
// For example, if we are given two polynomials P0(X) and P1(X) and the following mapping: map[int][]int{0:[0, 1, 2], 1:[3, 4, 5]},
// then the polynomial evaluation on a vector [a, b, c, d, e, f, g, h] will evaluate to [P0(a), P0(b), P0(c), P1(d), P1(e), P1(f), 0, 0]
func NewPolynomialVector(polys []bignum.Polynomial, mapping map[int][]int) (PolynomialVector, error) {
	_ = "STUB: not implemented"
	return *new(PolynomialVector), nil
}

// IsEven returns true if all underlying polynomials are even,
// i.e. all odd powers are zero.
func (p PolynomialVector) IsEven() (even bool) { _ = "STUB: not implemented"; return false }

// IsOdd returns true if all underlying polynomials are odd,
// i.e. all even powers are zero.
func (p PolynomialVector) IsOdd() (odd bool) { _ = "STUB: not implemented"; return false }

// Factorize factorizes the underlying Polynomial vector p into p = polyq * X^{n} + polyr.
func (p PolynomialVector) Factorize(n int) (polyq, polyr PolynomialVector) {
	_ = "STUB: not implemented"
	return *new(PolynomialVector), *new(PolynomialVector)
}

// PatersonStockmeyerPolynomialVector is a struct implementing the
// Paterson Stockmeyer decomposition of a PolynomialVector.
// See [PatersonStockmeyerPolynomial] for additional information.
type PatersonStockmeyerPolynomialVector struct {
	Value   []PatersonStockmeyerPolynomial
	Mapping map[int][]int
}

// PatersonStockmeyerPolynomial returns the Paterson Stockmeyer polynomial decomposition of the target PolynomialVector.
// The decomposition is done with the power of two basis
func (p PolynomialVector) PatersonStockmeyerPolynomial(params rlwe.Parameters, inputLevel int, inputScale, outputScale rlwe.Scale, eval SimEvaluator) PatersonStockmeyerPolynomialVector {
	_ = "STUB: not implemented"
	return *new(PatersonStockmeyerPolynomialVector)
}
