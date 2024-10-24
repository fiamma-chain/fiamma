package keeper

import (
	"fiamma/x/zkpverify/types"
	"fiamma/x/zkpverify/verifiers"
)

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
