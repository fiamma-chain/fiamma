package keeper

import (
	"bytes"

	"fiamma/x/zkpverify/types"
	"fiamma/x/zkpverify/verifiers/sp1"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
)

func (k Keeper) verifyProof(verifyData *types.VerifyData) bool {
	switch verifyData.ProofSystem {
	case uint64(types.PlonkBls12_381):
		verifyResult := k.verifyPlonkProofBLS12_381(verifyData.Proof, verifyData.PublicInput, verifyData.Vk)

		k.Logger().Info("PLONK BLS12-381 proof verification:", "result", verifyResult)
		return verifyResult
	case uint64(types.PlonkBn254):
		verifyResult := k.verifyPlonkProofBN254(verifyData.Proof, verifyData.PublicInput, verifyData.Vk)

		k.Logger().Info("PLONK BN254 proof verification:", "result", verifyResult)
		return verifyResult
	case uint64(types.Groth16Bn254):
		verifyResult := k.verifyGroth16ProofBN254(verifyData.Proof, verifyData.PublicInput, verifyData.Vk)

		k.Logger().Info("GROTH16 BN254 proof verification:", "result", verifyResult)
		return verifyResult

	case uint64(types.Groth16Bn254_BTC):
		panic("Not implemented")

	case uint64(types.SP1):
		proofLen := (uint32)(len(verifyData.Proof))
		// For the verification of the SP1 proof system, we consider the ELF file as public input.
		elfLen := (uint32)(len(verifyData.PublicInput))
		verifyResult := sp1.VerifySp1Proof(verifyData.Proof, proofLen, verifyData.PublicInput, elfLen)
		k.Logger().Info("SP1 proof verification:", "result", verifyResult)
		return verifyResult

	default:
		k.Logger().Error("Unrecognized proof system ID")
		return false
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
