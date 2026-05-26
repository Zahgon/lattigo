package ring

// IsPrime applies the Baillie-PSW, which is 100% accurate for numbers bellow 2^64.
func IsPrime(x uint64) bool { _ = "STUB: not implemented"; return false }

// NTTFriendlyPrimesGenerator is a struct used to generate NTT friendly primes.
type NTTFriendlyPrimesGenerator struct {
	Size                           float64
	NextPrime, PrevPrime, NthRoot  uint64
	CheckNextPrime, CheckPrevPrime bool
}

// NewNTTFriendlyPrimesGenerator instantiates a new NTTFriendlyPrimesGenerator.
// Primes generated are of the form 2^{BitSize} +/- k * {NthRoot} + 1.
func NewNTTFriendlyPrimesGenerator(BitSize, NthRoot uint64) NTTFriendlyPrimesGenerator {
	_ = "STUB: not implemented"
	return *new(NTTFriendlyPrimesGenerator)
}

// NextUpstreamPrimes returns the next k primes of the form 2^{BitSize} + k * {NthRoot} + 1.
func (n *NTTFriendlyPrimesGenerator) NextUpstreamPrimes(k int) (primes []uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextDownstreamPrimes returns the next k primes of the form 2^{BitSize} - k * {NthRoot} + 1.
func (n *NTTFriendlyPrimesGenerator) NextDownstreamPrimes(k int) (primes []uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextAlternatingPrimes returns the next k primes of the form 2^{BitSize} +/- k * {NthRoot} + 1.
func (n *NTTFriendlyPrimesGenerator) NextAlternatingPrimes(k int) (primes []uint64, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NextUpstreamPrime returns the next prime of the form 2^{BitSize} + k * {NthRoot} + 1.
func (n *NTTFriendlyPrimesGenerator) NextUpstreamPrime() (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Stops if the next prime would overlap with primes of the next bit-size or if an uint64 overflow would occur.

// NextDownstreamPrime returns the next prime of the form 2^{BitSize} - k * {NthRoot} + 1.
func (n *NTTFriendlyPrimesGenerator) NextDownstreamPrime() (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Stops if the next prime would overlap with the primes of the previous bit-size or if an uint64 overflow would occur.

// NextAlternatingPrime returns the next prime of the form 2^{BitSize} +/- k * {NthRoot} + 1.
func (n *NTTFriendlyPrimesGenerator) NextAlternatingPrime() (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Stops if the next prime would overlap with primes of the next bit-size or if an uint64 overflow would occure.

// Stops if the next prime would overlap with the primes of the previous bit-size or if an uint64 overflow would occure.
