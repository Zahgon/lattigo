package ring

import (
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

const ternarySamplerPrecision = uint64(56)

// TernarySampler keeps the state of a polynomial sampler in the ternary distribution.
type TernarySampler struct {
	*baseSampler
	matrixProba  [2][ternarySamplerPrecision - 1]uint8
	matrixValues [][3]uint64
	invDensity   float64
	hw           int
	sample       func(poly Poly, f func(a, b, c uint64) uint64)
}

// NewTernarySampler creates a new instance of TernarySampler from a PRNG, the ring definition and the distribution
// parameters (see type Ternary). If "montgomery" is set to true, polynomials read from this sampler are in Montgomery form.
// WARNING: If the PRNG is deterministic/keyed (of type [sampling.KeyedPRNG]), *concurrent* calls to the sampler will not necessarily result in a deterministic output.
func NewTernarySampler(prng sampling.PRNG, baseRing *Ring, X Ternary, montgomery bool) (ts *TernarySampler, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AtLevel returns an instance of the target TernarySampler to sample at the given level.
// The returned sampler cannot be used concurrently to the original sampler.
func (ts *TernarySampler) AtLevel(level int) Sampler {
	_ = "STUB: not implemented"
	return *new(Sampler)
}

// Read samples a polynomial into pol.
func (ts *TernarySampler) Read(pol Poly) { _ = "STUB: not implemented"; return }

// ReadNew allocates and samples a polynomial at the max level.
func (ts *TernarySampler) ReadNew() (pol Poly) { _ = "STUB: not implemented"; return *new(Poly) }

func (ts *TernarySampler) ReadAndAdd(pol Poly) { _ = "STUB: not implemented"; return }

func (ts *TernarySampler) initializeMatrix(montgomery bool) { _ = "STUB: not implemented"; return }

// [0] = 0
// [1] = 1 * 2^64 mod qi
// [2] = (qi - 1) * 2^64 mod qi

func (ts *TernarySampler) computeMatrixTernary(p float64) { _ = "STUB: not implemented"; return }

/* #nosec G115 -- value is 1 bit */

/* #nosec G115 -- value is 1 bit */

func (ts *TernarySampler) sampleProba(pol Poly, f func(a, b, c uint64) uint64) {
	_ = "STUB: not implemented"

	// Sanity check for invalid parameters
	return
}

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

func (ts *TernarySampler) sampleSparse(pol Poly, f func(a, b, c uint64) uint64) {
	_ = "STUB: not implemented"
	return
}

// We sample ceil(hw/8) bytes

// Sanity check, this error should not happen.

/* #nosec G115 -- N-i and bits(N-i) cannot be negative */
// rejection sampling of a random variable between [0, len(index)]

/* #nosec G115 -- N-i cannot be negative */

// random binary digit [0, 1] from the random bytes (0 = 1, 1 = -1)

// Remove the element in position j of the slice (order not preserved)

// kysampling uses the binary expansion and random bytes matrix to sample a discrete Gaussian value and its sign.
func (ts *TernarySampler) kysampling(prng sampling.PRNG, randomBytes []byte, pointer uint8, bytePointer, byteLength int) (uint64, uint64, []byte, uint8, int) {
	_ = "STUB: not implemented"
	return 0, 0, nil, 0, 0
}

// Use one random byte per cycle and cycle through the randomBytes

// There is small probability that it will get out of the bound, then
// rerun until it gets a proper output

// Sign

// If the last bit of the array was read, sample a new one

// Sanity check, this error should not happen.

// Otherwise, the sign is the next bit of the byte

/* #nosec G115 -- row and sign cannot be negative */

// Reset the bit pointer and discard the used byte

// If the last bit of the array was read, sample a new one

// Sanity check, this error should not happen.
