// This example demonstrates the use of the [multiparty] package to perform a basic N-party private information retrieval (PIR) protocol.
// This protocol relies on the t-out-of-N-threshold variant of the BGV scheme.
//
// In a nutshell, each party uploads an data row to a helper server, as an encrypted integer vector. The result is an encrypted database of N rows. At a later stage,
// a party queries a row in the database to the helper, without revealing the selector query index. The querying party does so by encoding its query as
// a binary vector with a single 1 component corresponding to the row it wants to retrieve, and send this encrypted vector to the helper.
// The helper can then compute the response (under encryption) by multiplying each row i of the database with the i-th query component, and summing
// the resulting vectors together. Finally, the result can be decrypted by any group of t parties.
//
// Thanks to the t-out-of-N-threshold scheme, only t parties need to be online at query time.
//
// For more details about the PIR circuit example see the paper [Multiparty Homomorphic Encryption from Ring-Learning-With-Errors] by Mouchet, Troncoso-Pastoriza, Bossuat, and Hubaux.
// For more details about the t-out-of-N-threshold scheme, see the paper [An Efficient Threshold Access-Structure for RLWE-Based Multiparty Homomorphic Encryption] by Mouchet, Bertrand and Hubaux.
//
// To run the example, use the following command:
//
//	go run main.go N T NGoRoutines
//
// where N is the number of parties (default:3) T is the threshold (default: 2) and NGoRoutines is the number of Go routines (default: 1) to use during the homomorphic evaluation.
// All parties are run in the same process.
//
// The example demonstrates the following steps:
//
//  1. Setup:
//     a. The parties generate threshold secret key with t-out-of-N access structure.
//     b. The parties generate a collective public encryption key, a relinearization key, and a set of rotation (or, galois) keys.
//  2. Database inputs: Each party encrypts its input row vector and send it to a helper server.
//  3. Query evaluation: A party requests a row in the database. The helper server computes the query output as described above.
//  4. Query output decryption: the helper, with the help of at least t parties, decrypts the query output.
//
// [Multiparty Homomorphic Encryption from Ring-Learning-With-Errors]: https://eprint.iacr.org/2020/304
// [An Efficient Threshold Access-Structure for RLWE-Based Multiparty Homomorphic Encryption]: https://eprint.iacr.org/2022/780
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/multiparty"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// party is a type for the parties' state in the protocol, to be kept accross the different phases.
type party struct {
	multiparty.Combiner

	sk         *rlwe.SecretKey // secret key of the party
	tsk        multiparty.ShamirSecretShare
	rlkEphemSk *rlwe.SecretKey // ephemeral state to be used in the RKG protocol

	shamirPt multiparty.ShamirPublicPoint

	input []uint64 // the input of the party, encoding a set as a binary vector
}

// getOnlineParties is a utility function that returns t random parties from a list of parties.
// This simulates a dynamic setting where the system rely on any t parties to be online at query time to
// execute the various protocols.
func getOnlineParties(t int, parties []party) []party { _ = "STUB: not implemented"; return nil }

// randomizes a subset of t parties

// getShamirPoints is a utility function that returns the Shamir public points of a group of parties.
func getShamirPoints(parties []party) []multiparty.ShamirPublicPoint {
	_ = "STUB: not implemented"
	return nil
}

var l = log.New(os.Stderr, "", 0)

