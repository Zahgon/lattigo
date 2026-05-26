package sampling

import (
	"io"
	"sync"

	"golang.org/x/crypto/blake2b"
)

// PRNG is an interface for secure generation of random bytes
type PRNG interface {
	io.Reader
}

type ThreadSafePRNG struct {
}

// NewPRNG returns a new PRNG that is thread-safe
func NewPRNG() (*ThreadSafePRNG, error) { _ = "STUB: not implemented"; return nil, nil }

// Read reads bytes from the KeyedPRNG on sum.
func (prng *ThreadSafePRNG) Read(sum []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0,

		// KeyedPRNG is a structure storing the parameters used to securely and *deterministically* generate shared
		// sequences of random bytes among different parties using the hash function blake2b. Backward sequence
		// security (given the digest i, compute the digest i-1) is ensured by default, however forward sequence
		// security (given the digest i, compute the digest i+1) is only ensured if the KeyedPRNG is keyed.
		// WARNING: If KeyedPRNG is called concurrently by multiple threads, the resulting sequences will be independent and no error will be triggered. However, the result will not be deterministic and therefore, in most cases, it does not make sense to use KeyedPRNG in a concurrent setting.
		// NOTE: For a PRNG securely seeded with a private key use [ThreadSafePRNG].
		nil
}

type KeyedPRNG struct {
	mutex sync.Mutex
	key   []byte
	xof   blake2b.XOF
}

// NewKeyedPRNG creates a new instance of KeyedPRNG.
// Accepts an optional key, else set key=nil which is treated as key=[]byte{}
// WARNING: A PRNG INITIALISED WITH key=nil IS INSECURE!
// WARNING: KeyedPRNG can be called by multiple threads BUT the generated sequences will not be deterministic.
func NewKeyedPRNG(key []byte) (*KeyedPRNG, error) { _ = "STUB: not implemented"; return nil, nil }

// Key returns a copy of the key used to seed the PRNG.
// This value can be used with `NewKeyedPRNG` to instantiate
// a new PRNG that will produce the same stream of bytes.
func (prng *KeyedPRNG) Key() (key []byte) { _ = "STUB: not implemented"; return nil }

// Read reads bytes from the KeyedPRNG on sum.
// WARNING: Read() should NOT be called concurrently by multiple threads. If that occurs, the generated sequence will not be deterministic.
func (prng *KeyedPRNG) Read(sum []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Reset resets the PRNG to its initial state.
// WARNING: KeyedPRNG's methods should not be called concurrently. If that occurs, the generated sequence will not be deterministic.
func (prng *KeyedPRNG) Reset() { _ = "STUB: not implemented"; return }
