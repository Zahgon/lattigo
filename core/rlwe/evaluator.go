package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
)

// Evaluator is a struct that holds the necessary elements to execute general homomorphic
// operation on RLWE ciphertexts, such as automorphisms, key-switching and relinearization.
type Evaluator struct {
	params Parameters
	EvaluationKeySet

	automorphismIndex map[uint64][]uint64

	BasisExtender *ring.BasisExtender
	Decomposer    *ring.Decomposer
	pool          *BufferPool
}

// NewEvaluator creates a new [Evaluator].
func NewEvaluator(params ParameterProvider, evk EvaluationKeySet) (eval *Evaluator) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should not happen.

func (eval *Evaluator) GetRLWEParameters() *Parameters { _ = "STUB: not implemented"; return nil }

// CheckAndGetGaloisKey returns an error if the [GaloisKey] for the given Galois element is missing or the [EvaluationKey] interface is nil.
func (eval Evaluator) CheckAndGetGaloisKey(galEl uint64) (evk *GaloisKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sanity check, this error should not happen.

// CheckAndGetRelinearizationKey returns an error if the [RelinearizationKey] is missing or the [EvaluationKey] interface is nil.
func (eval Evaluator) CheckAndGetRelinearizationKey() (evk *RelinearizationKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InitOutputBinaryOp initializes the output [Element] opOut for receiving the result of a binary operation over
// op0 and op1. The method also performs the following checks:
//
//  1. Inputs are not nil
//  2. MetaData are not nil
//  3. op0.Degree() + op1.Degree() != 0 (i.e at least one [Element] is a ciphertext)
//  4. op0.IsNTT == op1.IsNTT == DefaultNTTFlag
//  5. op0.IsBatched == op1.IsBatched
//
// The opOut metadata are initilized as:
// IsNTT <- DefaultNTTFlag
// IsBatched <- op0.IsBatched
// LogDimensions <- max(op0.LogDimensions, op1.LogDimensions)
//
// The method returns max(op0.Degree(), op1.Degree(), opOut.Degree()) and min(op0.Level(), op1.Level(), opOut.Level())
func (eval Evaluator) InitOutputBinaryOp(op0, op1 *Element[ring.Poly], opInTotalMaxDegree int, opOut *Element[ring.Poly]) (degree, level int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// InitOutputUnaryOp initializes the output [Element] opOut for receiving the result of a unary operation over
// op0. The method also performs the following checks:
//
//  1. Input and output are not nil
//  2. Inoutp and output Metadata are not nil
//  3. op0.IsNTT == DefaultNTTFlag
//
// The method will also update the metadata of opOut:
//
// IsNTT <- NTTFlag
// IsBatched <- op0.IsBatched
// LogDimensions <- op0.LogDimensions
//
// The method returns max(op0.Degree(), opOut.Degree()) and min(op0.Level(), opOut.Level()).
func (eval Evaluator) InitOutputUnaryOp(op0, opOut *Element[ring.Poly]) (degree, level int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// WithKey creates a shallow copy of the receiver [Evaluator] for which the new [EvaluationKey] is evaluationKey
// and where the temporary buffers are shared. The receiver and the returned evaluators cannot be used concurrently.
func (eval Evaluator) WithKey(evk EvaluationKeySet) *Evaluator {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check, this error should not happen.

func (eval Evaluator) AutomorphismIndex(galEl uint64) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

func (eval Evaluator) ModDownQPtoQNTT(levelQ, levelP int, p1Q, p1P, p2Q ring.Poly) {
	_ = "STUB: not implemented"
	return
}
