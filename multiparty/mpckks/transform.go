package mpckks

import (
	"math/big"

	"github.com/tuneinsight/lattigo/v6/multiparty"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// MaskedLinearTransformationProtocol is a struct storing the parameters for the [MaskedLinearTransformationProtocol] protocol.
type MaskedLinearTransformationProtocol struct {
	e2s EncToShareProtocol
	s2e ShareToEncProtocol

	noise ring.DistributionParameters

	defaultScale *big.Int
	prec         uint

	encoder *ckks.Encoder
}

// WithParams creates a shallow copy of the target [MaskedLinearTransformationProtocol] but with new output parameters.
// The expected input parameters remain unchanged.
func (mltp MaskedLinearTransformationProtocol) WithParams(paramsOut ckks.Parameters) MaskedLinearTransformationProtocol {
	_ = "STUB: not implemented"
	return *new(MaskedLinearTransformationProtocol)
}

// Sanity check, this error should not happen.

// MaskedLinearTransformationFunc represents a user-defined in-place function that can be evaluated on masked float plaintexts, as a part of the
// Masked Transform Protocol.
// The function is called with a vector of *Complex modulo ckks.Parameters.Slots() as input, and must write
// its output on the same buffer.
// Transform can be the identity.
//
//   - Decode: if true, then the masked float plaintext will be decoded before applying Transform.
//   - Recode: if true, then the masked float plaintext will be recoded after applying Transform.
//
// Decode (true/false) -> Transform -> Recode (true/false).
type MaskedLinearTransformationFunc struct {
	Decode bool
	Func   func(coeffs []*bignum.Complex)
	Encode bool
}

// NewMaskedLinearTransformationProtocol creates a new instance of the PermuteProtocol.
// paramsIn: the ckks.Parameters of the ciphertext before the protocol.
// paramsOut: the ckks.Parameters of the ciphertext after the protocol.
// prec : the log2 of decimal precision of the internal encoder.
// The method will return an error if the maximum number of slots of the output parameters is smaller than the number of slots of the input ciphertext.
func NewMaskedLinearTransformationProtocol(paramsIn, paramsOut ckks.Parameters, prec uint, noise ring.DistributionParameters) (mltp MaskedLinearTransformationProtocol, err error) {
	_ = "STUB: not implemented"
	return *new(MaskedLinearTransformationProtocol), nil
}

// AllocateShare allocates the shares of the PermuteProtocol
func (mltp MaskedLinearTransformationProtocol) AllocateShare(levelDecrypt, levelRecrypt int) multiparty.RefreshShare {
	_ = "STUB: not implemented"
	return *new(multiparty.RefreshShare)
}

// SampleCRP samples a common random polynomial to be used in the Masked-Transform protocol from the provided
// common reference string. The CRP is considered to be in the NTT domain.
func (mltp MaskedLinearTransformationProtocol) SampleCRP(level int, crs sampling.PRNG) multiparty.KeySwitchCRP {
	_ = "STUB: not implemented"
	return *new(multiparty.KeySwitchCRP)
}

// GenShare generates the shares of the PermuteProtocol
// This protocol requires additional inputs which are:
//
//   - skIn     : the secret-key if the input ciphertext.
//   - skOut    : the secret-key of the output ciphertext.
//   - logBound : the bit length of the masks.
//   - ct1      : the degree 1 element the ciphertext to refresh, i.e. ct1 = ckk.Ciphetext.Value[1].
//   - scale    : the scale of the ciphertext when entering the refresh.
//
// The method [GetMinimumLevelForRefresh] should be used to get the minimum level at which the masked transform can be called while still ensure 128-bits of security, as well as the
// value for logBound.
func (mltp MaskedLinearTransformationProtocol) GenShare(skIn, skOut *rlwe.SecretKey, logBound uint, ct *rlwe.Ciphertext, crs multiparty.KeySwitchCRP, transform *MaskedLinearTransformationFunc, shareOut *multiparty.RefreshShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Generates the decryption share
// Returns [M_i] on mltp.tmpMask and [a*s_i -M_i + e] on EncToShareShare

// Applies LT(M_i)

// Stores the metadata of the ciphertext

// Returns [-a*s_i + LT(M_i) * diffscale + e] on ShareToEncShare

// AggregateShares sums share1 and share2 on shareOut.
func (mltp MaskedLinearTransformationProtocol) AggregateShares(share1, share2, shareOut *multiparty.RefreshShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Transform decrypts the ciphertext to LSSS-shares, applies the linear transformation on the LSSS-shares and re-encrypts the LSSS-shares to an RLWE ciphertext.
// The re-encrypted ciphertext's scale is set to the default scaling factor of the output parameters.
func (mltp MaskedLinearTransformationProtocol) Transform(ct *rlwe.Ciphertext, transform *MaskedLinearTransformationFunc, crs multiparty.KeySwitchCRP, share multiparty.RefreshShare, ciphertextOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Returns -sum(M_i) + x (outside of the NTT domain)

// Returns LT(-sum(M_i) + x)

// Extend the levels of the ciphertext for future allocation

// Updates the ciphertext metadata if the output dimensions is smaller

// Sets LT(-sum(M_i) + x) * diffscale in the RNS domain
// Positional -> RNS -> NTT

// LT(-sum(M_i) + x) * diffscale + [-a*s + LT(M_i) * diffscale + e] = [-a*s + LT(x) * diffscale + e]

// Copies the result on the out ciphertext

func (mltp MaskedLinearTransformationProtocol) applyTransformAndScale(transform *MaskedLinearTransformationFunc, metadata rlwe.MetaData, mask []*big.Int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Extracts sparse coefficients

// Decodes if asked to

// Applies the linear transform

// Recodes if asked to

// Puts the coefficient back

// Applies LT(M_i) * diffscale

// .Int truncates (i.e. does not round to the nearest integer)
// Thus we check if we are below, and if yes add 1, which acts as rounding to the nearest integer

// Scales the mask by the ratio between the two scales
