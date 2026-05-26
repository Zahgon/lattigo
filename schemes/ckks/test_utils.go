package ckks

import (
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/ring"
	"github.com/tuneinsight/lattigo/v6/utils/bignum"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

type TestContext struct {
	Params Parameters
	Ecd    *Encoder

	Prng sampling.PRNG

	Kgen *rlwe.KeyGenerator
	Sk   *rlwe.SecretKey
	Pk   *rlwe.PublicKey

	Enc *rlwe.Encryptor
	Dec *rlwe.Decryptor

	Evl *Evaluator
}

func NewTestContext(params ParametersLiteral) *TestContext { _ = "STUB: not implemented"; return nil }

func (tc *TestContext) NewTestVector(a, b complex128) (values []*bignum.Complex, pt *rlwe.Plaintext, ct *rlwe.Ciphertext) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func randomConst(tp ring.Type, prec uint, a, b complex128) (constant *bignum.Complex) {
	_ = "STUB: not implemented"
	return nil
}

var (
	// testInsecurePrec45 are insecure parameters used for the sole purpose of fast testing.
	testInsecurePrec45 = ParametersLiteral{
		LogN:            10,
		LogQ:            []int{55, 45, 45, 45, 45, 45, 45},
		LogP:            []int{60},
		LogDefaultScale: 45,
	}

	// testInsecurePrec90 are insecure parameters used for the sole purpose of fast testing.
	testInsecurePrec90 = ParametersLiteral{
		LogN:            10,
		LogQ:            []int{55, 55, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45},
		LogP:            []int{60, 60},
		LogDefaultScale: 90,
	}

	testParametersLiteral = []ParametersLiteral{testInsecurePrec45, testInsecurePrec90}
)

// testInsecurePrec45 = ParametersLiteral{
// 	LogN: 10,
// 	Q: []uint64{
// 		0x80000000080001,
// 		0x2000000a0001,
// 		0x2000000e0001,
// 		0x2000001d0001,
// 		0x1fffffcf0001,
// 		0x1fffffc20001,
// 		0x200000440001,
// 	},
// 	P: []uint64{
// 		0x80000000130001,
// 		0x7fffffffe90001,
// 	},
// 	LogDefaultScale: 45,
// }

// testInsecurePrec90 = ParametersLiteral{
// 	LogN: 10,
// 	Q: []uint64{
// 		0x80000000080001,
// 		0x80000000440001,
// 		0x2000000a0001,
// 		0x2000000e0001,
// 		0x1fffffc20001,
// 		0x200000440001,
// 		0x200000500001,
// 		0x200000620001,
// 		0x1fffff980001,
// 		0x2000006a0001,
// 		0x1fffff7e0001,
// 		0x200000860001,
// 	},
// 	P: []uint64{
// 		0xffffffffffc0001,
// 		0x10000000006e0001,
// 	},
// 	LogDefaultScale: 90,
// }

// testParamsLiteral = []ParametersLiteral{testInsecurePrec45, testInsecurePrec90}
