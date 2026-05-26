// This example demonstrates the use of the [multiparty] package to perform a basic N-party private set intersection (PSI) protocol with external receiver.
// This protocol relies on the N-out-of-N threshold variant of the BGV scheme.
// In a nutshell, the parties encode their sets as binary vectors, and a helper server compute (under encryption) the intersection as the compoenent-wise product of all vectors.
// The result is then re-encrypted towards the receiver's key.
// For more details about the PSI circuit example see the paper [Multiparty Homomorphic Encryption from Ring-Learning-With-Errors] by by Christian Mouchet, Juan Troncoso-Pastoriza, Jean-Philippe Bossuat, and Jean-Pierre Hubaux.
//
// To run the example, use the following command:
//
//	go run main.go N NGoRoutines
//
// where N is the number of parties (default:16) and NGoRoutines is the number of Go routines (default: 1) to use during the homomorphic evaluation.
// All parties are run in the same process.
//
// The example demonstrates the following steps:
//
//  1. Setup: The parties generate a collectice public encryption key and a collective relinearization key.
//  2. Inputs: Each party encrypts its input vector and send it to a helper server.
//  3. Evaluation: The helper server computes the multiplication of the input vectors and relinearizes the result.
//  4. Output: The helper server, with the help of the N parties, switches the encryption of the result to the target public key.
//  5. Decryption: The target party decrypts the result with its secret key.
//
// [Multiparty Homomorphic Encryption from Ring-Learning-With-Errors]: https://eprint.iacr.org/2020/304
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
	"github.com/tuneinsight/lattigo/v6/utils/sampling"
)

// party is a type for the parties' state in the protocol, to be kept accross the different phases.
type party struct {
	sk         *rlwe.SecretKey // secret key of the party
	rlkEphemSk *rlwe.SecretKey // ephemeral state to be used in the RKG protocol

	input []uint64 // the input of the party, encoding a set as a binary vector
}

var l = log.New(os.Stderr, "", 0)

func main() {

	// $go run main.go N NGoRoutine
	// N: number of parties
	// NGoRoutines: number of Go routines
	N := 16 // Default number of parties
	var err error
	if len(os.Args[1:]) >= 1 {
		N, err = strconv.Atoi(os.Args[1])
		check(err)

		if N < 2 || N > 128 {
			l.Fatal("N must be in the range [2, ..., 128]")
		}
	}

	NGoRoutine := 1 // Default number of Go routines
	if len(os.Args[1:]) >= 2 {
		NGoRoutine, err = strconv.Atoi(os.Args[2])
		check(err)

		if NGoRoutine < 1 {
			l.Fatal("NGoRoutine must be at least 1")
		}
	}

	// Creating encryption parameters from a default params with logN=14, logQP=438 with a plaintext modulus T=65537
	params, err := bgv.NewParametersFromLiteral(bgv.ParametersLiteral{
		LogN:             14,
		LogQ:             []int{56, 55, 55, 54, 54, 54},
		LogP:             []int{55, 55},
		PlaintextModulus: 65537,
	})
	check(err)

	// Creates a PRNG that will be used to sample the common reference string (crs)
	crs, err := sampling.NewKeyedPRNG([]byte{'l', 'a', 't', 't', 'i', 'g', 'o'})
	check(err)

	// Generate some keys for the receiver (target party)
	tsk, tpk := rlwe.NewKeyGenerator(params).GenKeyPairNew()

	// Create the N input parties and generate their secret keys
	P := genparties(params, N)

	// Step 1: Setup of the collective public key and relinearization key
	l.Printf("========= Setup Phase =========")

	pk := execCKGProtocol(params, crs, P) // generates the collective public key

	rlk := execRKGProtocol(params, crs, P) // generates the collective relinearization key

	evk := rlwe.NewMemEvaluationKeySet(rlk) // creates the evaluation key from the relinearization key

	l.Printf("Setup done (cloud: %s, party: %s)\n",
		elapsedRKGCloud+elapsedCKGCloud, elapsedRKGParty+elapsedCKGParty)

	// Step 2: Each party encrypts its input vector
	l.Printf("========= Computation Phase =========")

	expRes := genInputs(params, P) // generates the input vectors and the expected result

	encoder := bgv.NewEncoder(params)
	encInputs := inputPhase(params, P, pk, encoder) // encrypts the input vectors

	// Step 3: The helper server computes the multiplication of the input vectors and relinearizes the result
	encRes := evalPhase(params, NGoRoutine, encInputs, evk)

	// Step 4: The helper server switches the encryption of the result to the target public key
	encOut := execPCKSProtocol(params, tpk, encRes, P)

	// Step 5: The target party decrypts the result with their secret key
	l.Println("> Result Decryption")
	decryptor := rlwe.NewDecryptor(params, tsk)
	ptres := bgv.NewPlaintext(params, params.MaxLevel())
	elapsedDecParty := runTimed(func() {
		decryptor.Decrypt(encOut, ptres)
	})
	l.Printf("\tdone (cloud: %s, party: %s)\n", time.Duration(0), elapsedDecParty)

	// Check the result
	res := make([]uint64, params.MaxSlots())
	err = encoder.Decode(ptres, res)
	check(err)

	l.Printf("\tResult: %v\n", res[:16])
	for i := range expRes {
		if expRes[i] != res[i] {
			//l.Printf("\t%v\n", expRes)
			l.Println("\tincorrect")
			return
		}
	}
	l.Println("\tCorrect")
	l.Printf("Finished (total cloud: %s, total party: %s)\n",
		elapsedCKGCloud+elapsedRKGCloud+elapsedEncryptCloud+elapsedEvalCloud+elapsedPCKSCloud,
		elapsedCKGParty+elapsedRKGParty+elapsedEncryptParty+elapsedEvalParty+elapsedPCKSParty+elapsedDecParty)
}

