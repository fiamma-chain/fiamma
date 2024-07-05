package keeper

import (
	"context"
	"encoding/hex"
	"strconv"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitProof(goCtx context.Context, msg *types.MsgSubmitProof) (*types.MsgSubmitProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proofSystemId, err := types.ProofSystemIdFromString(msg.ProofSystem)
	if err != nil {
		k.Logger().Info("Error parsing proof system:", "error", err)
		return nil, types.ErrInvalidProofSystem

	}
	proofData := types.ProofData{
		ProofSystem: uint64(proofSystemId),
		Proof:       msg.Proof,
		PublicInput: msg.PublicInput,
		Vk:          msg.Vk,
	}

	proofId, err := k.GetProofId(proofData)
	if err != nil {
		k.Logger().Info("Error getting proof id:", "error", err)
		return nil, types.ErrGetProofId

	}

	// submit proof data to DA
	dataCommitmentStr, dataLocationId, err := k.SubmitProofData(ctx, proofId[:], proofData)
	if err != nil {
		k.Logger().Info("Error submitting proof to DA:", "error", err)
		return nil, types.ErrSubmitProof
	}

	proofIdStr := hex.EncodeToString(proofId[:])

	// The chain first verifies the correctness of the proofs submitted by the user, and saves the results.
	// The observer may challenge the result at a later stage.
	result, witness := k.verifyProof(&proofData)

	// store witness if the proof system is BitVM
	if proofData.ProofSystem == uint64(types.Groth16Bn254_BitVM) {
		k.SetBitVMWitness(ctx, proofId[:], witness)
	}

	k.Logger().Info("Proof verification:", "result", result)

	// store verify data in the store
	verifyResult := types.VerifyResult{
		ProofId:                    proofIdStr,
		DataCommitment:             dataCommitmentStr,
		DataLocation:               uint64(dataLocationId),
		Result:                     result,
		Status:                     types.VerificationStatus_INITIALVALIDATION,
		CommunityVerificationCount: uint64(0),
	}

	k.SetVerifyResult(ctx, proofId[:], verifyResult)

	event := sdk.NewEvent("InitVerifyFinished",
		sdk.NewAttribute("proofId", proofIdStr),
		sdk.NewAttribute("dataCommitment", dataCommitmentStr),
		sdk.NewAttribute("dataLocation", dataLocationId.String()),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)),
		sdk.NewAttribute("proofSystem", msg.ProofSystem))
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitProofResponse{}, nil
}
