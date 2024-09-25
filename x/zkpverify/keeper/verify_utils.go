package keeper

import (
	"bytes"
	"crypto/sha256"

	"fiamma/x/zkpverify/types"
	"fiamma/x/zkpverify/verifiers"
)

// GetProofId returns the proof id
func (k Keeper) GetProofId(proofData types.ProofData) ([32]byte, error) {
	var buf bytes.Buffer
	buf.Write([]byte(proofData.Namespace))
	buf.Write([]byte(proofData.ProofSystem.String()))
	buf.Write(proofData.Proof)
	buf.Write(proofData.PublicInput)
	buf.Write(proofData.Vk)

	hash := sha256.Sum256(buf.Bytes())
	return hash, nil
}

func (k Keeper) verifyProof(proofData *types.ProofData) (bool, []byte) {
	switch proofData.ProofSystem {
	case types.ProofSystem_GROTH16_BN254_BITVM:

		vkLen := (uint32)(len(proofData.Vk))
		proofLen := (uint32)(len(proofData.Proof))
		pubInputLen := (uint32)(len(proofData.PublicInput))

		// TODO: The bitvm validation process returns an intermediate comment
		//  which needs to be submitted to bitcoin at a later date.
		verifyResult, witness := verifiers.VerifyBitvmProof(proofData.Vk, vkLen, proofData.Proof, proofLen, proofData.PublicInput, pubInputLen)
		k.Logger().Info("GROTH16 BN254 BitVM proof verification:", "result", verifyResult)
		return verifyResult, witness

	default:
		k.Logger().Error("Unrecognized proof system ID")
		return false, nil
	}
}
