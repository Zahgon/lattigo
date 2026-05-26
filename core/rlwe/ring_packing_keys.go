package rlwe

// RingPackingEvaluationKey is a struct storing the
// ring packing evaluation keys.
// All fields of this struct are public, enabling
// custom instantiations.
type RingPackingEvaluationKey struct {
	// Parameters are the different Parameters among
	// which a ciphertext will be switched during the
	// procedure. These parameters share the same primes
	// but support different ring degrees.
	Parameters map[int]ParameterProvider

	// RingSwitchingKeys are the ring degree switching keys
	// indexed as map[inputLogN][outputLogN]
	RingSwitchingKeys map[int]map[int]*EvaluationKey

	// RepackKeys are the [EvaluationKey] used for the
	// RLWE repacking.
	RepackKeys map[int]EvaluationKeySet

	// ExtractKeys are the [EvaluationKey] used for the
	// RLWE extraction.
	ExtractKeys map[int]EvaluationKeySet
}

// MinLogN returns the minimum Log(N) among the supported ring degrees.
// This method requires that the field Parameters of [RingPackingEvaluationKey]
// has been populated.
func (rpk RingPackingEvaluationKey) MinLogN() (minLogN int) { _ = "STUB: not implemented"; return 0 }

// MaxLogN returns the maximum Log(N) among the supported ring degrees.
// This method requires that the field Parameters of [RingPackingEvaluationKey]
// has been populated.
func (rpk RingPackingEvaluationKey) MaxLogN() (maxLogN int) { _ = "STUB: not implemented"; return 0 }

// GenRingSwitchingKeys generates the [Parameter]s and [EvaluationKey]s
// to be able to split an [Ciphertext] into two [Ciphertext]s of half
// the ring degree and merge two [Ciphertext]s into one [Ciphertext]
// of twice the ring degree.
//
// The method returns the [Parameter]s, [EvaluationKey]s and ephemeral
// [SecretKey]s used to generate the ring-switching [EvaluationKey]s.
//
// See the methods [RingPackingEvaluator.Split] and [RingPackingEvaluator.Repack].
//
// This function will return an error if minLogN >= params.LogN().
func (rpk *RingPackingEvaluationKey) GenRingSwitchingKeys(params ParameterProvider, sk *SecretKey, minLogN int, evkParams EvaluationKeyParameters) (ski map[int]*SecretKey, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ring switching evaluation keys

// GenRepackEvaluationKeys generates the set of params.LogN() [EvaluationKey]s necessary to perform the repacking operation.
// See [RingPackingEvaluator.Repack] for additional information.
func (rpk *RingPackingEvaluationKey) GenRepackEvaluationKeys(params ParameterProvider, sk *SecretKey, evkParams EvaluationKeyParameters) {
	_ = "STUB: not implemented"
	return
}

// GenExtractEvaluationKeys generates the set of params.LogN() [EvaluationKey]s necessary to perform the extraction operation.
// See [RingPackingEvaluator.Extract] for additional information.
func (rpk *RingPackingEvaluationKey) GenExtractEvaluationKeys(params ParameterProvider, sk *SecretKey, evkParams EvaluationKeyParameters) {
	_ = "STUB: not implemented"
	return
}

// GaloisElementsForExpand returns the list of Galois elements required
// to perform the `Expand` operation with parameter `logN`.
func GaloisElementsForExpand(params ParameterProvider, logN int) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// GaloisElementsForPack returns the list of Galois elements required to perform the `Pack` operation.
func GaloisElementsForPack(params ParameterProvider, logGap int) (galEls []uint64) {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check
