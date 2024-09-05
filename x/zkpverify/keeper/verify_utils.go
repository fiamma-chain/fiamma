package keeper

import (
	"bytes"
	"crypto/sha256"

	"fiamma/x/zkpverify/types"
	"fiamma/x/zkpverify/verifiers"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
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
	case types.ProofSystem_PLONK_BLS12_381:
		verifyResult := k.verifyPlonkProofBLS12_381(proofData.Proof, proofData.PublicInput, proofData.Vk)

		k.Logger().Info("PLONK BLS12-381 proof verification:", "result", verifyResult)
		return verifyResult, nil
	case types.ProofSystem_PLONK_BN254:
		verifyResult := k.verifyPlonkProofBN254(proofData.Proof, proofData.PublicInput, proofData.Vk)

		k.Logger().Info("PLONK BN254 proof verification:", "result", verifyResult)
		return verifyResult, nil
	case types.ProofSystem_GROTH16_BN254:
		verifyResult := k.verifyGroth16ProofBN254(proofData.Proof, proofData.PublicInput, proofData.Vk)

		k.Logger().Info("GROTH16 BN254 proof verification:", "result", verifyResult)
		return verifyResult, nil

	case types.ProofSystem_GROTH16_BN254_BITVM:

		vkLen := (uint32)(len(proofData.Vk))
		proofLen := (uint32)(len(proofData.Proof))
		pubInputLen := (uint32)(len(proofData.PublicInput))

		// TODO: The bitvm validation process returns an intermediate comment
		//  which needs to be submitted to bitcoin at a later date.
		verifyResult, witness := verifiers.VerifyBitvmProof(proofData.Vk, vkLen, proofData.Proof, proofLen, proofData.PublicInput, pubInputLen)
		k.Logger().Info("GROTH16 BN254 BitVM proof verification:", "result", verifyResult)
		return verifyResult, witness

	case types.ProofSystem_SP1:
		proofLen := (uint32)(len(proofData.Proof))
		// For the verification of the SP1 proof system, we consider the ELF file as public input.
		elfLen := (uint32)(len(proofData.PublicInput))
		verifyResult := verifiers.VerifySp1Proof(proofData.Proof, proofLen, proofData.PublicInput, elfLen)
		k.Logger().Info("SP1 proof verification:", "result", verifyResult)
		return verifyResult, nil

	default:
		k.Logger().Error("Unrecognized proof system ID")
		return false, nil
	}
}

// VerifyPlonkProofBLS12_381 verifies a PLONK proof using BLS12-381 curve.
func (k Keeper) verifyPlonkProofBLS12_381(proofBytes []byte, pubInputBytes []byte, verificationKeyBytes []byte) bool {
	return k.verifyPlonkProof(proofBytes, pubInputBytes, verificationKeyBytes, ecc.BLS12_381)
}

// VerifyGroth16ProofBN254 verifies a GROTH16 proof using BN254 curve.
func (k Keeper) verifyGroth16ProofBN254(proofBytes []byte, pubInputBytes []byte, verificationKeyBytes []byte) bool {
	return k.verifyGroth16Proof(proofBytes, pubInputBytes, verificationKeyBytes, ecc.BN254)
}

// VerifyPlonkProofBN254 verifies a PLONK proof using BN254 curve.
func (o Keeper) verifyPlonkProofBN254(proofBytes []byte, pubInputBytes []byte, verificationKeyBytes []byte) bool {
	return o.verifyPlonkProof(proofBytes, pubInputBytes, verificationKeyBytes, ecc.BN254)
}

// verifyPlonkProof contains the common proof verification logic.
func (k Keeper) verifyPlonkProof(proofBytes []byte, pubInputBytes []byte, verificationKeyBytes []byte, curve ecc.ID) bool {
	proofReader := bytes.NewReader(proofBytes)
	proof := plonk.NewProof(curve)
	if _, err := proof.ReadFrom(proofReader); err != nil {
		k.Logger().Info("Could not deserialize proof:", "error", err)
		return false
	}

	pubInputReader := bytes.NewReader(pubInputBytes)
	pubInput, err := witness.New(curve.ScalarField())
	if err != nil {
		k.Logger().Info("Error instantiating witness:", "error", err)
		return false
	}
	if _, err = pubInput.ReadFrom(pubInputReader); err != nil {
		k.Logger().Info("Could not read PLONK public input:", "error", err)
		return false
	}

	verificationKeyReader := bytes.NewReader(verificationKeyBytes)
	verificationKey := plonk.NewVerifyingKey(curve)
	if _, err = verificationKey.ReadFrom(verificationKeyReader); err != nil {
		k.Logger().Info("Could not read PLONK verifying key from bytes:", "error", err)
		return false
	}

	err = plonk.Verify(proof, verificationKey, pubInput)
	return err == nil
}

// verifyGroth16Proof contains the common proof verification logic.
func (k Keeper) verifyGroth16Proof(proofBytes []byte, pubInputBytes []byte, verificationKeyBytes []byte, curve ecc.ID) bool {
	proofReader := bytes.NewReader(proofBytes)
	proof := groth16.NewProof(curve)
	if _, err := proof.ReadFrom(proofReader); err != nil {
		k.Logger().Info("Could not deserialize proof:", "error", err)
		return false
	}

	pubInputReader := bytes.NewReader(pubInputBytes)
	pubInput, err := witness.New(curve.ScalarField())
	if err != nil {
		k.Logger().Info("Error instantiating witness:", "error", err)
		return false
	}
	if _, err = pubInput.ReadFrom(pubInputReader); err != nil {
		k.Logger().Info("Could not read Groth16 public input:", "error", err)
		return false
	}

	verificationKeyReader := bytes.NewReader(verificationKeyBytes)
	verificationKey := groth16.NewVerifyingKey(curve)
	if _, err = verificationKey.ReadFrom(verificationKeyReader); err != nil {
		k.Logger().Info("Could not read Groth16 verifying key from bytes:", "error", err)
		return false
	}

	err = groth16.Verify(proof, verificationKey, pubInput)
	return err == nil
}
