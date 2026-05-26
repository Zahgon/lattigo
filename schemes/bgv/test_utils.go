package bgv

import (
	"testing"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

type TestContext struct {
	Params Parameters
	Ecd    *Encoder

	Prng    sampling.PRNG
	Sampler *ring.UniformSampler

	Kgen *rlwe.KeyGenerator
	Sk   *rlwe.SecretKey
	Pk   *rlwe.PublicKey

	Enc *rlwe.Encryptor
	Dec *rlwe.Decryptor

	Evl *Evaluator
}

func NewTestContext(params ParametersLiteral, scaleInvariant bool) *TestContext {
	_ = "STUB: not implemented"
	return nil
}

func (tc TestContext) String() string { _ = "STUB: not implemented"; return "" }

func VerifyTestVectors(params Parameters, encoder *Encoder, decryptor *rlwe.Decryptor, have interface{}, want []uint64, isBatched bool, t *testing.T) {
	_ = "STUB: not implemented"
	return
}

// fmt.Println("have", values[:10])
// fmt.Println("want", want[:10])

func NewTestVector(params Parameters, encoder *Encoder, encryptor *rlwe.Encryptor, level int, scale rlwe.Scale, batched bool) (values []uint64, pt *rlwe.Plaintext, ct *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
