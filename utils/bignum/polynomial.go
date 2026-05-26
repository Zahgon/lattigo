package bignum

import (
	"math/big"
)

type PolynomialBSGS struct {
	MetaData
	Coeffs [][]*Complex
}

func OptimalSplit(logDegree int) (logSplit int) { _ = "STUB: not implemented"; return 0 }

type Polynomial struct {
	MetaData
	Coeffs []*Complex
}

func (p Polynomial) Clone() Polynomial { _ = "STUB: not implemented"; return *new(Polynomial) }

// NewPolynomial creates a new polynomial from the input parameters:
// basis: either `Monomial` or `Chebyshev`
// coeffs: []Complex128, []float64, []*Complex or []*big.Float
// interval: [2]float64{a, b} or *Interval
func NewPolynomial(basis Basis, coeffs interface{}, interval interface{}) Polynomial {
	_ = "STUB: not implemented"
	return *new(Polynomial)
}

// ChangeOfBasis returns change of basis required to evaluate the polynomial
// Change of basis is defined as follow:
//   - Monomial: scalar=1, constant=0.
//   - Chebyshev: scalar=2/(b-a), constant = (-a-b)/(b-a).
func (p *Polynomial) ChangeOfBasis() (scalar, constant *big.Float) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 2 / (b-a)

// (-b-a)/(b-a)

// Depth returns the number of sequential multiplications needed to evaluate the polynomial.
func (p Polynomial) Depth() int { _ = "STUB: not implemented"; return 0 }

// Degree returns the degree of the polynomial.
func (p Polynomial) Degree() int { _ = "STUB: not implemented"; return 0 }

// EvaluateModP evalutes the polynomial modulo p, treating each coefficient as
// integer variables and returning the result as *big.Int in the interval [0, P-1].
func (p Polynomial) EvaluateModP(xInt, PInt *big.Int) (yInt *big.Int) {
	_ = "STUB: not implemented"
	return nil
}

// Evaluate takes x a *big.Float or *big.Complex and returns y = P(x).
// The precision of x is used as reference precision for y.
func (p *Polynomial) Evaluate(x interface{}) (y *Complex) { _ = "STUB: not implemented"; return nil }

// Factorize factorizes p as X^{n} * pq + pr.
func (p Polynomial) Factorize(n int) (pq, pr Polynomial) {
	_ = "STUB: not implemented"
	return *new(Polynomial), *new(Polynomial)
}

// ns a polynomial p such that p = q*C^degree + r.
