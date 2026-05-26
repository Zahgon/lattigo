package ring

import (
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// UniformSampler wraps a util.PRNG and represents the state of a sampler of uniform polynomials.
type UniformSampler struct {
	*baseSampler
}

// NewUniformSampler creates a new instance of UniformSampler from a PRNG and ring definition.
// WARNING: If the PRNG is deterministic/keyed (of type [sampling.KeyedPRNG]), *concurrent* calls to the sampler will not necessarily result in a deterministic output.
func NewUniformSampler(prng sampling.PRNG, baseRing *Ring) (u *UniformSampler) {
	_ = "STUB: not implemented"
	return nil
}

// AtLevel returns an instance of the target UniformSampler to sample at the given level.
// The returned sampler cannot be used concurrently to the original sampler.
func (u *UniformSampler) AtLevel(level int) Sampler {
	_ = "STUB: not implemented"
	return *new(Sampler)
}

func (u *UniformSampler) Read(pol Poly) { _ = "STUB: not implemented"; return }

func (u *UniformSampler) ReadAndAdd(pol Poly) { _ = "STUB: not implemented"; return }

func (u *UniformSampler) read(pol Poly, f func(a, b, c uint64) uint64) {
	_ = "STUB: not implemented"
	return
}

// Sanity check, this error should not happen.

// Starts by computing the mask

// Iterates for each modulus over each coefficient

// Samples an integer between [0, qi-1]

// Refills the buff if it runs empty

// Sanity check, this error should not happen.

// Reads bytes from the buff

// If the integer is between [0, qi-1], breaks the loop

// ReadNew generates a new polynomial with coefficients following a uniform distribution over [0, Qi-1].
// Polynomial is created at the max level.
func (u *UniformSampler) ReadNew() (pol Poly) { _ = "STUB: not implemented"; return *new(Poly) }

// RandUniform samples a uniform randomInt variable in the range [0, mask] until randomInt is in the range [0, v-1].
// mask needs to be of the form 2^n -1.
func RandUniform(prng sampling.PRNG, v uint64, mask uint64) (randomInt uint64) {
	_ = "STUB: not implemented"
	return 0
}

// randInt32 samples a uniform variable in the range [0, mask], where mask is of the form 2^n-1, with n in [0, 32].
func randInt32(prng sampling.PRNG, mask uint64) uint64 {
	_ = "STUB: not implemented"

	// generate random 4 bytes
	return 0
}

// Sanity check, this error should not happen.

// convert 4 bytes to a uint32

// return required bits

// randInt64 samples a uniform variable in the range [0, mask], where mask is of the form 2^n-1, with n in [0, 64].
func randInt64(prng sampling.PRNG, mask uint64) uint64 {
	_ = "STUB: not implemented"

	// generate random 8 bytes
	return 0
}

// Sanity check, this error should not happen.

// convert 8 bytes to a uint64

// return required bits
