// Package lintrans implements homomorphic linear transformations for the CKKS scheme.
package lintrans

import (
	"github.com/tuneinsight/lattigo/v6/circuits/common/lintrans"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// Diagonals is a wrapper of [lintrans.Diagonals].
type Diagonals[T ckks.Float] lintrans.Diagonals[T]

// DiagonalsIndexList returns the list of the non-zero diagonals of the square matrix.
// A non zero diagonals is a diagonal with a least one non-zero element.
func (m Diagonals[T]) DiagonalsIndexList() (indexes []int) { _ = "STUB: not implemented"; return nil }

// Evaluate evaluates the linear transformation on the provided vector.
// add: c = a + b
// muladd: c = c + a * b
func (m Diagonals[T]) Evaluate(vector []T, newVec func(size int) []T, add func(a, b, c []T), muladd func(a, b, c []T)) (res []T) {
	_ = "STUB: not implemented"
	return nil
}

// PermutationMapping is a struct storing
// a mapping: From -> To and a scaling value.
type PermutationMapping[T ckks.Float] struct {
	From    int
	To      int
	Scaling T
}

// Permutation is a struct that defines generic permutations
// over vectors.
//
// For example, given the vector [a, b, c, d] that would be
// mapped to the vector [1b, 2c, 3d, 4a] then the Permutation
// would contain the following map:
// {0: {3, 4}, 1:{0, 1}, 2:{1, 2}, 3:{2, 3}}
type Permutation[T ckks.Float] []PermutationMapping[T]

// GetDiagonals returns the non-zero diagonals of the matrix
// representation of the permutation, which can be used to
// instantiate [Parameters].
func (p Permutation[T]) GetDiagonals(logSlots int) Diagonals[T] {
	_ = "STUB: not implemented"
	return nil
}

// Parameters is a wrapper of [lintrans.Parameters].
type Parameters lintrans.Parameters

// LinearTransformation is a wrapper of [lintrans.LinearTransformation].
type LinearTransformation lintrans.LinearTransformation

// GaloisElements returns the list of Galois elements required to evaluate the linear transformation.
func (lt LinearTransformation) GaloisElements(params rlwe.ParameterProvider) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

// NewTransformation instantiates a new LinearTransformation and is a wrapper of [lintrans.LinearTransformation].
func NewTransformation(params rlwe.ParameterProvider, lt Parameters) LinearTransformation {
	_ = "STUB: not implemented"
	return *new(LinearTransformation)
}

// Encode is a method used to encode a transformation and a wrapper of [lintrans.Encode].
func Encode[T ckks.Float](ecd schemes.Encoder, diagonals Diagonals[T], allocated LinearTransformation) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GaloisElements returns the list of Galois elements required to evaluate the linear transformation.
func GaloisElements(params rlwe.ParameterProvider, lt Parameters) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// Evaluator is an evaluator providing an API to evaluate linear transformations on [rlwe.Ciphertexts].
// All fields of this struct are public, enabling custom instantiations.
type Evaluator struct {
	lintrans.Evaluator
}

// NewEvaluator instantiates a new [Evaluator] from a circuit.EvaluatorForLinearTransformation.
// This method is allocation free.
func NewEvaluator(eval schemes.Evaluator) (linTransEval *Evaluator) {
	_ = "STUB: not implemented"
	return nil
}

// EvaluateNew takes as input a ciphertext ctIn and a linear transformation M and evaluate and returns opOut: M(ctIn).
func (eval Evaluator) EvaluateNew(ctIn *rlwe.Ciphertext, linearTransformation LinearTransformation) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Evaluate takes as input a ciphertext ctIn, a linear transformation M and evaluates opOut: M(ctIn).
func (eval Evaluator) Evaluate(ctIn *rlwe.Ciphertext, linearTransformation LinearTransformation, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EvaluateManyNew takes as input a ciphertext ctIn and a list of linear transformations [M0, M1, M2, ...] and returns opOut:[M0(ctIn), M1(ctIn), M2(ctInt), ...].
func (eval Evaluator) EvaluateManyNew(ctIn *rlwe.Ciphertext, linearTransformations []LinearTransformation) (opOut []*rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EvaluateMany takes as input a ciphertext ctIn, a list of linear transformations [M0, M1, M2, ...] and a list of pre-allocated receiver opOut
// and evaluates opOut: [M0(ctIn), M1(ctIn), M2(ctIn), ...]
func (eval Evaluator) EvaluateMany(ctIn *rlwe.Ciphertext, linearTransformations []LinearTransformation, opOut []*rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EvaluateSequentialNew takes as input a ciphertext ctIn and a list of linear transformations [M0, M1, M2, ...] and returns opOut:...M2(M1(M0(ctIn))
func (eval Evaluator) EvaluateSequentialNew(ctIn *rlwe.Ciphertext, linearTransformations []LinearTransformation) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EvaluateSequential takes as input a ciphertext ctIn and a list of linear transformations [M0, M1, M2, ...] and returns opOut:...M2(M1(M0(ctIn))
func (eval Evaluator) EvaluateSequential(ctIn *rlwe.Ciphertext, linearTransformations []LinearTransformation, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}
