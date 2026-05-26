package multiparty

import (
	"io"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

// Thresholdizer is a type for generating secret-shares of [ringqp.Poly] types such that
// the resulting sharing has a t-out-of-N-threshold access-structure. It implements the
// "Thresholdize" operation as presented in "An Efficient Threshold Access-Structure
// for RLWE-Based Multiparty Homomorphic Encryption" (2022) by Mouchet, C., Bertrand, E.,
// and Hubaux, J. P. (https://eprint.iacr.org/2022/780).
//
// See the [multiparty] package README.md.
type Thresholdizer struct {
	params   *rlwe.Parameters
	ringQP   *ringqp.Ring
	usampler ringqp.UniformSampler
}

// Combiner is a type for generating t-out-of-t additive shares from local t-out-of-N
// shares. It implements the "Combine" operation as presented in "An Efficient Threshold
// Access-Structure for RLWE-Based Multiparty Homomorphic Encryption" (2022) by Mouchet, C.,
// Bertrand, E., and Hubaux, J. P. (https://eprint.iacr.org/2022/780).
type Combiner struct {
	ringQP         *ringqp.Ring
	threshold      int
	tmp1, tmp2     []uint64
	one            ring.RNSScalar
	lagrangeCoeffs map[ShamirPublicPoint]ring.RNSScalar
}

// ShamirPublicPoint is a type for Shamir public point associated with a party identity within
// the t-out-of-N-threshold scheme.
//
// See [Thresholdizer] and [Combiner] types.
type ShamirPublicPoint uint64

// ShamirPolynomial represents a polynomial with [ringqp.Poly] coefficients. It is used by the
// Thresholdizer type to produce t-out-of-N-threshold shares of an [ringqp.Poly].
//
// See [Thresholdizer] type.
type ShamirPolynomial struct {
	Value structs.Vector[ringqp.Poly]
}

// ShamirSecretShare represents a t-out-of-N-threshold secret-share.
//
// See [Thresholdizer] and [Combiner] types.
type ShamirSecretShare struct {
	ringqp.Poly
}

// NewThresholdizer creates a new [Thresholdizer] instance from parameters.
func NewThresholdizer(params rlwe.ParameterProvider) Thresholdizer {
	_ = "STUB: not implemented"
	return *new(Thresholdizer)
}

// Sanity check, this error should not happen.

// GenShamirPolynomial generates a new secret [ShamirPolynomial] to be used in the [Thresholdizer.GenShamirSecretShare] method.
// It does so by sampling a random polynomial of degree threshold - 1 and with its constant term equal to secret.
func (thr Thresholdizer) GenShamirPolynomial(threshold int, secret *rlwe.SecretKey) (ShamirPolynomial, error) {
	_ = "STUB: not implemented"
	return *new(ShamirPolynomial), nil
}

// AllocateThresholdSecretShare allocates a [ShamirSecretShare] struct.
func (thr Thresholdizer) AllocateThresholdSecretShare() ShamirSecretShare {
	_ = "STUB: not implemented"
	return *new(ShamirSecretShare)
}

// GenShamirSecretShare generates a secret share for the given recipient, identified by its [ShamirPublicPoint].
// The result is stored in ShareOut and should be sent to this party.
func (thr Thresholdizer) GenShamirSecretShare(recipient ShamirPublicPoint, secretPoly ShamirPolynomial, shareOut *ShamirSecretShare) {
	_ = "STUB: not implemented"
	return
}

// AggregateShares aggregates two [ShamirSecretShare] and stores the result in outShare.
func (thr Thresholdizer) AggregateShares(share1, share2 ShamirSecretShare, outShare *ShamirSecretShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// NewCombiner creates a new [Combiner] struct from the parameters and the set of [ShamirPublicPoints]. Note that the other
// parameter may contain the instantiator's own [ShamirPublicPoint].
func NewCombiner(params rlwe.ParameterProvider, own ShamirPublicPoint, others []ShamirPublicPoint, threshold int) Combiner {
	_ = "STUB: not implemented"
	return *new(Combiner)
}

// precomputes lagrange coefficient factors

// GenAdditiveShare generates a t-out-of-t additive share of the secret from a local aggregated share ownSecret and the set of active identities, identified
// by their [ShamirPublicPoint]. It stores the resulting additive share in skOut.
func (cmb Combiner) GenAdditiveShare(activesPoints []ShamirPublicPoint, ownPoint ShamirPublicPoint, ownShare ShamirSecretShare, skOut *rlwe.SecretKey) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//Lagrange Interpolation with the public threshold key of other active players

func (cmb Combiner) lagrangeCoeff(thisKey ShamirPublicPoint, thatKey ShamirPublicPoint, lagCoeff []uint64) {
	_ = "STUB: not implemented"
	return
}

// BinarySize returns the serialized size of the object in bytes.
func (s ShamirSecretShare) BinarySize() int { _ = "STUB: not implemented"; return 0 }

// WriteTo writes the object on an [io.Writer]. It implements the [io.WriterTo]
// interface, and will write exactly object.BinarySize() bytes on w.
//
// Unless w implements the [buffer.Writer] interface (see lattigo/utils/buffer/writer.go),
// it will be wrapped into a [bufio.Writer]. Since this requires allocations, it
// is preferable to pass a [buffer.Writer] directly:
//
//   - When writing multiple times to a [io.Writer], it is preferable to first wrap the
//     [io.Writer] in a pre-allocated [bufio.Writer].
//   - When writing to a pre-allocated var b []byte, it is preferable to pass
//     buffer.NewBuffer(b) as w (see lattigo/utils/buffer/buffer.go).
func (s ShamirSecretShare) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0,

		// ReadFrom reads on the object from an [io.Writer]. It implements the
		// [io.ReaderFrom] interface.
		//
		// Unless r implements the [buffer.Reader] interface (see see lattigo/utils/buffer/reader.go),
		// it will be wrapped into a [bufio.Reader]. Since this requires allocation, it
		// is preferable to pass a [buffer.Reader] directly:
		//
		//   - When reading multiple values from a [io.Reader], it is preferable to first
		//     first wrap [io.Reader] in a pre-allocated [bufio.Reader].
		//   - When reading from a var b []byte, it is preferable to pass a buffer.NewBuffer(b)
		//     as w (see lattigo/utils/buffer/buffer.go).
		nil
}

func (s *ShamirSecretShare) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0,

		// MarshalBinary encodes the object into a binary form on a newly allocated slice of bytes.
		nil
}

func (s ShamirSecretShare) MarshalBinary() (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil

	// UnmarshalBinary decodes a slice of bytes generated by
	// [ShamirSecretShare.MarshalBinary] or [ShamirSecretShare.WriteTo] on the object.
}

func (s *ShamirSecretShare) UnmarshalBinary(p []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
