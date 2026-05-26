package mpckks

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/multiparty"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// EncToShareProtocol is the structure storing the parameters and temporary buffers
// required by the encryption-to-shares protocol.
type EncToShareProtocol struct {
	multiparty.KeySwitchProtocol

	params ckks.Parameters
	zero   *rlwe.SecretKey
}

func NewAdditiveShare(params ckks.Parameters, logSlots int) multiparty.AdditiveShareBigint {
	_ = "STUB: not implemented"
	return *new(multiparty.AdditiveShareBigint)
}

// NewEncToShareProtocol creates a new EncToShareProtocol struct from the passed parameters.
func NewEncToShareProtocol(params ckks.Parameters, noise ring.DistributionParameters) (EncToShareProtocol, error) {
	_ = "STUB: not implemented"
	return *new(EncToShareProtocol), nil
}

// AllocateShare allocates a share of the EncToShare protocol
func (e2s EncToShareProtocol) AllocateShare(level int) (share multiparty.KeySwitchShare) {
	_ = "STUB: not implemented"
	return *new(multiparty.KeySwitchShare)
}

// GenShare generates a party's share in the encryption-to-shares protocol. This share consist in the additive secret-share of the party
// which is written in secretShareOut and in the public masked-decryption share written in publicShareOut.
// This protocol requires additional inputs which are:
//
//   - logBound : the bit length of the masks
//   - ct: the ciphertext to share
//
// publicShareOut is always returned in the NTT domain.
// The method [GetMinimumLevelForRefresh] should be used to get the minimum level at which EncToShare can be called while still ensure 128-bits of security, as well as the
// value for logBound.
func (e2s EncToShareProtocol) GenShare(sk *rlwe.SecretKey, logBound uint, ct *rlwe.Ciphertext, secretShareOut *multiparty.AdditiveShareBigint, publicShareOut *multiparty.KeySwitchShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Get the upperbound on the norm
// Ensures that bound >= 2^{128+logbound}

// Generate the mask in Z[Y] for Y = X^{N/(2*slots)}

// Encrypt the mask
// Generates an encryption of zero and subtracts the mask

// Positional -> RNS -> NTT

// Subtracts the mask to the encryption of zero

// GetShare is the final step of the encryption-to-share protocol. It performs the masked decryption of the target ciphertext followed by a
// the removal of the caller's secretShare as generated in the GenShare method.
// If the caller is not secret-key-share holder (i.e., didn't generate a decryption share), secretShare can be set to nil.
// Therefore, in order to obtain an additive sharing of the message, only one party should call this method, and the other parties should use
// the secretShareOut output of the GenShare method.
func (e2s EncToShareProtocol) GetShare(secretShare *multiparty.AdditiveShareBigint, aggregatePublicShare multiparty.KeySwitchShare, ct *rlwe.Ciphertext, secretShareOut *multiparty.AdditiveShareBigint) {
	_ = "STUB: not implemented"
	return
}

// Adds the decryption share on the ciphertext and stores the result in a buff

// INTT -> RNS -> Positional

// Subtracts the last mask

// ShareToEncProtocol is the structure storing the parameters and temporary buffers
// required by the shares-to-encryption protocol.
type ShareToEncProtocol struct {
	multiparty.KeySwitchProtocol
	params ckks.Parameters
	zero   *rlwe.SecretKey
}

// NewShareToEncProtocol creates a new ShareToEncProtocol struct from the passed parameters.
func NewShareToEncProtocol(params ckks.Parameters, noise ring.DistributionParameters) (ShareToEncProtocol, error) {
	_ = "STUB: not implemented"
	return *new(ShareToEncProtocol), nil
}

// AllocateShare allocates a share of the ShareToEnc protocol
func (s2e ShareToEncProtocol) AllocateShare(level int) (share multiparty.KeySwitchShare) {
	_ = "STUB: not implemented"
	return *new(multiparty.KeySwitchShare)
}

// GenShare generates a party's in the shares-to-encryption protocol given the party's secret-key share `sk`, a common
// polynomial sampled from the CRS `crs` and the party's secret share of the message.
func (s2e ShareToEncProtocol) GenShare(sk *rlwe.SecretKey, crs multiparty.KeySwitchCRP, metadata *rlwe.MetaData, secretShare multiparty.AdditiveShareBigint, c0ShareOut *multiparty.KeySwitchShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Generates an encryption share

// Positional -> RNS -> NTT

// GetEncryption computes the final encryption of the secret-shared message when provided with the aggregation `c0Agg` of the parties'
// share in the protocol and with the common, CRS-sampled polynomial `crs`.
func (s2e ShareToEncProtocol) GetEncryption(c0Agg multiparty.KeySwitchShare, crs multiparty.KeySwitchCRP, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}
