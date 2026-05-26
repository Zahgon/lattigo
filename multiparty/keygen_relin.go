package multiparty

import (
	"io"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/ring/ringqp"
	"github.com/tuneinsight/lattigo/v6/utils/structs"
)

// RelinearizationKeyGenProtocol is the structure storing the parameters and and precomputations for the collective relinearization key generation protocol.
type RelinearizationKeyGenProtocol struct {
	params rlwe.Parameters

	gaussianSamplerQ ring.Sampler
	ternarySamplerQ  ring.Sampler
}

// RelinearizationKeyGenShare is a share in the RelinearizationKeyGen protocol.
type RelinearizationKeyGenShare struct {
	rlwe.GadgetCiphertext
}

// RelinearizationKeyGenCRP is a type for common reference polynomials in the RelinearizationKeyGen protocol.
type RelinearizationKeyGenCRP struct {
	Value structs.Matrix[ringqp.Poly]
}

// NewRelinearizationKeyGenProtocol creates a new RelinearizationKeyGen protocol struct.
func NewRelinearizationKeyGenProtocol(params rlwe.ParameterProvider) RelinearizationKeyGenProtocol {
	_ = "STUB: not implemented"
	return *new(RelinearizationKeyGenProtocol)
}

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

// Sanity check, this error should not happen.

// SampleCRP samples a common random polynomial to be used in the RelinearizationKeyGen protocol from the provided
// common reference string.
func (ekg RelinearizationKeyGenProtocol) SampleCRP(crs CRS, evkParams ...rlwe.EvaluationKeyParameters) RelinearizationKeyGenCRP {
	_ = "STUB: not implemented"
	return *new(RelinearizationKeyGenCRP)
}

// GenShareRoundOne is the first of three rounds of the [RelinearizationKeyGenProtocol] protocol. Each party generates a pseudo encryption of
// its secret share of the key s_i under its ephemeral key u_i : [-u_i*a + s_i*w + e_i] and broadcasts it to the other
// j-1 parties.
//
// round1 = [-u_i * a + s_i * P + e_0i, s_i* a + e_i1]
func (ekg RelinearizationKeyGenProtocol) GenShareRoundOne(sk *rlwe.SecretKey, crp RelinearizationKeyGenCRP, ephSkOut *rlwe.SecretKey, shareOut *RelinearizationKeyGenShare) {
	_ = "STUB: not implemented"
	// Given a base decomposition w_i (here the CRT decomposition)
	// computes [-u*a_i + P*s_i + e_i, s_i * a + e_i]
	// where a_i = crp_i
	return
}

// Computes P * sk

// u

// h = e

// h = sk*CrtBaseDecompQi + e

// Handles the case where nb pj does not divides nb qi

// h = sk*CrtBaseDecompQi + -u*a + e

// Second Element
// e_2i

// s*a + e_2i

// GenShareRoundTwo is the second of three rounds of the [RelinearizationKeyGenProtocol] protocol. Upon receiving the j-1 shares, each party computes :
//
//   - round1 = sum([-u_i * a + s_i * P + e_0i, s_i* a + e_i1]) = [-ua + sP + e0, sa + e1]
//
//   - round2 = [s_i * round1[0] + (u_i - s_i) * round1[1] + e_i2] = [s_i * {-ua + s * P + e0} + (u_i - s_i) * {sa + e1} + e_i2]
//
// and broadcasts both values to the other j-1 parties.
func (ekg RelinearizationKeyGenProtocol) GenShareRoundTwo(ephSk, sk *rlwe.SecretKey, round1 RelinearizationKeyGenShare, shareOut *RelinearizationKeyGenShare) {
	_ = "STUB: not implemented"
	return
}

// (u_i - s_i)

// Each sample is of the form [-u*a_i + s*w_i + e_i]
// So for each element of the base decomposition w_i:

// Computes [(sum samples)*sk + e_1i, sk*a + e_2i]

// (AggregateShareRoundTwo samples) * sk

// (AggregateShareRoundTwo samples) * sk + e_1i

// second part
// (AggRound1Samples[0])*sk + (u_i - s_i) * (AggRound1Samples[1]) + e_1

// AggregateShares combines two RelinearizationKeyGen shares into a single one.
func (ekg RelinearizationKeyGenProtocol) AggregateShares(share1, share2 RelinearizationKeyGenShare, shareOut *RelinearizationKeyGenShare) {
	_ = "STUB: not implemented"
	return
}

// deg(round 1 shares) = 1, deg(round 2 shares) = 0

// GenRelinearizationKey computes the generated RLK from the public shares and write the result in evalKeyOut.
//
//   - round1 = [-ua + sP + e0, sa + e1]
//   - round2 = sum([s_i * {-ua + sP + e0} + (u_i - s_i) * {sa + e1} + e_i2]) = [-sua + Ps^2 + se0 + e2, sua + ue1 - s^2a -se1]
//   - [round2[0] + round2[1], round1[1]] = [-{s^2a + se1} + Ps^2 + {se0 + ue1 + e2}, sa + e1] = [sb + Ps^2 + e, b]
func (ekg RelinearizationKeyGenProtocol) GenRelinearizationKey(round1 RelinearizationKeyGenShare, round2 RelinearizationKeyGenShare, evalKeyOut *rlwe.RelinearizationKey) {
	_ = "STUB: not implemented"
	return
}

// AllocateShare allocates the share of the EKG protocol.
// To satisfy the correctness of the multi-party protocol, linearization keys shares cannot be allocated in the compressed format.
func (ekg RelinearizationKeyGenProtocol) AllocateShare(evkParams ...rlwe.EvaluationKeyParameters) (ephSk *rlwe.SecretKey, r1 RelinearizationKeyGenShare, r2 RelinearizationKeyGenShare) {
	_ = "STUB: not implemented"
	return nil, *new(RelinearizationKeyGenShare), *new(RelinearizationKeyGenShare)
}

// BinarySize returns the serialized size of the object in bytes.
func (share RelinearizationKeyGenShare) BinarySize() int { _ = "STUB: not implemented"; return 0 }

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
func (share RelinearizationKeyGenShare) WriteTo(w io.Writer) (n int64, err error) {
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
func (share *RelinearizationKeyGenShare) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MarshalBinary encodes the object into a binary form on a newly allocated slice of bytes.
func (share RelinearizationKeyGenShare) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalBinary decodes a slice of bytes generated by
// [RelinearizationKeyGenShare.MarshalBinary] or [RelinearizationKeyGenShare.WriteTo] on the object.
func (share *RelinearizationKeyGenShare) UnmarshalBinary(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
