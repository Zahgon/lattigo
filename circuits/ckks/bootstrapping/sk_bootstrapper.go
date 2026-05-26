package bootstrapping

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/ckks"
)

// SecretKeyBootstrapper is an implementation of the [rlwe.Bootstrapping] interface that
// uses the secret-key to decrypt and re-encrypt the bootstrapped ciphertext.
type SecretKeyBootstrapper struct {
	ckks.Parameters
	*ckks.Encoder
	*rlwe.Decryptor
	*rlwe.Encryptor
	sk       *rlwe.SecretKey
	Counter  int // records the number of bootstrapping
	MinLevel int
}

func NewSecretKeyBootstrapper(params ckks.Parameters, sk *rlwe.SecretKey) *SecretKeyBootstrapper {
	_ = "STUB: not implemented"
	return nil
}

func (d *SecretKeyBootstrapper) Bootstrap(ct *rlwe.Ciphertext) (*rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d SecretKeyBootstrapper) BootstrapMany(cts []rlwe.Ciphertext) ([]rlwe.Ciphertext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d SecretKeyBootstrapper) Depth() int { _ = "STUB: not implemented"; return 0 }

func (d SecretKeyBootstrapper) MinimumInputLevel() int { _ = "STUB: not implemented"; return 0 }

func (d SecretKeyBootstrapper) OutputLevel() int { _ = "STUB: not implemented"; return 0 }
