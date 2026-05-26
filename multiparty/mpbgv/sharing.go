package mpbgv

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/multiparty"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
)

// EncToShareProtocol is the structure storing the parameters and temporary buffers
// required by the encryption-to-shares protocol.
type EncToShareProtocol struct {
	multiparty.KeySwitchProtocol
	params bgv.Parameters

	maskSampler *ring.UniformSampler
	encoder     *bgv.Encoder

	zero *rlwe.SecretKey
}

func NewAdditiveShare(params bgv.Parameters) multiparty.AdditiveShare {
	_ = "STUB: not implemented"
	return *new(multiparty.AdditiveShare)
}

// NewEncToShareProtocol creates a new EncToShareProtocol struct from the passed bgv.Parameters.
func NewEncToShareProtocol(params bgv.Parameters, noiseFlooding ring.DistributionParameters) (EncToShareProtocol, error) {
	_ = "STUB: not implemented"
	return *new(EncToShareProtocol), nil
}

// Sanity check, this error should not happen.

// AllocateShare allocates a share of the EncToShare protocol
func (e2s EncToShareProtocol) AllocateShare(level int) (share multiparty.KeySwitchShare) {
	_ = "STUB: not implemented"
	return *new(multiparty.KeySwitchShare)
}

// GenShare generates a party's share in the encryption-to-shares protocol. This share consist in the additive secret-share of the party
// which is written in secretShareOut and in the public masked-decryption share written in publicShareOut.
// ct1 is degree 1 element of a rlwe.Ciphertext, i.e. rlwe.Ciphertext.Value[1].
func (e2s EncToShareProtocol) GenShare(sk *rlwe.SecretKey, ct *rlwe.Ciphertext, secretShareOut *multiparty.AdditiveShare, publicShareOut *multiparty.KeySwitchShare) {
	_ = "STUB: not implemented"
	return
}

// GetShare is the final step of the encryption-to-share protocol. It performs the masked decryption of the target ciphertext followed by a
// the removal of the caller's secretShare as generated in the [EncToShareProtocol.GenShare] method.
// If the caller is not secret-key-share holder (i.e., didn't generate a decryption share), secretShare can be set to nil.
// Therefore, in order to obtain an additive sharing of the message, only one party should call this method, and the other parties should use
// the secretShareOut output of the GenShare method.
func (e2s EncToShareProtocol) GetShare(secretShare *multiparty.AdditiveShare, aggregatePublicShare multiparty.KeySwitchShare, ct *rlwe.Ciphertext, secretShareOut *multiparty.AdditiveShare) {
	_ = "STUB: not implemented"
	return
}

// ShareToEncProtocol is the structure storing the parameters and temporary buffers
// required by the shares-to-encryption protocol.
type ShareToEncProtocol struct {
	multiparty.KeySwitchProtocol
	params bgv.Parameters

	encoder *bgv.Encoder

	zero *rlwe.SecretKey
}

// NewShareToEncProtocol creates a new ShareToEncProtocol struct from the passed integer parameters.
func NewShareToEncProtocol(params bgv.Parameters, noiseFlooding ring.DistributionParameters) (ShareToEncProtocol, error) {
	_ = "STUB: not implemented"
	return *new(ShareToEncProtocol), nil
}

// AllocateShare allocates a share of the ShareToEnc protocol
func (s2e ShareToEncProtocol) AllocateShare(level int) (share multiparty.KeySwitchShare) {
	_ = "STUB: not implemented"
	return *new(multiparty.KeySwitchShare)
}

// GenShare generates a party's in the shares-to-encryption protocol given the party's secret-key share `sk`, a common
// polynomial sampled from the CRS `crp` and the party's secret share of the message.
func (s2e ShareToEncProtocol) GenShare(sk *rlwe.SecretKey, crp multiparty.KeySwitchCRP, secretShare multiparty.AdditiveShare, c0ShareOut *multiparty.KeySwitchShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GetEncryption computes the final encryption of the secret-shared message when provided with the aggregation `c0Agg` of the parties'
// shares in the protocol and with the common, CRS-sampled polynomial `crp`.
func (s2e ShareToEncProtocol) GetEncryption(c0Agg multiparty.KeySwitchShare, crp multiparty.KeySwitchCRP, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}
