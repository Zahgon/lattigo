package mod1

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
)

// Type is the type of function/approximation used to evaluate x mod 1.
type Type uint64

// Sin and Cos are the two proposed functions for [Type].
// These trigonometric functions offer a good approximation of the function x mod 1 when the values are close to the origin.
const (
	CosDiscrete   = Type(0) // Special approximation (Han and Ki) of pow((1/2pi), 1/2^r) * cos(2pi(x-0.25)/2^r); this method requires a minimum degree of 2*(K-1).
	SinContinuous = Type(1) // Standard Chebyshev approximation of (1/2pi) * sin(2pix) on the full interval
	CosContinuous = Type(2) // Standard Chebyshev approximation of pow((1/2pi), 1/2^r) * cos(2pi(x-0.25)/2^r) on the full interval
)

// ParametersLiteral a struct for the parameters of the mod 1 procedure.
// The x mod 1 procedure goal is to homomorphically evaluate a modular reduction by Q[0] (the first prime of the moduli chain) on the encrypted plaintext.
// This struct is consumed by [NewParametersFromLiteral] to generate the [ParametersLiteral] struct, which notably stores
// the coefficient of the polynomial approximating the function x mod Q[0].
type ParametersLiteral struct {
	LevelQ          int     // Starting level of x mod 1
	LogScale        int     // Log2 of the scaling factor used during x mod 1
	Mod1Type        Type    // Chose between [Sin(2*pi*x)] or [cos(2*pi*x/r) with double angle formula]
	Scaling         float64 // Value by which the output is scaled by
	LogMessageRatio int     // Log2 of the ratio between Q0 and m, i.e. Q[0]/|m|
	K               int     // K parameter (interpolation in the range -K to K)
	Mod1Degree      int     // Degree of f: x mod 1
	DoubleAngle     int     // Number of rescale and double angle formula (only applies for cos and is ignored if sin is used)
	Mod1InvDegree   int     // Degree of f^-1: (x mod 1)^-1
}

// MarshalBinary returns a JSON representation of the the target Mod1ParametersLiteral struct on a slice of bytes.
// See Marshal from the [encoding/json] package.
func (evm ParametersLiteral) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalBinary reads a JSON representation on the target Mod1ParametersLiteral struct.
		// See Unmarshal from the [encoding/json] package.
		nil
}

func (evm *ParametersLiteral) UnmarshalBinary(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Depth returns the depth required to evaluate x mod 1.
func (evm ParametersLiteral) Depth() (depth int) { _ = "STUB: not implemented"; return 0 }

// this method requires a minimum degree of 2*K-1.
/* #nosec G115 -- Mod1Degree cannot be negative */

/* #nosec G115 -- Mod1Degree cannot be negative */

/* #nosec G115 -- Mod1InvDegree cannot be negative */

// Parameters is a struct storing the parameters and polynomials approximating the function x mod Q[0] (the first prime of the moduli chain).
type Parameters struct {
	LevelQ          int                // starting level of the operation
	LogDefaultScale int                // log2 of the default scaling factor
	Mod1Type        Type               // type of approximation for the f: x mod 1 function
	LogMessageRatio int                // Log2 of the ratio between Q0 and m, i.e. Q[0]/|m|
	DoubleAngle     int                // Number of rescale and double angle formula (only applies for cos and is ignored if sin is used)
	QDiff           float64            // Q / 2^round(Log2(Q))
	Sqrt2Pi         float64            // (1/2pi)^(1.0/scFac)
	Mod1Poly        bignum.Polynomial  // Polynomial for f: x mod 1
	Mod1InvPoly     *bignum.Polynomial // Polynomial for f^-1: (x mod 1)^-1
	K               float64            // interval [-K, K]
}

// IntervalShrinkFactor returns 2^{DoubleAngle}
func (evp Parameters) IntervalShrinkFactor() float64 { _ = "STUB: not implemented"; return 0 }

// KShrinked returns K / IntervalShrinkFactor()
func (evp Parameters) KShrinked() float64 { _ = "STUB: not implemented"; return 0 }

// ScalingFactor returns scaling factor used during the x mod 1.
func (evp Parameters) ScalingFactor() rlwe.Scale {
	_ = "STUB: not implemented"
	return *new(rlwe.Scale)
}

// MessageRatio returns the pre-set ratio Q[0]/|m|.
func (evp Parameters) MessageRatio() float64 { _ = "STUB: not implemented"; return 0 }

// NewParametersFromLiteral generates an [Parameters] struct from the [ParametersLiteral] struct.
// The [Parameters] struct is to instantiates a [Mod1Evaluator], which homomorphically evaluates x mod 1.
func NewParametersFromLiteral(params ckks.Parameters, evm ParametersLiteral) (Parameters, error) {
	_ = "STUB: not implemented"
	return *new(Parameters), nil
}

func sin2pi(x *big.Float) (y *big.Float) { _ = "STUB: not implemented"; return nil }

func cos2pi(x *big.Float) (y *big.Float) { _ = "STUB: not implemented"; return nil }
