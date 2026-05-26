package bootstrapping

import (
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/dft"
	"github.com/tuneinsight/lattigo/v6/circuits/ckks/mod1"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// Parameters is a struct storing the parameters
// of the bootstrapping circuit.
type Parameters struct {
	// ResidualParameters: Parameters outside of the bootstrapping circuit
	ResidualParameters ckks.Parameters
	// BootstrappingParameters: Parameters during the bootstrapping circuit
	BootstrappingParameters ckks.Parameters
	// SlotsToCoeffsParameters Parameters of the homomorphic decoding linear transformation
	SlotsToCoeffsParameters dft.MatrixLiteral
	// Mod1ParametersLiteral: Parameters of the homomorphic modular reduction
	Mod1ParametersLiteral mod1.ParametersLiteral
	// CoeffsToSlotsParameters: Parameters of the homomorphic encoding linear transformation
	CoeffsToSlotsParameters dft.MatrixLiteral
	// IterationsParameters: Parameters of the bootstrapping iterations (META-BTS)
	IterationsParameters *IterationsParameters
	// EphemeralSecretWeight: Hamming weight of the ephemeral secret. If 0, no ephemeral secret is used during the bootstrapping.
	EphemeralSecretWeight int
	// CircuitOrder: Value indicating the order of the circuit (default: ModUpThenEncode)
	CircuitOrder CircuitOrder
}

// NewParametersFromLiteral instantiates a [Parameters] from the residual [ckks.Parameters] and
// a [ParametersLiteral] struct.
//
// The residualParameters corresponds to the [ckks.Parameters] that are left after the bootstrapping circuit is evaluated.
// These are entirely independent of the bootstrapping parameters with one exception: the ciphertext primes Qi must be
// congruent to 1 mod 2N of the bootstrapping parameters (note that the auxiliary primes Pi do not need to be).
// This is required because the primes Qi of the residual parameters and the bootstrapping parameters are the same between
// the two sets of parameters.
//
// The user can ensure that this condition is met by setting the appropriate LogNThRoot in the [ckks.ParametersLiteral] before
// instantiating them.
//
// The method NewParametersFromLiteral will automatically allocate the [ckks.Parameters] of the bootstrapping circuit based on
// the provided residualParameters and the information given in the [ParametersLiteral].
func NewParametersFromLiteral(residualParameters ckks.Parameters, btpLit ParametersLiteral) (Parameters, error) {
	_ = "STUB: not implemented"

	// Retrieve the LogN of the bootstrapping circuit
	return *new(Parameters), nil
}

// Retrieve the NthRoot

// If ConjugateInvariant, then the bootstrapping LogN must be at least 1 greater
// than the residualParameters LogN

// Takes the greatest NthRoot between the residualParameters NthRoot and the bootstrapping NthRoot
/* #nosec G115 -- N cannot be negative */

// The LogN of the bootstrapping parameters cannot be smaller than the LogN of the residualParameters.

// Takes the greatest NthRoot between the residualParameters NthRoot and the bootstrapping NthRoot
/* #nosec G115 -- N cannot be negative */

// Checks that all primes Qi of the residualParameters are congruent to 1 mod NthRoot of the bootstrapping parameters.

// Retrieves the variable LogSlots, which is used to instantiates the encoding/decoding matrices.

// Retrieves the factorization depth and scaling factor of the encoding matrix

// Retrieves the factorization depth and scaling factor of the decoding matrix

// Slots To Coeffs params

// Number of bootstrapping iterations

// Boolean if there is a reserved prime for the bootstrapping iterations

// SlotsToCoeffs parameters (homomorphic decoding)

// Scaling factor of the homomorphic modular reduction x mod 1

// Type of polynomial approximation of x mod 1

// Degree of the taylor series of arc sine

// Log2 ratio between Q[0] and |m| (i.e. gap between the message and Q[0])

// Interval [-K+1, K-1] of the polynomial approximation of x mod 1

// Number of double angle evaluation if x mod 1 is approximated with cos

// Degree of the polynomial approximation of x mod 1

// Parameters of the homomorphic modular reduction x mod 1

// Hamming weight of the ephemeral secret key to which the ciphertext is
// switched to during the ModUp step.

// Coeffs To Slots params

// Parameters of the CoeffsToSlots (homomorphic encoding)

// List of the prime-size of all primes required by the bootstrapping circuit.

// appends the reserved prime first for multiple iteration, if any

// Appends all other primes in reverse order of the circuit

// Extracts all the different primes Qi that are
// in the residualParameters

// Maps the number of primes per bit size

// Retrieve the number of primes #Pi of the bootstrapping circuit
// and adds them to the list of bit-size

// Map to store [bit-size][]primes

// For each bit-size sample a pair-wise coprime prime

// Creates a new prime generator
/* #nosec G115 -- logqi cannot be negative */

// Populates the list with primes that aren't yet in primesHave

// Constructs the set of primes Qi

// Appends to the residual moduli

// Constructs the set of primes Pi

// Ensure that ckks.PrecisionMode = PREC64 when using PREC128 residual parameters.

// Instantiates the ckks.Parameters of the bootstrapping circuit.

func (p Parameters) Equal(other *Parameters) (res bool) { _ = "STUB: not implemented"; return false }

// LogMaxDimensions returns the log plaintext dimensions of the target Parameters.
func (p Parameters) LogMaxDimensions() ring.Dimensions {
	_ = "STUB: not implemented"
	return *new(ring.Dimensions)
}

// LogMaxSlots returns the log of the maximum number of slots.
func (p Parameters) LogMaxSlots() int { _ = "STUB: not implemented"; return 0 }

// DepthCoeffsToSlots returns the depth of the Coeffs to Slots of the bootstrapping.
func (p Parameters) DepthCoeffsToSlots() (depth int) { _ = "STUB: not implemented"; return 0 }

// DepthEvalMod returns the depth of the EvalMod step of the bootstrapping.
func (p Parameters) DepthEvalMod() (depth int) { _ = "STUB: not implemented"; return 0 }

// DepthSlotsToCoeffs returns the depth of the Slots to Coeffs step of the bootstrapping.
func (p Parameters) DepthSlotsToCoeffs() (depth int) { _ = "STUB: not implemented"; return 0 }

// Depth returns the depth of the full bootstrapping circuit.
func (p Parameters) Depth() (depth int) { _ = "STUB: not implemented"; return 0 }

// MarshalBinary returns a JSON representation of the Parameters struct.
// See Marshal from the [encoding/json] package.
func (p Parameters) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalBinary reads a JSON representation on the target Parameters struct.
		// See Unmarshal from the [encoding/json] package.
		nil
}

func (p *Parameters) UnmarshalBinary(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p Parameters) MarshalJSON() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parameters) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

// GaloisElements returns the list of Galois elements required to evaluate the bootstrapping.
func (p Parameters) GaloisElements(params ckks.Parameters) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil

	// List of the rotation key values to needed for the bootstrap
}

//SubSum rotation needed X -> Y^slots rotations