func main() {

	// Parse command line arguments

	N := 3 // Default number of parties
	var err error
	if len(os.Args[1:]) >= 1 {
		N, err = strconv.Atoi(os.Args[1])
		check(err)

		if N < 3 || N > 128 {
			l.Fatal("N must be in the range [3, 128]")
		}
	}

	t := 2 // Default Threshold
	if len(os.Args[1:]) >= 2 {
		t, err = strconv.Atoi(os.Args[2])
		check(err)

		if t < 2 || t >= N {
			l.Fatal("T must be in the range [2, N-1]")
		}
	}

	nGoRoutine := 1 // Default number of Go routines
	if len(os.Args[1:]) >= 3 {
		nGoRoutine, err = strconv.Atoi(os.Args[2])
		check(err)

		if nGoRoutine < 1 {
			l.Fatal("NGoRoutine must be at least 1")
		}
	}

	// Creating encryption parameters
	// LogN = 13 & LogQP = 218
	params, err := bgv.NewParametersFromLiteral(bgv.ParametersLiteral{
		LogN:             13,
		LogQ:             []int{54, 54, 54},
		LogP:             []int{55},
		PlaintextModulus: 65537,
	})
	if err != nil {
		panic(err)
	}

	// The circuit relies on rotations/automorphisms for computing an inner-sum-like operation.
	// This obtains the corresponding Galois elements.
	galEls := append(params.GaloisElementsForInnerSum(1, params.N()>>1), params.GaloisElementForRowRotation())

	// Creates a PRNG that will be used to sample the common reference string (crs)
	crs, err := sampling.NewKeyedPRNG([]byte{'l', 'a', 't', 't', 'i', 'g', 'o'})
	if err != nil {
		panic(err)
	}

	// Create the N input parties and generate their secret keys and private inputs
	P := genparties(params, N, t)

	l.Printf("========= Setup phase =========")

	// Step 1.a: Generation of the threshold secret key

	thresholdizer := multiparty.NewThresholdizer(params.Parameters)

	// Allocates the memory for the parties' shares in the protocol
	tskShares := make([][]multiparty.ShamirSecretShare, N)
	for i := range P {
		tskShares[i] = make([]multiparty.ShamirSecretShare, N)
		for j := range P {
			tskShares[i][j] = thresholdizer.AllocateThresholdSecretShare()
		}
	}

	l.Println("> Threshold Secret Key Generation")
	// The parties generate their shares for the threshold secret key
	elapsedThresholdParty := runTimedParty(func() {
		for i, pi := range P {
			pol, err := thresholdizer.GenShamirPolynomial(t, pi.sk)
			check(err)
			for j, pj := range P {
				thresholdizer.GenShamirSecretShare(pj.shamirPt, pol, &tskShares[i][j])
			}
		}
	}, N)

	// Each party aggregates the shares it has received from the other parties
	elapsedThresholdParty += runTimedParty(func() {
		for i := range P {
			P[i].tsk = thresholdizer.AllocateThresholdSecretShare()
			for j := range P {
				err := thresholdizer.AggregateShares(tskShares[j][i], P[i].tsk, &P[i].tsk)
				check(err)
			}
		}
	}, N)

	l.Printf("\tdone (cloud: %s, party: %s)\n", time.Duration(0), elapsedThresholdParty)

	// Step 1.b: Setup of the collective public encryption and evaluation keys
	// Note 1: because we are using the t-out-of-N-threshold scheme, generating a key only requires t parties to be online.
	// Note 2: the above note means that the workload could be balanced between the online parties whenever more than t parties are online.

	pk := execCKGProtocol(params, crs, getOnlineParties(t, P)) // Collective public key generation

	rlk := execRKGProtocol(params, crs, getOnlineParties(t, P)) // Collective RelinearizationKey generation

	galKeys := execGTGProtocol(params, crs, galEls, getOnlineParties(t, P)) // Collective GaloisKeys generation

	// Creates the evaluation key set from the rlk and the galKeys
	evk := rlwe.NewMemEvaluationKeySet(rlk, galKeys...)

	l.Printf("Setup done (cloud: %s, party: %s)\n",
		elapsedCKGCloud+elapsedRKGCloud+elapsedGKGCloud,
		elapsedCKGParty+elapsedRKGParty+elapsedGKGParty)

	// Step 2: Database inputs

	// Pre-allocates the encrypted database: one ciphertext per party
	encInputs := make([]*rlwe.Ciphertext, N)
	for i := range encInputs {
		encInputs[i] = bgv.NewCiphertext(params, 1, params.MaxLevel())
	}

	// Each party encrypts its input row under the collective public key
	l.Println("========= Database input phase ==========")
	l.Println("> Input Encryption")
	encoder := bgv.NewEncoder(params)
	encryptor := rlwe.NewEncryptor(params, pk)
	pt := bgv.NewPlaintext(params, params.MaxLevel())
	elapsedEncryptParty := runTimedParty(func() {
		for i, pi := range P {
			if err := encoder.Encode(pi.input, pt); err != nil {
				panic(err)
			}
			if err := encryptor.Encrypt(pt, encInputs[i]); err != nil {
				panic(err)
			}
		}
	}, N)

	l.Printf("\tdone (cloud: %s, party: %s)\n", elapsedEncryptCloud, elapsedEncryptParty)

	// Step 3: Query evaluation

	l.Println("========= Query evaluation phase =========")

	queryIndex := 2 // Index of the ciphertext to retrieve.
	querier := P[0] // Party performing the query
	others := P[1:] // Other parties

	// Creates a query ciphertext from the query index
	encQuery := genQuery(params, queryIndex, encoder, encryptor)

	// Executes the requests
	encResult := execRequest(params, nGoRoutine, encQuery, encInputs, evk)

	// Step 4: Query output decryption

	// The helper, with the help of at least t parties, performs a re-encryption of the result towards the querier
	participants := getOnlineParties(t-1, others)
	encOut := execCKSProtocol(params, participants, querier, encResult)

	// The querier decrypts the final result
	sk := rlwe.NewSecretKey(params)
	err = querier.Combiner.GenAdditiveShare(append(getShamirPoints(participants), querier.shamirPt), querier.shamirPt, querier.tsk, sk)
	check(err)

	decryptor := rlwe.NewDecryptor(params, sk)
	ptres := bgv.NewPlaintext(params, params.MaxLevel())
	elapsedDecParty := runTimed(func() {
		decryptor.Decrypt(encOut, ptres)
	})

	res := make([]uint64, params.MaxSlots())
	if err := encoder.Decode(ptres, res); err != nil {
		panic(err)
	}

	l.Printf("> Finished (total cloud: %s, total party: %s)\n",
		elapsedCKGCloud+elapsedRKGCloud+elapsedGKGCloud+elapsedEncryptCloud+elapsedRequestCloudCPU+elapsedCKSCloud,
		elapsedCKGParty+elapsedRKGParty+elapsedGKGParty+elapsedEncryptParty+elapsedRequestParty+elapsedCKSParty+elapsedDecParty)

	l.Printf("Result: %v...%v\n", res[:8], res[params.N()-8:])
}

