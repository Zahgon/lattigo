// Package main implements an example of smooth function approximation using Chebyshev polynomial interpolation.
package main

import (
	"math"

	"github.com/tuneinsight/lattigo/v6/circuits/ckks/polynomial"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

func main() {
	var err error
	var params ckks.Parameters

	// 128-bit secure parameters enabling depth-7 circuits.
	// LogN:14, LogQP: 431.
	if params, err = ckks.NewParametersFromLiteral(
		ckks.ParametersLiteral{
			LogN:            14,                                    // log2(ring degree)
			LogQ:            []int{55, 45, 45, 45, 45, 45, 45, 45}, // log2(primes Q) (ciphertext modulus)
			LogP:            []int{61},                             // log2(primes P) (auxiliary modulus)
			LogDefaultScale: 45,                                    // log2(scale)
			RingType:        ring.ConjugateInvariant,
		}); err != nil {
		panic(err)
	}

	// Key Generator
	kgen := rlwe.NewKeyGenerator(params)

	// Secret Key
	sk := kgen.GenSecretKeyNew()

	// Encoder
	ecd := ckks.NewEncoder(params)

	// Encryptor
	enc := rlwe.NewEncryptor(params, sk)

	// Decryptor
	dec := rlwe.NewDecryptor(params, sk)

	// Relinearization Key
	rlk := kgen.GenRelinearizationKeyNew(sk)

	// Evaluation Key Set with the Relinearization Key
	evk := rlwe.NewMemEvaluationKeySet(rlk)

	// Evaluator
	eval := ckks.NewEvaluator(params, evk)

	// Samples values in [-K, K]
	K := 25.0

	// Allocates a plaintext at the max level.
	pt := ckks.NewPlaintext(params, params.MaxLevel())

	// Vector of plaintext values
	values := make([]float64, pt.Slots())

	// Populates the vector of plaintext values
	for i := range values {
		values[i] = sampling.RandFloat64(-K, K)
	}

	// Encodes the vector of plaintext values
	if err = ecd.Encode(values, pt); err != nil {
		panic(err)
	}

	// Encrypts the vector of plaintext values
	var ct *rlwe.Ciphertext
	if ct, err = enc.EncryptNew(pt); err != nil {
		panic(err)
	}

	sigmoid := func(x float64) (y float64) {
		return 1 / (math.Exp(-x) + 1)
	}

	// Chebyhsev approximation of the sigmoid in the domain [-K, K] of degree 63.
	poly := polynomial.NewPolynomial(GetChebyshevPoly(K, 63, sigmoid))

	// Instantiates the polynomial evaluator
	polyEval := polynomial.NewEvaluator(params, eval)

	// Retrieves the change of basis y = scalar * x + constant
	scalar, constant := poly.ChangeOfBasis()

	// Performes the change of basis Standard -> Chebyshev
	if err := eval.Mul(ct, scalar, ct); err != nil {
		panic(err)
	}

	if err := eval.Add(ct, constant, ct); err != nil {
		panic(err)
	}

	if err := eval.Rescale(ct, ct); err != nil {
		panic(err)
	}

	// Evaluates the polynomial
	if ct, err = polyEval.Evaluate(ct, poly, params.DefaultScale()); err != nil {
		panic(err)
	}

	// Allocates a vector for the reference values and
	// evaluates the same circuit on the plaintext values
	want := make([]float64, ct.Slots())
	for i := range want {
		want[i], _ = poly.Evaluate(values[i])[0].Float64()
		//want[i] = sigmoid(values[i])
	}

	// Decrypts and print the stats about the precision.
	PrintPrecisionStats(params, ct, want, ecd, dec)
}

// GetChebyshevPoly returns the Chebyshev polynomial approximation of f the
// in the interval [-K, K] for the given degree.
func GetChebyshevPoly(K float64, degree int, f64 func(x float64) (y float64)) bignum.Polynomial {
	_ = "STUB: not implemented"
	return *new(bignum.Polynomial)
}

// Returns the polynomial.

// PrintPrecisionStats decrypts, decodes and prints the precision stats of a ciphertext.
func PrintPrecisionStats(params ckks.Parameters, ct *rlwe.Ciphertext, want []float64, ecd *ckks.Encoder, dec *rlwe.Decryptor) {
	_ = "STUB: not implemented"

	// Decrypts the vector of plaintext values
	return
}

// Decodes the plaintext

// Pretty prints some values

// Pretty prints the precision stats
