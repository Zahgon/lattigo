// Package inverse implements a homomorphic inversion circuit for the CKKS scheme.
package inverse

import (
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/bootstrapping"
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/minimax"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// Evaluator is an evaluator used to evaluate the inverses of ciphertexts.
// All fields of this struct are public, enabling custom instantiations.
type Evaluator struct {
	Parameters ckks.Parameters
	*minimax.Evaluator
}

// NewEvaluator instantiates a new InverseEvaluator.
// This method is allocation free.
func NewEvaluator(params ckks.Parameters, eval *minimax.Evaluator) Evaluator {
	_ = "STUB: not implemented"
	return *new(Evaluator)
}

// EvaluateFullDomainNew computes 1/x for x in [-2^{log2max}, -2^{log2min}] U [2^{log2min}, 2^{log2max}].
//  1. Reduce the interval from [-max, -min] U [min, max] to [-1, -min] U [min, 1] by computing an approximate
//     inverse c such that |c * x| <= 1. For |x| > 1, c tends to 1/x while for |x| < c tends to 1.
//     This is done by using the work Efficient Homomorphic Evaluation on Large Intervals (https://eprint.iacr.org/2022/280.pdf).
//  2. Compute |c * x| = sign(x * c) * (x * c), this is required for the next step, which can only accept positive values.
//  3. Compute y' = 1/(|c * x|) with the iterative Goldschmidt division algorithm.
//  4. Compute y = y' * c * sign(x * c)
//
// The user can provide a minimax composite polynomial (signMinimaxPoly) for the sign function in the interval
// [-1-e, -2^{log2min}] U [2^{log2min}, 1+e] (where e is an upperbound on the scheme error).
// If no such polynomial is provided, then the [DefaultMinimaxCompositePolynomialForSign] is used by default.
// Note that the precision of the output of sign(x * c) does not impact the circuit precision since this value ends up being both at
// the numerator and denominator, thus cancelling itself.
func (eval Evaluator) EvaluateFullDomainNew(ct *rlwe.Ciphertext, log2min, log2max float64, signMinimaxPoly ...minimax.Polynomial) (cInv *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EvaluatePositiveDomainNew computes 1/x for x in [2^{log2min}, 2^{log2max}].
//  1. Reduce the interval from [min, max] to [min, 1] by computing an approximate
//     inverse c such that |c * x| <= 1. For |x| > 1, c tends to 1/x while for |x| < c tends to 1.
//     This is done by using the work Efficient Homomorphic Evaluation on Large Intervals (https://eprint.iacr.org/2022/280.pdf).
//  2. Compute y' = 1/(c * x) with the iterative Goldschmidt division algorithm.
//  3. Compute y = y' * c
func (eval Evaluator) EvaluatePositiveDomainNew(ct *rlwe.Ciphertext, log2min, log2max float64) (cInv *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EvaluateNegativeDomainNew computes 1/x for x in [-2^{log2max}, -2^{log2min}].
//  1. Reduce the interval from [-max, -min] to [-1, -min] by computing an approximate
//     inverse c such that |c * x| <= 1. For |x| > 1, c tends to 1/x while for |x| < c tends to 1.
//     This is done by using the work Efficient Homomorphic Evaluation on Large Intervals (https://eprint.iacr.org/2022/280.pdf).
//  2. Compute y' = 1/(c * x) with the iterative Goldschmidt division algorithm.
//  3. Compute y = y' * c
func (eval Evaluator) EvaluateNegativeDomainNew(ct *rlwe.Ciphertext, log2min, log2max float64) (cInv *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (eval Evaluator) evaluateNew(ct *rlwe.Ciphertext, log2min, log2max float64, fulldomain bool, signMinimaxPoly minimax.Polynomial) (cInv *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If max > 1, then normalizes the ciphertext interval from  [-max, -min] U [min, max]
// to [-1, -min] U [min, 1], and returns the encrypted normalization factor.

// Computes the sign with precision [-1, -2^-a] U [2^-a, 1]

// Checks that cInv have at least one level remaining above the minimum
// level required for the bootstrapping.

// Gets |x| = x * sign(x)

// Computes the inverse of x in [min = 2^-a, 1]

// If x > 1 then multiplies back with the encrypted normalization vector

// Multiplies back with the encrypted sign

// GoldschmidtDivisionNew homomorphically computes 1/x in the domain [0, 2].
// input: ct: Enc(x) with values in the interval [0+2^{-log2min}, 2-2^{-log2min}].
// output: Enc(1/x - e), where |e| <= (1-x)^2^(#iterations+1) -> the bit-precision doubles after each iteration.
// This method automatically estimates how many iterations are needed to
// achieve the optimal precision, which is derived from the plaintext scale.
// This method will return an error if the input ciphertext does not have enough
// remaining level and if the InverseEvaluator was instantiated with no bootstrapper.
// This method will return an error if something goes wrong with the bootstrapping or the rescaling operations.
func (eval Evaluator) GoldschmidtDivisionNew(ct *rlwe.Ciphertext, log2min float64) (ctInv *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 2^{-(prec - LogN + 1)}

// Estimates the number of iterations required to achieve the desired precision, given the interval [min, 2-min]

// Doubles the bit-precision at each iteration

// Minimum of 3 iterations
// This minimum is set in the case where log2min is close to 0.

// a is at a higher level than tmp but at the same scale magnitude
// We consume a level to bring a to the same level as tmp

// IntervalNormalization applies a modified version of Algorithm 2 of Efficient Homomorphic Evaluation on Large Intervals (https://eprint.iacr.org/2022/280)
// to normalize the interval from [-max, max] to [-1, 1]. Also returns the encrypted normalization factor.
//
// The original algorithm of https://eprint.iacr.org/2022/280 works by successive evaluation of a function that compresses values greater than some threshold
// to this threshold and let values smaller than the threshold untouched (mostly). The process is iterated, each time reducing the threshold by a pre-defined
// factor L. We can modify the algorithm to keep track of the compression factor so that we can get back the original values (before the compression) afterward.
//
// Given ct with values [-max, max], the method will compute y such that ct * y has values in [-1, 1].
// The normalization factor is independant to each slot:
//   - values smaller than 1 will have a normalization factor that tends to 1
//   - values greater than 1 will have a normalization factor that tends to 1/x
func (eval Evaluator) IntervalNormalization(ct *rlwe.Ciphertext, log2Max float64, btp bootstrapping.Bootstrapper) (ctNorm, ctNormFac *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Compression factor (experimental)

// n = log_{L}(max)

// c = 2/sqrt(27 * L^(2 * (n-1-i)))

// Depth 2
// Computes: z = 1 - (y * c)^2

// 1 level

// (c * y)

// 1 level
// (c * y)^2

// -(c * y)^2

// 1-(c * y)^2

// Updates the normalization factor

// 1 level

// Updates the ciphertext
// 1 level