func genparties(params bgv.Parameters, N, t int) []party { _ = "STUB: not implemented"; return nil }

/* #nosec G115 -- i cannot be negative */

/* #nosec G115 -- i cannot be negative */

func execCKGProtocol(params bgv.Parameters, crs sampling.PRNG, participants []party) *rlwe.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

// Creates a protocol type for the collective public key generation.
// The type is stateless and can be used to generate as many public keys as needed.

// Allocates the memory for the parties' shares in the protocol

// the public CKG shares
// the t-out-of-t secret keys for group P

// Allocate the memory for the combined share

// sample the common reference polynomial (crp) from the common reference string (crs)

// Generate the parties' shares

// Generate the t-out-of-t secret key of the party within the group of participants

// Generate the public key share of the party from the t-out-of-t secret key

// Aggregate the parties' shares into a collective public key

// Aggregate the parties' shares into a combined share

// Generate the public key from the combined share

func execRKGProtocol(params bgv.Parameters, crs sampling.PRNG, participants []party) *rlwe.RelinearizationKey {
	_ = "STUB: not implemented"
	return nil
}

// Creates a protocol type for the collective relinearization key generation.
// The type is stateless and can be used to generate as many relinearization keys as needed.
// The RKG protocol has two rounds. Because the ephemeral secret key is not re-shared,
// the same set of participants must participate to the two rounds.

// Allocates the memory for the parties' shares in the protocol

// the parties have a private ephemeral secret key in the RKGen protocol

// Allocate the memory for the combined public shares

// Sample the common reference polynomial (crp) common reference string (crs)

// The parties generate their shares for round one

// Generate the t-out-of-t secret key of the party within the group of participants

// Generate the shares for round one from the t-out-of-t secret key

// the helper aggregates the parties' shares for round one

// The parties generate their shares for round two

// Generate the shares for round two from the t-out-of-t secret key
// Note: the same set of participants must participate to the two rounds, so the same tsk is used.

// the helper aggregates the parties' shares for round two and generates the relinearization key

