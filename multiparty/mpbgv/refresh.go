package mpbgv

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/multiparty"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
)

// RefreshProtocol is a struct storing the relevant parameters for the Refresh protocol.
type RefreshProtocol struct {
	MaskedTransformProtocol
}

// NewRefreshProtocol creates a new Refresh protocol instance.
func NewRefreshProtocol(params bgv.Parameters, noiseFlooding ring.DistributionParameters) (rfp RefreshProtocol, err error) {
	_ = "STUB: not implemented"
	return *new(RefreshProtocol), nil
}

// AllocateShare allocates the shares of the PermuteProtocol
func (rfp RefreshProtocol) AllocateShare(inputLevel, outputLevel int) multiparty.RefreshShare {
	_ = "STUB: not implemented"
	return *new(multiparty.RefreshShare)
}

// GenShare generates a share for the Refresh protocol.
// ct1 is degree 1 element of a rlwe.Ciphertext, i.e. rlwe.Ciphertext.Value[1].
func (rfp RefreshProtocol) GenShare(sk *rlwe.SecretKey, ct *rlwe.Ciphertext, crp multiparty.KeySwitchCRP, shareOut *multiparty.RefreshShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AggregateShares aggregates two parties' shares in the Refresh protocol.
func (rfp RefreshProtocol) AggregateShares(share1, share2 multiparty.RefreshShare, shareOut *multiparty.RefreshShare) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Finalize applies Decrypt, Recode and Recrypt on the input ciphertext.
func (rfp RefreshProtocol) Finalize(ctIn *rlwe.Ciphertext, crp multiparty.KeySwitchCRP, share multiparty.RefreshShare, opOut *rlwe.Ciphertext) (err error) {
	_ = "STUB: not implemented"
	return nil
}
