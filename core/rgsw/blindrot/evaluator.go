package blindrot

import (
	"github.com/tuneinsight/lattigo/v6/core/rgsw"
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
)

// Evaluator is a struct that stores the necessary
// data to handle LWE <-> RLWE conversion and
// blind rotations.
type Evaluator struct {
	*rgsw.Evaluator
	paramsBR  rlwe.Parameters
	paramsLWE rlwe.Parameters

	poolMod2N [2]ring.Poly

	accumulator *rlwe.Ciphertext

	galoisGenDiscreteLog map[uint64]int
}

// NewEvaluator instantiates a new [Evaluator].
func NewEvaluator(paramsBR, paramsLWE rlwe.ParameterProvider) (eval *Evaluator) {
	_ = "STUB: not implemented"
	return nil
}

// This flag is always true

// Generates a map for the discrete log of (+/- 1) * GaloisGen^k for 0 <= k < N-1.
// galoisGenDiscreteLog: map[+/-G^{k} mod 2N] = k

// Evaluate extracts on the fly LWE samples and evaluates the provided blind rotation on the LWE.
// testPolyWithSlotIndex : a map with [slot_index] -> blind rotation
// Returns a map[slot_index] -> BlindRotate(ct[slot_index])
func (eval *Evaluator) Evaluate(ct *rlwe.Ciphertext, testPolyWithSlotIndex map[int]*ring.Poly, BRK BlindRotationEvaluationKeySet) (res map[int]*rlwe.Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Switch modulus from Q to 2N and ensure they are odd

// Conversion from Convolution(a, sk) to DotProd(a, sk) for LWE decryption.
// Copy coefficients multiplied by X^{N-1} in reverse order:
// a_{0} -a_{N-1} -a2_{N-2} ... -a_{1}

/* #nosec G115 -- N cannot be negative */

// Switch modulus from Q to 2N

// Line 2 of Algorithm 7 of https://eprint.iacr.org/2022/198
// Acc = (f(X^{-g}) * X^{-g * b}, 0)
/* #nosec G115 -- b is ensured to be small enough */

// use unused buffer because AutomorphismNTT is not in place

// Line 3 of Algorithm 7 https://eprint.iacr.org/2022/198 (Algorithm 3 of https://eprint.iacr.org/2022/198)

// f(X) * X^{b + <a, s>}

// BlindRotateCore implements Algorithm 3 of https://eprint.iacr.org/2022/198
func (eval *Evaluator) BlindRotateCore(a []uint64, acc *rlwe.Ciphertext, BRK BlindRotationEvaluationKeySet) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GaloisElement(k) = GaloisGen^{k} mod 2N

// Maps a[i] to (+/-) g^{k} mod 2N

// Algorithm 3 of https://eprint.iacr.org/2022/198

// Lines 3 to 9 (negative set of a[i] = -g^{k} mod 2N)

// Line 10 (0 in the negative set is 2N)

// Line 12
// acc = acc(X^{-g})

// Lines 13 - 19 (positive set of a[i] = g^{k} mod 2N)

// Lines 20 - 21 (0 in the positive set is 0)

// evaluateFromDiscreteLogSets loops of Algorithm 3 of https://eprint.iacr.org/2022/198
func (eval *Evaluator) evaluateFromDiscreteLogSets(GaloisElement func(k int) (galEl uint64), sets map[int][]int, k, v int, acc *rlwe.Ciphertext, BRK BlindRotationEvaluationKeySet) (int, error) {
	_ = "STUB: not implemented"

	// Checks if k is in the discrete log sets
	return 0, nil
}

// First condition of line 7 or 17

// acc = acc * RGSW(X^{s[j]})

// Second and third conditions of line 7 or 17

// getGaloisElementInverseMap generates a map [(+/-) g^{k} mod 2N] = +/- k
func getGaloisElementInverseMap(GaloisGen uint64, N int) (GaloisGenDiscreteLog map[uint64]int) {
	_ = "STUB: not implemented"
	return nil
}

/* #nosec G115 -- previous check ensures twoN is greater than zero */

/* #nosec G115 -- twoN cannot be negative */

// getDiscreteLogSets returns map[+/-k] = [i...] for a[0 <= i < N] = {(+/-) g^{k} mod 2N for +/- k}
func (eval *Evaluator) getDiscreteLogSets(a []uint64) (discreteLogSets map[int][]int) {
	_ = "STUB: not implemented"
	return nil
}

// Maps (2*N*a[i]/QLWE) to -N/2 < k <= N/2 for a[i] = (+/- 1) * g^{k}

// modSwitchRLWETo2NLvl applies round(x * 2N / Q) to the coefficients of polQ and returns the result on pol2N.
// makeOdd ensures that output coefficients are odd by xoring with 1 (if not already zero).
func (eval *Evaluator) modSwitchRLWETo2NLvl(level int, polQ, pol2N ring.Poly, makeOdd bool) {
	_ = "STUB: not implemented"
	return
}

/* #nosec G115 -- N cannot be negative */