func execGTGProtocol(params bgv.Parameters, crs sampling.PRNG, galEls []uint64, participants []party) (galKeys []*rlwe.GaloisKey) {
	_ = "STUB: not implemented"
	return nil
}

// Creates a protocol type for the collective galois key generation.
// The type is stateless and can be used to generate as many galois keys as needed.
// Rotation keys generation

// Allocates the memory for the parties' shares in the protocol

// Allocate a slice for storing the output keys

// Runs the GKG protocol for each required Galois key
// Note: this demo re-uses the allocated shares for each execution.

// Sample the common reference polynomial (crp) common reference string (crs)

// The parties generate their shares for the Galois key generation protocol

// Generate the t-out-of-t secret key of the party within the group of participants

// Generate the shares for the Galois key generation protocol from the t-out-of-t secret key

// The helper aggregates the parties' shares and generates the Galois key

// Allocate the memory for the combined share

func genQuery(params bgv.Parameters, queryIndex int, encoder *bgv.Encoder, encryptor *rlwe.Encryptor) *rlwe.Ciphertext {
	_ = "STUB: not implemented"
	return nil
}

// Creates a query vector from the query index

// Encrypts the query vector

func execRequest(params bgv.Parameters, NGoRoutine int, encQuery *rlwe.Ciphertext, encInputs []*rlwe.Ciphertext, evk rlwe.EvaluationKeySet) *rlwe.Ciphertext {
	_ = "STUB: not implemented"
	return nil
}

// First, pre-compute the plaintext masks for the query evaluation as:
// plainmask[i] = encode([0, ..., 0, 1, 0, ..., 0])  (zero with a 1 at the i-th position).
// In practice, the masks are pre-computed and reused accross queries.

// Buffer for the intermediate computation done by the helper

// Creates an evaluator for the homomorphic evaluation

// Split the task among the Go routines

// maskTask is a type for the task to be executed by the Go routines
// The task computes the multiplication of the query with a mask, and the multiplication of the result with a row of the database.

// 1) Multiplication BFV-style of the query with the plaintext mask

// 2) Inner sum (populate all the slots with the sum of all the slots)

// 3) Multiplication of 2) with the i-th ciphertext stored in the cloud

// Wait for all the workers to finish

// collects the elapsed time for each task

// Creates ciphertexts to store the final result
// to receive the sum of the partial results.
// to receive the relinearized final result

// Summation of all the partial result among the different Go routines
// The sum is computed over the degree-2 ciphertexts from the previous step. Then, the result is relinearized.
// This avoids performing N relinearizations.

func execCKSProtocol(params bgv.Parameters, participants []party, receiver party, ctIn *rlwe.Ciphertext) *rlwe.Ciphertext {
	_ = "STUB: not implemented"
	return nil
}

// Creates a protocol type for the collective key-switching protocol, with smudging distribution parameter of 2^30.
// The type is stateless and can be used to generate as many key-switching keys as needed.

// Allocates the memory for the parties' shares in the protocol

// Allocate the memory for the public share
// Allocate the memory for the t-out-of-t secret key

// Allocate the memory for the combined share

// To generate a re-encryption of the result ciphertexts towards the querier,
// each party except for the receiver generates a key-switching share towards
// secret-key zero (i.e., a decryption share).

// The parties (except the receiver) generate their shares for the key-switching protocol

// Generate the t-out-of-t secret key with the reciever and t-1 other parties

// Generate the key-switching share with the t-out-of-t secret key

// The helper aggregates the parties' shares and generates the key-switching key

// Aggregate the parties' shares into a combined share

// Generate the re-encryption from the combined share

var (
	elapsedCKGCloud        time.Duration
	elapsedCKGParty        time.Duration
	elapsedRKGCloud        time.Duration
	elapsedRKGParty        time.Duration
	elapsedGKGCloud        time.Duration
	elapsedGKGParty        time.Duration
	elapsedCKSCloud        time.Duration
	elapsedCKSParty        time.Duration
	elapsedEncryptCloud    time.Duration
	elapsedRequestParty    time.Duration
	elapsedRequestCloud    time.Duration
	elapsedRequestCloudCPU time.Duration
)

func check(err error) { _ = "STUB: not implemented"; return }

func runTimed(f func()) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func runTimedParty(f func(), N int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
