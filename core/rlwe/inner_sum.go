package rlwe

// Trace maps X -> sum((-1)^i * X^{i*n+1}) for n <= i < N
// Monomial X^k vanishes if k is not divisible by (N/n), otherwise it is multiplied by (N/n).
// Ciphertext is pre-multiplied by (N/n)^-1 to remove the (N/n) factor.
// Examples of full Trace for [0 + 1X + 2X^2 + 3X^3 + 4X^4 + 5X^5 + 6X^6 + 7X^7]
//
// 1.
//
//	  [1 + 2X + 3X^2 + 4X^3 + 5X^4 + 6X^5 + 7X^6 + 8X^7]
//	+ [1 - 6X - 3X^2 + 8X^3 + 5X^4 + 2X^5 - 7X^6 - 4X^7]  {X-> X^(i * 5^1)}
//	= [2 - 4X + 0X^2 +12X^3 +10X^4 + 8X^5 - 0X^6 + 4X^7]
//
// 2.
//
//	  [2 - 4X + 0X^2 +12X^3 +10X^4 + 8X^5 - 0X^6 + 4X^7]
//	+ [2 + 4X + 0X^2 -12X^3 +10X^4 - 8X^5 + 0X^6 - 4X^7]  {X-> X^(i * 5^2)}
//	= [4 + 0X + 0X^2 - 0X^3 +20X^4 + 0X^5 + 0X^6 - 0X^7]
//
// 3.
//
//	  [4 + 0X + 0X^2 - 0X^3 +20X^4 + 0X^5 + 0X^6 - 0X^7]
//	+ [4 + 0X + 0X^2 - 0X^3 -20X^4 + 0X^5 + 0X^6 - 0X^7]  {X-> X^(i * -1)}
//	= [8 + 0X + 0X^2 - 0X^3 + 0X^4 + 0X^5 + 0X^6 - 0X^7]
//
// The method will return an error if the input and output ciphertexts degree is not one.
func (eval Evaluator) Trace(ctIn *Ciphertext, logN int, opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// We skip the last step that applies phi(5^{-1})

/* #nosec G115 -- gap cannot be negative */

// pre-multiplication by (N/n)^-1

// GaloisElementsForTrace returns the list of Galois elements required for the for the `Trace` operation.
// Trace maps X -> sum((-1)^i * X^{i*n+1}) for 2^{LogN} <= i < N.
func GaloisElementsForTrace(params ParameterProvider, logN int) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// PartialTracesSum applies a set of automorphisms on the input ciphertext and sum the results.
// The automorphisms are of the form phi(i*offset, X), 0 <= i < n, where phi(k, X): X -> X^{5^k}
// i.e. opOut = \sum_{i = 0}^{n-1} phi(i*offset, ctIn).
// At the scheme level, this function is used to perform inner sums or efficiently replicate slots.
func (eval Evaluator) PartialTracesSum(ctIn *Ciphertext, offset, n int, opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Accumulator mod QP (i.e. opOut Mod QP)

// Buffer mod QP (i.e. to store the result of lazy gadget products)

// Buffer mod Q (i.e. to store the result of gadget products)

// Sanity check, this error should not happen unless the
// evaluator's buffer has been improperly tempered with.

// Binary reading of the input n

// Starts by decomposing the input ciphertext

// If the binary reading scans a 1 (j is odd)

// If the rotation is not zero

// opOutQP = opOutQP + Rotate(ctInNTT, k)

// j is even

// if n is not a power of two, then at least one j was odd, and thus the buffer opOutQP is not empty

// opOut = opOutQP/P + ctInNTT
// Division by P
// Division by P

// ctInNTT = ctInNTT + Rotate(ctInNTT, 2^i)

// InnerFunction applies an user defined function on the [Ciphertext] with a tree-like combination requiring log2(n) + HW(n) rotations.
//
// InnerFunction with f = eval.Add(a, b, c) is equivalent to [Evaluator.InnerSum] (although slightly slower).
//
// The operation assumes that `ctIn` encrypts Slots/`batchSize` sub-vectors of size `batchSize` and will add them together (in parallel) in groups of `n`.
// It outputs in opOut a [Ciphertext] for which the "leftmost" sub-vector of each group is equal to the pair-wise recursive evaluation of
// function over the group.
//
// The inner function is computed in a tree fashion. Example for batchSize=2 & n=4 (garbage slots are marked by 'x'):
//
//  1. [{a, b}, {c, d}, {e, f}, {g, h}, {a, b}, {c, d}, {e, f}, {g, h}]
//
//  2. [{a, b}, {c, d}, {e, f}, {g, h}, {a, b}, {c, d}, {e, f}, {g, h}]
//     f
//     [{c, d}, {e, f}, {g, h}, {x, x}, {c, d}, {e, f}, {g, h}, {x, x}] (rotate batchSize * 2^{0})
//     =
//     [{f(a, c), f(b, d)}, {f(c, e), f(d, f)}, {f(e, g), f(f, h)}, {x, x}, {f(a, c), f(b, d)}, {f(c, e), f(d, f)}, {f(e, g), f(f, h)}, {x, x}]
//
//  3. [{f(a, c), f(b, d)}, {x, x}, {f(e, g), f(f, h)}, {x, x}, {f(a, c), f(b, d)}, {x, x}, {f(e, g), f(f, h)}, {x, x}] (rotate batchSize * 2^{1})
//     +
//     [{f(e, g), f(f, h)}, {x, x}, {x, x}, {x, x}, {f(e, g), f(f, h)}, {x, x}, {x, x}, {x, x}] =
//     =
//     [{f(f(a,c),f(e,g)), f(f(b, d), f(f, h))}, {x, x}, {x, x}, {x, x}, {f(f(a,c),f(e,g)), f(f(b, d), f(f, h))}, {x, x}, {x, x}, {x, x}]
func (eval Evaluator) InnerFunction(ctIn *Ciphertext, batchSize, n int, f func(a, b, c *Ciphertext) (err error), opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Accumulator mod Q

// Sanity check, this error should not happen unless the
// evaluator's buffer has been improperly tempered with.

// Buffer mod Q

// Sanity check, this error should not happen unless the
// evaluator's buffer has been improperly tempered with.

// Binary reading of the input n

// If the binary reading scans a 1 (j is odd)

// If the rotation is not zero

// opOutQ = f(opOutQ, Rotate(ctInNTT, k), opOutQ)

// j is even

// if n is not a power of two, then at least one j was odd, and thus the buffer opOutQ is not empty

// ctInNTT = f(ctInNTT, Rotate(ctInNTT, 2^i), ctInNTT)

// GaloisElementsForInnerSum returns the list of Galois elements necessary to apply the method
// [Evaluator.InnerSum] operation with parameters batch and n.
func GaloisElementsForInnerSum(params ParameterProvider, batch, n int) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// Replicate applies an optimized replication on the [Ciphertext] (log2(n) + HW(n) rotations with double hoisting).
// It acts as the inverse of a inner sum (summing elements from left to right).
// The replication is parameterized by the size of the sub-vectors to replicate batchSize and
// the number of times n they need to be replicated.
// To ensure correctness, a gap of zero values of size batchSize * (n-1) must exist between
// two consecutive sub-vectors to replicate.
// This method is faster than Replicate when the number of rotations is large and it uses log2(n) + HW(n) instead of n.
func (eval Evaluator) Replicate(ctIn *Ciphertext, batchSize, n int, opOut *Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GaloisElementsForReplicate returns the list of Galois elements necessary to perform the
// [Evaluator.Replicate] operation with parameters batch and n.
func GaloisElementsForReplicate(params ParameterProvider, batch, n int) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}
