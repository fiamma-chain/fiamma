package keeper

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"strconv"

	"fiamma/x/zkpverify/types"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/plonk"
	"github.com/consensys/gnark/backend/witness"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitGnarkPlonk(goCtx context.Context, msg *types.MsgSubmitGnarkPlonk) (*types.MsgSubmitGnarkPlonkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proofBytes, err := base64.StdEncoding.DecodeString(msg.Proof)
	if err != nil {
		return nil, types.ErrInvalidVerifyData
	}
	publicInputBytes, err := base64.StdEncoding.DecodeString(msg.PublicInputs)

	if err != nil {
		return nil, types.ErrInvalidVerifyData
	}

	verificationKeyBytes, err := base64.StdEncoding.DecodeString(msg.VerifyingKey)

	if err != nil {
		return nil, types.ErrInvalidVerifyData

	}
	verfiyData := types.VerificationData{
		ProvingTypeId:   types.GnarkPlonk,
		Proof:           proofBytes,
		PubInput:        publicInputBytes,
		VerificationKey: verificationKeyBytes,
	}
	// submit proof data to DA
	proofCommitment, proofDAIds, err := k.SubmitVerifyDataToDA(ctx, verfiyData)
	if err != nil {
		k.Logger().Info("Error submitting proof to DA: %v", err)
		return nil, types.ErrSubmitProof
	}

	// The chain first verifies the correctness of the proofs submitted by the user, and saves the results.
	// The observer may challenge the result at a later stage.
	result := k.verifyGnarkPlonk(proofBytes, publicInputBytes, verificationKeyBytes)

	k.logger.Info("Proof verification result: %v", result)

	// store verify data in the store
	verifyData := types.ZkpVerify{
		Creator:     msg.Creator,
		ProofId:     hex.EncodeToString(proofDAIds[0]),
		ProofType:   "PLONK",
		ProofStatus: 1,
	}

	k.SetVerifyData(ctx, hex.EncodeToString(proofCommitment[:]), verifyData)

	event := sdk.NewEvent("verification_finished",
		sdk.NewAttribute("verify_result", strconv.FormatBool(result)),
		sdk.NewAttribute("prover_type", "PLONK"))
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitGnarkPlonkResponse{}, nil
}

func (k msgServer) verifyGnarkPlonk(proofBytes []byte, publicInputBytes []byte, verifyKeyBytes []byte) bool {
	proofReader := bytes.NewReader(proofBytes)
	proof := plonk.NewProof(ecc.BN254)

	if _, err := proof.ReadFrom(proofReader); err != nil {
		k.Logger().Info("Could not deserialize proof: %v", err)
		return false
	}

	publicInputsReader := bytes.NewReader(publicInputBytes)
	publicInput, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		k.Logger().Info("Error instantiating witness: %v", err)
		return false
	}

	if _, err = publicInput.ReadFrom(publicInputsReader); err != nil {
		k.Logger().Info("Could not read PLONK public input: %v", err)
		return false
	}

	verificationKeyReader := bytes.NewReader(verifyKeyBytes)
	verifyingkey := plonk.NewVerifyingKey(ecc.BN254)
	if _, err := verifyingkey.ReadFrom(verificationKeyReader); err != nil {
		k.Logger().Info("Could not read plonk verifying key from bytes: %v", err)
		return false
	}

	err = plonk.Verify(proof, verifyingkey, publicInput)

	return err == nil
}
