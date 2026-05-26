package ring

// SubRing is a struct storing precomputation
// for fast modular reduction and NTT for
// a given modulus.
type SubRing struct {
	ntt NumberTheoreticTransformer

	// Polynomial nb.Coefficients
	N int

	// Modulus
	Modulus uint64

	// Unique factors of Modulus-1
	Factors []uint64

	// 2^bit_length(Modulus) - 1
	Mask uint64

	// Fast reduction constants
	BRedConstant [2]uint64 // Barrett Reduction
	MRedConstant uint64    // Montgomery Reduction

	*NTTTable // NTT related constants
}

// NewSubRing creates a new SubRing with the standard NTT.
// NTT constants still need to be generated using .GenNTTConstants(NthRoot uint64).
func NewSubRing(N int, Modulus uint64) (s *SubRing, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSubRingWithCustomNTT creates a new SubRing with degree N and modulus Modulus with user-defined NTT transform and primitive Nth root of unity.
// Modulus should be equal to 1 modulo the root of unity.
// N must be a power of two larger than 8. An error is returned with a nil *SubRing in the case of non NTT-enabling parameters.
func NewSubRingWithCustomNTT(N int, Modulus uint64, ntt func(*SubRing, int) NumberTheoreticTransformer, NthRoot int) (s *SubRing, err error) {
	_ = "STUB: not implemented"

	// Checks if N is a power of 2
	return nil, nil
}

/* #nosec G115 -- Modulus is ensured to be greater than 0 */

// Computes the fast modular reduction constants for the Ring

// If qi is not a power of 2, we can compute the MRed (otherwise, it
// would return an error as there is no valid Montgomery form mod a power of 2)

// Type returns the Type of subring which might be either `Standard` or `ConjugateInvariant`.
func (s *SubRing) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// Sanity check

// generateNTTConstants generates the NTT constant for the target SubRing.
// The fields `PrimitiveRoot` and `Factors` can be set manually to
// bypass the search for the primitive root (which requires to
// factor Modulus-1) and speedup the generation of the constants.
func (s *SubRing) generateNTTConstants() (err error) { _ = "STUB: not implemented"; return nil }

// Checks if each qi is prime and equal to 1 mod NthRoot

// It is possible to manually set the primitive root along with the factors of q-1.
// This is notably useful when marshalling the SubRing, to avoid re-factoring q-1.
// If both are set, then checks that that the root is indeed primitive.
// Else, factorize q-1 and finds a primitive root.

// 1.1 Computes N^(-1) mod Q in Montgomery form

// 1.2 Computes Psi and PsiInv in Montgomery form

// Computes Psi and PsiInv in Montgomery form

// Computes nttPsi[j] = nttPsi[j-1]*Psi and RootsBackward[j] = RootsBackward[j-1]*PsiInv

// PrimitiveRoot computes the smallest primitive root of the given prime q
// The unique factors of q-1 can be given to speed up the search for the root.
func PrimitiveRoot(q uint64, factors []uint64) (uint64, []uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

//Factor q-1, might be slow

// if for any factor of q-1, g^(q-1)/factor = 1 mod q, g is not a primitive root

// CheckFactors checks that the given list of factors contains
// all the unique primes of m.
func CheckFactors(m uint64, factors []uint64) (err error) { _ = "STUB: not implemented"; return nil }

// CheckPrimitiveRoot checks that g is a valid primitive root mod q,
// given the factors of q-1.
func CheckPrimitiveRoot(g, q uint64, factors []uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// subRingParametersLiteral is a struct to store the minimum information
// to uniquely identify a SubRing and be able to reconstruct it efficiently.
// This struct's purpose is to faciliate marshalling of SubRings.
type subRingParametersLiteral struct {
	Type          uint8    // Standard or ConjugateInvariant
	LogN          uint8    // Log2 of the ring degree
	NthRoot       uint8    // N/NthRoot
	Modulus       uint64   // Modulus
	Factors       []uint64 // Factors of Modulus-1
	PrimitiveRoot uint64   // Primitive root used
}

// ParametersLiteral returns the SubRingParametersLiteral of the SubRing.
func (s *SubRing) parametersLiteral() subRingParametersLiteral {
	_ = "STUB: not implemented"
	return *new(subRingParametersLiteral)
}

/* #nosec G115 -- s.Type has is 0 or 1 */

/* #nosec G115 -- N cannot be negative if SubRing is valid */

/* #nosec G115 -- NthRoot cannot be negative if SubRing is valid */

// newSubRingFromParametersLiteral creates a new SubRing from the provided subRingParametersLiteral.
func newSubRingFromParametersLiteral(p subRingParametersLiteral) (s *SubRing, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* #nosec G115 -- deserialization from valid subring -> N and NthRoot cannot be negative */

/* #nosec G115 -- Modulus cannot be negative */

// Computes the fast modular reduction parameters for the Ring

// If qi is not a power of 2, we can compute the MRed (otherwise, it
// would return an error as there is no valid Montgomery form mod a power of 2)

/* #nosec G115 -- library requires 64-bit system -> int = int64 */

/* #nosec G115 -- library requires 64-bit system -> int = int64 */

/* #nosec G115 -- library requires 64-bit system -> int = int64 */

/* #nosec G115 -- library requires 64-bit system -> int = int64 */
