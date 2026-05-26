package rlwe

import (
	"github.com/tuneinsight/lattigo/v6/ring"
)

// RingPackingEvaluator is an evaluator for Ring-LWE packing operations.
// All fields of this struct are public, enabling custom instantiations.
type RingPackingEvaluator struct {
	*RingPackingEvaluationKey

	Evaluators map[int]*Evaluator

	//XPow2NTT: [1, x, x^2, x^4, ..., x^2^s] / (X^2^s +1)
	XPow2NTT map[int][]ring.Poly

	//XInvPow2NTT: [1, x^-1, x^-2, x^-4, ..., x^-2^s/2] / (X^2^s +1)
	XInvPow2NTT map[int][]ring.Poly
}

// NewRingPackingEvaluator instantiates a new RingPackingEvaluator from a RingPackingEvaluationKey.
func NewRingPackingEvaluator(evk *RingPackingEvaluationKey) *RingPackingEvaluator {
	_ = "STUB: not implemented"
	return nil
}

// Extract takes as input a ciphertext encrypting P(X) = c[i] * X^i and returns a map of
// ciphertexts of degree eval.MinLogN(), each encrypting P(X) = c[i] * X^{0} for i in idx.
// All non-constant coefficients are zeroed and thus correctness is ensured if this method
// is composed with either Repack or RepackNaive.
func (eval RingPackingEvaluator) Extract(ct *Ciphertext, idx map[int]bool) (cts map[int]*Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractNaive takes as input a ciphertext encrypting P(X) = c[i] * X^i and returns a map of
// ciphertexts of degree eval.MinLogN(), each encrypting P(X) = c[i] * X^{0} for i in idx.
// Non-constant coefficients are NOT zeroed thus correctness is only ensured if this method
// is composed with Repack.
//
// If eval.MinLogN() = eval.MaxLogN(), no evaluation keys are required for this method.
// If eval.MinLogN() < eval.MaxLogN(), only RingSwitchingKeys are required for this method.
func (eval RingPackingEvaluator) ExtractNaive(ct *Ciphertext, idx map[int]bool) (cts map[int]*Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If naive = false, then all non-constant coefficients are zeroed.
func (eval RingPackingEvaluator) extract(ct *Ciphertext, idx map[int]bool, naive bool) (cts map[int]*Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First recursively splits the ciphertexts into smaller ciphertexts of half the ring
// degree until the minimum ring degre is reached

// Each split of the ring divides the gap a factor of two

// Applies the same split on the index map, but also update the
// indexes to take into account the new ordering

// For each small ciphertext, extracts the relevant values

// Rotates ciphertexts to move c[i] * X^{i} -> c[i] * X^{0}
// by sequentially multplying with the appropriate X^{-2^{i}}.

// Split splits a ciphertext of degree N into two ciphertexts of degree N/2:
// ctN[X] = ctEvenNHalf[Y] + X * ctOddNHalf[Y] where Y = X^2.
func (eval RingPackingEvaluator) Split(ctN, ctEvenNHalf, ctOddNHalf *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// SkN -> SkNHalf

// Maps to smaller ring degree X -> Y = X^{2}

// Maps to smaller ring degree X -> Y = X^{2}

// SplitNew splits a ciphertext of degree N into two ciphertexts of degree N/2:
// ctN[X] = ctEvenNHalf[Y] + X * ctOddNHalf[Y] where Y = X^2.
func (eval RingPackingEvaluator) SplitNew(ctN *Ciphertext) (ctEvenNHalf, ctOddNHalf *Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Repack takes as input a map of ciphertext and repacks the constant coefficient each ciphertext
// into a single ciphertext of degree eval.MaxLogN() following the indexing of the map.
//
// For example, if cts = map[int]*Ciphertext{0:ct0, 1:ct1, 4:ct2}, then the method will return
// a ciphertext encrypting P(X) = ct0[0] + ct1[0] * X + ct2[0] * X^4.
//
// The method accepts ciphertexts of a ring degree between eval.MinLogN() and eval.MaxLogN().
//
// All non-constant coefficient are zeroed during the repacking, thus correctness is ensured if this
// method can be composed with either Extract or ExtractNaive.
func (eval RingPackingEvaluator) Repack(cts map[int]*Ciphertext) (ct *Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil

	// RepackNaive takes as input a map of ciphertext and repacks the constant coefficient each ciphertext
	// into a single ciphertext of degree eval.MaxLogN() following the indexing of the map.
	//
	// For example, if cts = map[int]*Ciphertext{0:ct0, 1:ct1, 4:ct2}, then the method will return
	// a ciphertext encrypting P(X) = ct0[0] + ct1[0] * X + ct2[0] * X^4.
	//
	// The method accepts ciphertexts of a ring degree between eval.MinLogN() and eval.MaxLogN().
	//
	// If eval.MinLogN() = eval.MaxLogN(), no evaluation keys are required for this method.
	// If eval.MinLogN() < eval.MaxLogN(), only RingSwitchingKeys are required for this method.
	//
	// Unlike Repack, non-constant coefficient are NOT zeroed during the repacking, thus correctness is only
	// ensured if this method is composed with either Extract.
}

func (eval RingPackingEvaluator) RepackNaive(cts map[int]*Ciphertext) (ct *Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (eval RingPackingEvaluator) repack(cts map[int]*Ciphertext, naive bool) (ct *Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List of map containing the repacking of cts

// Assigns to each map the corresponding ciphertext.
// This takes into account the future merging, that merges
// ciphertexts in a base-2 tree-like fashion by evaluating
// ctN[X] = ctEvenNHalf[Y] + X * ctOddNHalf[Y] where Y = X^2.

// Map of repacked ciphertext that will then be merged together.
// Each merging takes two ciphertexts, doubles their ring degree
// and adds them together.

//X^(N/2^L)

// a = a + b * X^{N/2^{i}}

// if ct[jx] == nil, then simply re-assigns

// Merges the cipehrtexts in a base-2 tree like fashion.

// Merge merges two ciphertexts of degree N/2 into a ciphertext of degre N:
// ctN[X] = ctEvenNHalf[Y] + X * ctOddNHalf[Y] where Y = X^2.
func (eval RingPackingEvaluator) Merge(ctEvenNHalf, ctOddNHalf, ctN *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// SkNHalf -> SkN

// MergeNew merges two ciphertexts of degree N/2 into a ciphertext of degre N:
// ctN[X] = ctEvenNHalf[Y] + X * ctOddNHalf[Y] where Y = X^2.
func (eval RingPackingEvaluator) MergeNew(ctEvenNHalf, ctOddNHalf *Ciphertext) (ctN *Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expand expands a RLWE Ciphertext encrypting P(X) = ci * X^i and returns a map of
// ciphertexts, each encrypting P(X) = ci * X^0, indexed by i, for 0<= i < 2^{logN}
// and i divisible by 2^{logGap}.
//
// This method is a used as a sub-routine of the Extract method.
//
// The method will return an error if:
//   - The input ciphertext degree is not one
//   - The ring type is not ring.Standard
func (eval RingPackingEvaluator) Expand(ct *Ciphertext, logGap int) (cts map[int]*Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Multiplies by 2^{-logN} mod Q

/* #nosec G115 -- N and/or n cannot be negative */

// X -> X^{N/n + 1}
//[a, b, c, d] -> [a, -b, c, -d]

// Zeroes odd coeffs: [a, b, c, d] + [a, -b, c, -d] -> [2a, 0, 2b, 0]

// Zeroes even coeffs: [a, b, c, d] - [a, -b, c, -d] -> [0, 2b, 0, 2d]

// c1 * X^{-2^{i}}: [0, 2b, 0, 2d] * X^{-n} -> [2b, 0, 2d, 0]

// Zeroes odd coeffs: [a, b, c, d] + [a, -b, c, -d] -> [2a, 0, 2b, 0]

// Pack packs a map of of ciphertexts, each encrypting Pi(X) = ci * X^{i} for 0 <= i * 2^{inputLogGap} < 2^{LogN}
// and indexed by j, for 0<= j < 2^{eval.MaxLogN()} and returns ciphertext encrypting P(X) = Pi(X) * X^i.
// zeroGarbageSlots: if set to true, slots which are not multiples of X^{2^{logGap}} will be zeroed during the procedure.
//
// The method will return an error if:
//   - The number of ciphertexts is 0
//   - Any input ciphertext degree is not one
//   - Gaps between ciphertexts is smaller than inputLogGap > N
//   - The ring type is not ring.Standard
//
// Example: we want to pack 4 ciphertexts into one, and keep only coefficients which are a multiple of X^{4}.
//
//	To do so, we must set logGap = 2.
//	Here the `X` slots are treated as garbage slots that we want to discard during the procedure.
//
//	input: map[int]{
//	   0: [x00, X, X, X, x01, X, X, X],   with logGap = 2
//	   1: [x10, X, X, X, x11, X, X, X],
//	   2: [x20, X, X, X, x21, X, X, X],
//	   3: [x30, X, X, X, x31, X, X, X],
//		}
//
//	 Step 1:
//	         map[0]: 2^{-1} * (map[0] + X^2 * map[2] + phi_{5^2}(map[0] - X^2 * map[2]) = [x00, X, x20, X, x01, X, x21, X]
//	         map[1]: 2^{-1} * (map[1] + X^2 * map[3] + phi_{5^2}(map[1] - X^2 * map[3]) = [x10, X, x30, X, x11, X, x31, X]
//	 Step 2:
//	         map[0]: 2^{-1} * (map[0] + X^1 * map[1] + phi_{5^4}(map[0] - X^1 * map[1]) = [x00, x10, x20, x30, x01, x11, x21, x22]
func (eval RingPackingEvaluator) Pack(cts map[int]*Ciphertext, inputLogGap int, zeroGarbageSlots bool) (ct *Ciphertext, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//X^(N/2^L)

// tmpa = phi(a - b * X^{N/2^{i}}, 2^{i-1})

// a = a + b * X^{N/2^{i}}

// if ct[jx] == nil, then simply re-assigns

// a + b * X^{N/2^{i}} + phi(a - b * X^{N/2^{i}}, 2^{i-1})

// b * X^{N/2^{i}} - phi(b * X^{N/2^{i}}, 2^{i-1}))

// GenXPow2NTT generates X^({-1 if div else 1} * {2^{0 <= i < LogN}}) in NTT.
func GenXPow2NTT(r *ring.Ring, logN int, div bool) (xPow []ring.Poly) {
	_ = "STUB: not implemented"

	// Compute X^{-n} from 0 to LogN
	return nil
}

// X^{n} = X^{1} * X^{n-1}

func getMinimumGap(list []int) (gap, logGap int, err error) {
	_ = "STUB: not implemented"

	// The loops over to find the smallest gap
	return 0, 0, nil
}

// 2^{63}-1

// Sets gap to the largest power-of-two that divides it.
// We will then discart all coefficients that are not a
// multiple of this gap (and thus possibly entire ciph-
// ertexts).
