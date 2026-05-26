package ring

import (
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

const (
	discreteGaussianName = "DiscreteGaussian"
	ternaryDistName      = "Ternary"
	uniformDistName      = "Uniform"
)

// Sampler is an interface for random polynomial samplers.
// It has a single Read method which takes as argument the polynomial to be
// populated according to the Sampler's distribution.
type Sampler interface {
	Read(pol Poly)
	ReadNew() (pol Poly)
	ReadAndAdd(pol Poly)
	AtLevel(level int) Sampler
}

// DistributionParameters is an interface for distribution
// parameters in the ring.
// There are three implementation of this interface:
//   - DiscreteGaussian for sampling polynomials with discretized
//     gaussian coefficient of given standard deviation and bound.
//   - Ternary for sampling polynomials with coefficients in [-1, 1].
//   - Uniform for sampling polynomial with uniformly random
//     coefficients in the ring.
type DistributionParameters interface {
	// Type returns a string representation of the distribution name.
	Type() string
	mustBeDist()
}

// DiscreteGaussian represents the parameters of a
// discrete Gaussian distribution with standard
// deviation Sigma and bounds [-Bound, Bound].
type DiscreteGaussian struct {
	Sigma float64
	Bound float64
}

// Ternary represent the parameters of a distribution with coefficients
// in [-1, 0, 1]. Only one of its field must be set to a non-zero value:
//
//   - If P is set, each coefficient in the polynomial is sampled in [-1, 0, 1]
//     with probabilities [0.5*P, 1-P, 0.5*P].
//   - if H is set, the coefficients are sampled uniformly in the set of ternary
//     polynomials with H non-zero coefficients (i.e., of hamming weight H).
type Ternary struct {
	P float64
	H int
}

// Uniform represents the parameters of a uniform distribution
// i.e., with coefficients uniformly distributed in the given ring.
type Uniform struct{}

// NewSampler returns a new sampler that follows the distribution given by DistributionParameters.
// WARNING: If the PRNG is deterministic/keyed (of type [sampling.KeyedPRNG]), *concurrent* calls to the sampler will not necessarily result in a deterministic output.
func NewSampler(prng sampling.PRNG, baseRing *Ring, X DistributionParameters, montgomery bool) (Sampler, error) {
	_ = "STUB: not implemented"
	return *new(Sampler), nil
}

type baseSampler struct {
	prng     sampling.PRNG
	baseRing *Ring
}

// AtLevel returns an instance of the target base sampler that operates at the target level.
// This instance is not thread safe and cannot be used concurrently to the base instance.
func (b baseSampler) AtLevel(level int) *baseSampler { _ = "STUB: not implemented"; return nil }

func (d DiscreteGaussian) Type() string { _ = "STUB: not implemented"; return "" }

func (d DiscreteGaussian) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d DiscreteGaussian) mustBeDist() { _ = "STUB: not implemented"; return }

func (d Ternary) Type() string { _ = "STUB: not implemented"; return "" }

func (d Ternary) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d Ternary) mustBeDist() { _ = "STUB: not implemented"; return }

func (d Uniform) Type() string { _ = "STUB: not implemented"; return "" }

func (d Uniform) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d Uniform) mustBeDist() { _ = "STUB: not implemented"; return }

func getFloatFromMap(distDef map[string]interface{}, key string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getIntFromMap(distDef map[string]interface{}, key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ParametersFromMap(distDef map[string]interface{}) (DistributionParameters, error) {
	_ = "STUB: not implemented"
	return *new(DistributionParameters), nil
}

// a zero value for both P and H is interpreted as an unset value
