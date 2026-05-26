package mpbgv

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/multiparty"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// MaskedTransformProtocol is a struct storing the parameters for the [MaskedTransformProtocol] protocol.
type MaskedTransformProtocol struct {
	e2s EncToShareProtocol
	s2e ShareToEncProtocol
}

// MaskedTransformFunc is a struct containing a user-defined in-place function that can be applied to masked integer plaintexts, as a part of the
// Masked Transform Protocol.
// The function is called with a vector of integers modulo bgv.Parameters.PlaintextModulus() of size bgv.Parameters.N() as input, and must write
// its output on the same buffer.
// Transform can be the identity.
//
//   - Decode: if true, then the masked BFV plaintext will be decoded before applying Transform.
//   - Recode: if true, then the masked BFV plaintext will be recoded after applying Transform.
//
// Decode (true/false) -> Transform -> Recode (true/false).
type MaskedTransformFunc struct {
	Decode bool
	Func   func(coeffs []uint64)
	Encode bool
}

// NewMaskedTransformProtocol creates a new instance of the PermuteProtocol.
func NewMaskedTransformProtocol(paramsIn, paramsOut bgv.Parameters, noiseFlooding ring.DistributionParameters) (rfp MaskedTransformProtocol, err error) {
	_ = "STUB: not implemented"
	return *new(MaskedTransformProtocol), nil
}

// SampleCRP samples a common random polynomial to be used in the Masked-Transform protocol from the provided
// common reference string.
func (rfp *MaskedTransformProtocol) SampleCRP(level int, crs sampling.PRNG) multiparty.KeySwitchCRP {
	_ = "STUB: not implemented"
	return *new(multiparty.KeySwitchCRP)
}

// AllocateShare allocates the shares of the PermuteProtocol
func (rfp MaskedTransformProtocol) AllocateShare(levelDecrypt, levelRecrypt int) multiparty.RefreshShare {
	_ = "STUB: not implemented"
	return *new(multiparty.RefreshShare)
}

// GenShare generates the shares of the PermuteProtocol.
// ct1 is the degree 1 element of a rlwe.Ciphertext, i.e. rlwe.Ciphertext.Value[1].
func (rfp MaskedTransformProtocol) GenShare(skIn, skOut *rlwe.SecretKey, ct *rlwe.Ciphertext, crs multiparty.KeySwitchCRP, transform *MaskedTransformFunc, shareOut *multiparty.RefreshShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Stores the ciphertext metadata on the public share

// AggregateShares sums share1 and share2 on shareOut.
func (rfp MaskedTransformProtocol) AggregateShares(share1, share2 multiparty.RefreshShare, shareOut *multiparty.RefreshShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Transform applies Decrypt, Recode and Recrypt on the input ciphertext.
func (rfp MaskedTransformProtocol) Transform(ct *rlwe.Ciphertext, transform *MaskedTransformFunc, crs multiparty.KeySwitchCRP, share multiparty.RefreshShare, ciphertextOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// tmpMask RingT(m - sum M_i)
