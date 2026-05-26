package multiparty

import (
	"io"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

// EvaluationKeyGenProtocol is the structure storing the parameters for the collective EvaluationKey generation.
type EvaluationKeyGenProtocol struct {
	params           rlwe.Parameters
	gaussianSamplerQ ring.Sampler
}

// NewEvaluationKeyGenProtocol creates a [EvaluationKeyGenProtocol] instance.
func NewEvaluationKeyGenProtocol(params rlwe.ParameterProvider) (evkg EvaluationKeyGenProtocol) {
	_ = "STUB: not implemented"
	return *new(EvaluationKeyGenProtocol)
}

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

// AllocateShare allocates a party's share in the EvaluationKey Generation.
func (evkg EvaluationKeyGenProtocol) AllocateShare(evkParams ...rlwe.EvaluationKeyParameters) EvaluationKeyGenShare {
	_ = "STUB: not implemented"
	return *new(EvaluationKeyGenShare)
}

func (evkg EvaluationKeyGenProtocol) allocateShare(levelQ, levelP, BaseTwoDecomposition int) EvaluationKeyGenShare {
	_ = "STUB: not implemented"
	return *new(EvaluationKeyGenShare)
}

// SampleCRP samples a common random polynomial to be used in the EvaluationKey Generation from the provided
// common reference string.
func (evkg EvaluationKeyGenProtocol) SampleCRP(crs CRS, evkParams ...rlwe.EvaluationKeyParameters) EvaluationKeyGenCRP {
	_ = "STUB: not implemented"
	return *new(EvaluationKeyGenCRP)
}

func (evkg EvaluationKeyGenProtocol) sampleCRP(crs CRS, levelQ, levelP, BaseTwoDecomposition int) EvaluationKeyGenCRP {
	_ = "STUB: not implemented"
	return *new(EvaluationKeyGenCRP)
}

// GenShare generates a party's share in the EvaluationKey Generation.
func (evkg EvaluationKeyGenProtocol) GenShare(skIn, skOut *rlwe.SecretKey, crp EvaluationKeyGenCRP, shareOut *EvaluationKeyGenShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// e

// a is the CRP

// e + sk_in * (qiBarre*qiStar) * 2^w
// (qiBarre*qiStar)%qi = 1, else 0

// Handles the case where nb pj does not divides nb qi

// sk_in * (qiBarre*qiStar) * 2^w - a*sk + e

// AggregateShares computes share3 = share1 + share2.
func (evkg EvaluationKeyGenProtocol) AggregateShares(share1, share2 EvaluationKeyGenShare, share3 *EvaluationKeyGenShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GenEvaluationKey finalizes the EvaluationKey Generation and populates the input Evaluationkey with the computed collective EvaluationKey.
func (evkg EvaluationKeyGenProtocol) GenEvaluationKey(share EvaluationKeyGenShare, crp EvaluationKeyGenCRP, evk *rlwe.EvaluationKey) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EvaluationKeyGenCRP is a type for common reference polynomials in the EvaluationKey Generation protocol.
type EvaluationKeyGenCRP struct {
	Value structs.Matrix[ringqp.Poly]
}

// LevelQ returns the level of the ciphertext modulus of the target share.
func (crp EvaluationKeyGenCRP) LevelQ() int { _ = "STUB: not implemented"; return 0 }

// LevelP returns the level of the auxiliary switching key modulus of the target share.
func (crp EvaluationKeyGenCRP) LevelP() int { _ = "STUB: not implemented"; return 0 }

// BaseTwoDecompositionVectorSize returns the number of element in the Power of two decomposition basis for each prime of Q.
func (crp EvaluationKeyGenCRP) BaseTwoDecompositionVectorSize() (base []int) {
	_ = "STUB: not implemented"
	return nil
}

// BaseRNSDecompositionVectorSize returns the number of element in the RNS decomposition basis: Ceil(lenQi / lenPi)
func (crp EvaluationKeyGenCRP) BaseRNSDecompositionVectorSize() int {
	_ = "STUB: not implemented"
	return 0

	// EvaluationKeyGenShare is represent a Party's share in the EvaluationKey Generation protocol.
}

type EvaluationKeyGenShare struct {
	rlwe.GadgetCiphertext
}

// BinarySize returns the serialized size of the object in bytes.
func (share EvaluationKeyGenShare) BinarySize() int { _ = "STUB: not implemented"; return 0 }

// WriteTo writes the object on an [io.Writer]. It implements the [io.WriterTo]
// interface, and will write exactly object.BinarySize() bytes on w.
//
// Unless w implements the [buffer.Writer] interface (see lattigo/utils/buffer/writer.go),
// it will be wrapped into a [bufio.Writer]. Since this requires allocations, it
// is preferable to pass a [buffer.Writer] directly:
//
//   - When writing multiple times to a io.Writer, it is preferable to first wrap the
//     io.Writer in a pre-allocated [bufio.Writer].
//   - When writing to a pre-allocated var b []byte, it is preferable to pass
//     buffer.NewBuffer(b) as w (see lattigo/utils/buffer/buffer.go).
func (share EvaluationKeyGenShare) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom reads on the object from an [io.Writer]. It implements the
// [io.ReaderFrom] interface.
//
// Unless r implements the [buffer.Reader] interface (see see lattigo/utils/buffer/reader.go),
// it will be wrapped into a [bufio.Reader]. Since this requires allocation, it
// is preferable to pass a [buffer.Reader] directly:
//
//   - When reading multiple values from a io.Reader, it is preferable to first
//     first wrap [io.Reader] in a pre-allocated bufio.Reader.
//   - When reading from a var b []byte, it is preferable to pass a buffer.NewBuffer(b)
//     as w (see lattigo/utils/buffer/buffer.go).
func (share *EvaluationKeyGenShare) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MarshalBinary encodes the object into a binary form on a newly allocated slice of bytes.
func (share EvaluationKeyGenShare) MarshalBinary() (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalBinary decodes a slice of bytes generated by
// [EvaluationKeyGenShare.MarshalBinary] or [EvaluationKeyGenShare.WriteTo] on the object.
func (share *EvaluationKeyGenShare) UnmarshalBinary(p []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
