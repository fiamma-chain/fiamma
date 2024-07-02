package bitvm_test

import (
	"fiamma/x/zkpverify/verifiers/bitvm"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const MaxVkSize = 2 * 1024 * 1024
const MaxProofSize = 2 * 1024 * 1024
const MaxPublicInputSize = 2 * 1024 * 1024

func TestDummyBitvmProofVerifies(t *testing.T) {
	vkFile, err := os.Open("../../../../prover_examples/bitvm/vk.bitvm")
	if err != nil {
		t.Errorf("could not open vk file: %s", err)
	}
	vkBytes := make([]byte, MaxVkSize)
	nReadVkBytes, err := vkFile.Read(vkBytes)
	if err != nil {
		t.Errorf("could not read bytes from vk file")
	}

	proofFile, err := os.Open("../../../../prover_examples/bitvm/proof.bitvm")
	if err != nil {
		t.Errorf("could not open proof file: %s", err)
	}
	proofBytes := make([]byte, MaxProofSize)
	nReadProofBytes, err := proofFile.Read(proofBytes)
	if err != nil {
		t.Errorf("could not read bytes from proof file")
	}

	piFile, err := os.Open("../../../../prover_examples/bitvm/public_input.bitvm")
	if err != nil {
		t.Errorf("could not open public input file: %s", err)
	}

	piBytes := make([]byte, MaxPublicInputSize)
	nReadPiBytes, err := piFile.Read(piBytes)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	// witness is generated data from bitvm.
	success, witness := bitvm.VerifyBitvmProof(vkBytes, uint32(nReadVkBytes), proofBytes, uint32(nReadProofBytes), piBytes, uint32(nReadPiBytes))
	require.Equal(t, success, true)
	fmt.Println("witness: ", witness)
}
