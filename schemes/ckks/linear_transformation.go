package ckks

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
)

// TraceNew maps X -> sum((-1)^i * X^{i*n+1}) for 0 <= i < N and returns the result on a new ciphertext.
// For log(n) = logSlots.
func (eval Evaluator) TraceNew(ctIn *rlwe.Ciphertext, logSlots int) (opOut *rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Average returns the average of vectors of batchSize elements.
// The operation assumes that ctIn encrypts SlotCount/batchSize sub-vectors of size batchSize.
// It then replaces all values of those sub-vectors by the component-wise average between all the sub-vectors.
// Example for batchSize=4 and slots=8: [{a, b, c, d}, {e, f, g, h}] -> [0.5*{a+e, b+f, c+g, d+h}, 0.5*{a+e, b+f, c+g, d+h}]
// Operation requires log2(SlotCout/batchSize) rotations.
// Required rotation keys can be generated with RotationsForInnerSumLog(batchSize, SlotCount/batchSize).
func (eval Evaluator) Average(ctIn *rlwe.Ciphertext, logBatchSize int, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// pre-multiplication by n^-1

/* #nosec G115 -- n cannot be negative */
