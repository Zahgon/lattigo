// Package minimax implements a homomorphic minimax circuit for the CKKS scheme.
package minimax

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// Polynomial is a struct storing P(x) = pk(x) o pk-1(x) o ... o p1(x) o p0(x).
type Polynomial []bignum.Polynomial

// NewPolynomial creates a new Polynomial from a list of coefficients.
// Coefficients are expected to be given in the Chebyshev basis.
func NewPolynomial(coeffsStr [][]string) Polynomial {
	_ = "STUB: not implemented"
	return *new(Polynomial)
}

func (mcp Polynomial) MaxDepth() (depth int) { _ = "STUB: not implemented"; return 0 }

func (mcp Polynomial) Evaluate(x interface{}) (y *bignum.Complex) {
	_ = "STUB: not implemented"
	return nil
}

// CoeffsSignX2Cheby (from https://eprint.iacr.org/2019/1234.pdf) are the coefficients
// of 1.5*x - 0.5*x^3 in Chebyshev basis.
// Evaluating this polynomial on values already close to -1, or 1 ~doubles the number of
// of correct digits.
// For example, if x = -0.9993209 then p(x) = -0.999999308
// This polynomial can be composed after the minimax composite polynomial to double the
// output precision (up to the scheme precision) each time it is evaluated.
var CoeffsSignX2Cheby = []string{"0", "1.125", "0", "-0.125"}

// CoeffsSignX4Cheby (from https://eprint.iacr.org/2019/1234.pdf) are the coefficients
// of 35/16 * x - 35/16 * x^3 + 21/16 * x^5 - 5/16 * x^7 in Chebyshev basis.
// Evaluating this polynomial on values already close to -1, or 1 ~quadruples the number of
// of correct digits.
// For example, if x = -0.9993209 then p(x) = -0.9999999999990705
// This polynomial can be composed after the minimax composite polynomial to quadruple the
// output precision (up to the scheme precision) each time it is evaluated.
var CoeffsSignX4Cheby = []string{"0", "1.1962890625", "0", "-0.2392578125", "0", "0.0478515625", "0", "-0.0048828125"}

// GenMinimaxCompositePolynomialForSign generates the minimax composite polynomial
// P(x) = pk(x) o pk-1(x) o ... o p1(x) o p0(x) of the sign function in their interval
// [min-err, -2^{-alpha}] U [2^{-alpha}, max+err] where alpha is the desired distinguishing
// precision between two values and err an upperbound on the scheme error.
//
// The sign function is defined as: -1 if -1 <= x < 0, 0 if x = 0, 1 if 0 < x <= 1.
//
// See [GenMinimaxCompositePolynomial] for information about how to instantiate and
// parameterize each input value of the algorithm.
func GenMinimaxCompositePolynomialForSign(prec uint, logalpha, logerr int, deg []int) {
	_ = "STUB: not implemented"
	return
}

// GenMinimaxCompositePolynomial generates the minimax composite polynomial
// P(x) = pk(x) o pk-1(x) o ... o p1(x) o p0(x) for the provided function in the interval
// in their interval [min-err, -2^{-alpha}] U [2^{-alpha}, max+err] where alpha is
// the desired distinguishing precision between two values and err an upperbound on
// the scheme error.
//
// The user must provide the following inputs:
//   - prec: the bit precision of the big.Float values used by the algorithm to compute the polynomials.
//     This will impact the speed of the algorithm.
//     A too low precision can prevent convergence or induce a slope zero during the zero finding.
//     A sign that the precision is too low is when the iteration continue without the error getting smaller.
//   - logalpha: log2(alpha)
//   - logerr: log2(err), the upperbound on the scheme precision. Usually this value should be smaller or equal to logalpha.
//     Correctly setting this value is mandatory for correctness, because if x is outside of the interval
//     (i.e. smaller than -1-e or greater than 1+e), then the values will explode during the evaluation.
//     Note that it is not required to apply change of interval [-1, 1] -> [-1-e, 1+e] because the function to evaluate
//     is the sign (i.e. it will evaluate to the same value).
//   - deg: the degree of each polynomial, ordered as follow [deg(p0(x)), deg(p1(x)), ..., deg(pk(x))].
//     It is highly recommended that deg(p0) <= deg(p1) <= ... <= deg(pk) for optimal approximation.
//
// The polynomials are returned in the Chebyshev basis and pre-scaled for
// the interval [-1, 1] (no further scaling is required on the ciphertext).
//
// Be aware that finding the minimax polynomials can take a while (in the order of minutes for high precision when using large degree polynomials).
//
// The function will print information about each step of the computation in real time so that it can be monitored.
//
// The underlying algorithm use the multi-interval Remez algorithm of https://eprint.iacr.org/2020/834.pdf.
func GenMinimaxCompositePolynomial(prec uint, logalpha, logerr int, deg []int, f func(*big.Float) *big.Float) (coeffs [][]*big.Float) {
	_ = "STUB: not implemented"
	return nil
}

// Precision of the output value of the sign polynomial

// Expected upperbound scheme error

// Maximum number of iterations

// Scan step for finding zeroes of the error function

// Interval [-1, alpha] U [alpha, 1]

// Adds the error to the interval
// [A, -alpha] U [alpha, B] becomes [A-e, -alpha] U [alpha, B+e]

// Parameters of the minimax approximation

//r.ShowCoeffs(decimals)

// New interval as [-(1+max_err), -(1-min_err)] U [1-min_err, 1+max_err]

// Extends the new interval by the scheme error
// [-(1+max_err), -(1-min_err)] U [1-min_err, 1 + max_err] becomes [-(1+max_err+e), -(1-min_err-e)] U [1-min_err-e, 1+max_err+e]

// Interval normalization

//r.ShowCoeffs(decimals)

// Since this is the last polynomial, we can skip the interval scaling.

// PrettyPrintCoefficients prints the coefficients formatted.
// If odd = true, even coefficients are zeroed.
// If even = true, odd coefficients are zeroed.
func PrettyPrintCoefficients(decimals int, coeffs []*big.Float, odd, even, first bool) {
	_ = "STUB: not implemented"
	return
}

func parseCoeffs(coeffsStr []string) (coeffs []*big.Float) { _ = "STUB: not implemented"; return nil }

// max(float64, digits * log2(10))
