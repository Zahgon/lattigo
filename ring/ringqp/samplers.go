package ringqp

import (
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// UniformSampler is a type for sampling polynomials in Ring.
type UniformSampler struct {
	samplerQ, samplerP *ring.UniformSampler
}

// NewUniformSampler instantiates a new UniformSampler from a given PRNG.
// WARNING: If the PRNG is deterministic/keyed (of type [sampling.KeyedPRNG]), *concurrent* calls to the sampler will not necessarily result in a deterministic output.
func NewUniformSampler(prng sampling.PRNG, r Ring) (s UniformSampler) {
	_ = "STUB: not implemented"
	return *new(UniformSampler)
}

// AtLevel returns a shallow copy of the target sampler that operates at the specified levels.
func (s UniformSampler) AtLevel(levelQ, levelP int) UniformSampler {
	_ = "STUB: not implemented"
	return *new(UniformSampler)
}

// Read samples a new polynomial with uniform distribution and stores it into p.
func (s UniformSampler) Read(p Poly) { _ = "STUB: not implemented"; return }

// ReadNew samples a new polynomial with uniform distribution and returns it.
func (s UniformSampler) ReadNew() (p Poly) { _ = "STUB: not implemented"; return *new(Poly) }