func genparties(params bgv.Parameters, N int) []party {
	_ = "STUB: not implemented"

	// Create the parties and generates a secret key for each party
	return nil
}

func execCKGProtocol(params bgv.Parameters, crs sampling.PRNG, P []party) *rlwe.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

// Creates a protocol type for the collective public key generation.
// The type is stateless and can be used to generate as many public keys as needed.

// Allocates the memory for the parties' shares in the protocol

// Allocate the memory for the combined share

// sample the common reference polynomial (crp) from the common reference string (crs)

// Generate the parties' shares

// Aggregate the parties' shares into a collective public key

// Aggregate the parties' shares into a combined share

// Generate the public key from the combined share

func execRKGProtocol(params bgv.Parameters, crs sampling.PRNG, P []party) *rlwe.RelinearizationKey {
	_ = "STUB: not implemented"
	return nil
}

// Creates a protocol type for the collective relinearization key generation.
// The type is stateless and can be used to generate as many relinearization keys as needed.
// The RKG protocol has two rounds.

// Allocates the memory for the parties' shares in the protocol

// the parties have a private ephemeral secret key in the RKGen protocol

// Allocate the memory for the combined public shares

// Sample the common reference polynomial (crp) from the common reference string (crs)

// The parties generate their shares for round one

// the helper aggregates the parties' shares for round one

// The parties generate their shares for round two

// the helper aggregates the parties' shares for round two and generates the relinearization key

func genInputs(params bgv.Parameters, P []party) (expRes []uint64) {
	_ = "STUB: not implemented"

	// generate input vectors for the parties of max size
	return nil
}

func inputPhase(params bgv.Parameters, P []party, pk *rlwe.PublicKey, encoder *bgv.Encoder) (encInputs []*rlwe.Ciphertext) {
	_ = "STUB: not implemented"

	// Allocate the memory for the encrypted input vectors
	return nil
}

// Each party encrypts its input vector

func evalPhase(params bgv.Parameters, NGoRoutine int, encInputs []*rlwe.Ciphertext, evk rlwe.EvaluationKeySet) (encRes *rlwe.Ciphertext) {
	_ = "STUB: not implemented"

	// The eval phase performs the multiplication as a balanced binary tree.
	// For each level of the tree, it performs the multiplications in parallel using at most NGoRoutine Go routines.
	return nil
}

// Allocate the memory for the encrypted result at each level of the tree

// Creates a evaluator for the multiplication, with the evaluation key

// Split the task among the Go routines
// A multTask is a task that multiplies two ciphertexts and relinearizes the result

//l.Println("> Spawning", NGoRoutine, "evaluator goroutine")

// 1) Multiplication of two input vectors

// 2) Relinearization

// Start the tasks

//l.Println("> Shutting down workers")

func execPCKSProtocol(params bgv.Parameters, tpk *rlwe.PublicKey, encRes *rlwe.Ciphertext, P []party) (encOut *rlwe.Ciphertext) {
	_ = "STUB: not implemented"

	// Collective key switching from the collective secret key to
	// the target public key
	return nil
}

// Creates a protocol type for the collective public key switch.
// The type is stateless and can be used to generate as many public key switches as needed.

// Allocates the memory for the parties' shares in the protocol

// Allocates the memory for combined share

// Each party generates its share

// The helper server aggregates the parties' shares and combutes the output, re-encrpyted, ciphertext

// Aggregate the parties' shares into a combined share

// Generate the output ciphertext

var (
	elapsedEncryptParty,
	elapsedEncryptCloud,
	elapsedCKGCloud,
	elapsedCKGParty,
	elapsedRKGCloud,
	elapsedRKGParty,
	elapsedPCKSCloud,
	elapsedPCKSParty,
	elapsedEvalCloudCPU,
	elapsedEvalCloud,
	elapsedEvalParty time.Duration
)

func check(err error) { _ = "STUB: not implemented"; return }

func runTimed(f func()) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func runTimedParty(f func(), N int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
