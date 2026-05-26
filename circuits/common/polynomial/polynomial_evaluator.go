package polynomial

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// CoefficientGetter defines an interface to get the coefficients of a Polynomial.
type CoefficientGetter[T uint64 | *bignum.Complex] interface {
	// GetVectorCoefficient should return a slice []T containing the k-th coefficient
	// of each polynomial of PolynomialVector indexed by its Mapping.
	// See PolynomialVector for additional information about the Mapping.
	GetVectorCoefficient(pol PolynomialVector, k int) (values []T)
	// GetSingleCoefficient should return the k-th coefficient of Polynomial as the type T.
	GetSingleCoefficient(pol Polynomial, k int) (value T)
}

type Evaluator[T uint64 | *bignum.Complex] struct {
	schemes.Evaluator
	CoefficientGetter[T]
}

// Evaluate is a generic and scheme agnostic method to evaluate polynomials on rlwe.Ciphertexts.
func (eval Evaluator[T]) Evaluate(input interface{}, p interface{}, targetScale rlwe.Scale, levelsConsumedPerRescaling int, SimEval SimEvaluator) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* #nosec G115 -- Degree cannot be negative */

// Computes all the powers of two with relinearization
// This will recursively compute and store all powers of two up to 2^logDegree

// Computes the intermediate powers, starting from the largest, without relinearization if possible

// BabyStep is a struct storing the result of a baby-step
// of the Paterson-Stockmeyer polynomial evaluation algorithm.
type BabyStep struct {
	Degree int
	Value  *rlwe.Ciphertext
}

// EvaluatePatersonStockmeyerPolynomialVector evaluates a pre-decomposed PatersonStockmeyerPolynomialVector on a pre-computed power basis [1, X^{1}, X^{2}, ..., X^{2^{n}}, X^{2^{n+1}}, ..., X^{2^{m}}]
func (eval Evaluator[T]) EvaluatePatersonStockmeyerPolynomialVector(poly PatersonStockmeyerPolynomialVector, pb PowerBasis) (res *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Small steps

// eval & cg are not thread-safe

// Loops as long as there is more than one sub-polynomial

// Precomputes the ops to apply in the giant steps loop

// eval is not thread-safe

// Discards processed sub-polynomials

// EvaluateBabyStep evaluates a baby-step of the PatersonStockmeyer polynomial evaluation algorithm, i.e. the inner-product between the precomputed
// powers [1, T, T^2, ..., T^{n-1}] and the coefficients [ci0, ci1, ci2, ..., ci{n-1}].
func (eval Evaluator[T]) EvaluateBabyStep(i int, poly PatersonStockmeyerPolynomialVector, pb PowerBasis) (ct *BabyStep, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transposes the polynomial matrix

// EvaluateGiantStep evaluates a giant-step of the PatersonStockmeyer polynomial evaluation algorithm, which consists
// in combining the baby-steps <[1, T, T^2, ..., T^{n-1}], [ci0, ci1, ci2, ..., ci{n-1}]> together with powers T^{2^k}.
func (eval Evaluator[T]) EvaluateGiantStep(i int, giantSteps []int, babySteps []*BabyStep, pb PowerBasis) (err error) {
	_ = "STUB: not implemented"

	// If we reach the end of the list it means we weren't able to combine
	// the last two sub-polynomials which necessarily implies that that the
	// last one has degree smaller than the previous one and that there is
	// no next polynomial to combine it with.
	// Therefore we update it's degree to the one of the previous one.
	return nil
}

// If two consecutive sub-polynomials, from ascending degree order, have the
// same degree, we combine them.

/* #nosec G115 -- Degree cannot be negative */

// EvaluateMonomial evaluates a monomial of the form a + b * X^{pow} and writes the results in b.
func (eval Evaluator[T]) EvaluateMonomial(a, b, xpow *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EvaluatePolynomialVectorFromPowerBasis evaluates P(ct) = sum c_i * ct^{i}.
func (eval Evaluator[T]) EvaluatePolynomialVectorFromPowerBasis(targetLevel int, pol PolynomialVector, pb PowerBasis, targetScale rlwe.Scale) (res *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"

	// Map[int] of the powers [X^{0}, X^{1}, X^{2}, ...]
	return nil, nil
}

// Retrieve the degree of the highest degree non-zero coefficient
// TODO: optimize for nil/zero coefficients

// Gets the maximum degree of the ciphertexts among the power basis
// TODO: optimize for nil/zero coefficients, odd/even polynomial

// If an index slot is given (either multiply polynomials or masking)

// If the degree of the poly is zero

// Allocates the output ciphertext

// Allocates the output ciphertext

// Loops starting from the highest degree coefficient

// MulScalarAndAdd automatically scales c to match the scale of res.
