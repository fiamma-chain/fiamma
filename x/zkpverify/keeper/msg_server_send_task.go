package keeper

import (
	"context"
	"encoding/hex"
	"strconv"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendTask(goCtx context.Context, msg *types.MsgSendTask) (*types.MsgSendTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proofSystemId, err := types.ProofSystemIdFromString(msg.ProofSystem)
	if err != nil {
		k.Logger().Info("Error parsing proof system:", "error", err)
		return nil, types.ErrInvalidProofSystem

	}
	verfiyData := types.VerifyData{
		ProofSystem: uint64(proofSystemId),
		Proof:       msg.Proof,
		PublicInput: msg.PublicInput,
		Vk:          msg.Vk,
	}
	// submit proof data to DA
	verifyId, dataCommitments, dataLocationId, err := k.SubmitVerifyData(ctx, verfiyData)
	if err != nil {
		k.Logger().Info("Error submitting proof to DA:", "error", err)
		return nil, types.ErrSubmitProof
	}

	verifyIdStr := hex.EncodeToString(verifyId[:])
	dataCommitmentStr := hex.EncodeToString(dataCommitments[0])

	// The chain first verifies the correctness of the proofs submitted by the user, and saves the results.
	// The observer may challenge the result at a later stage.
	result := k.verifyProof(&verfiyData)

	k.Logger().Info("Proof verification:", "result", result)

	// store verify data in the store
	verifyResult := types.VerifyResult{
		VerifyId:       verifyIdStr,
		DataCommitment: dataCommitmentStr,
		DataLocation:   uint64(dataLocationId),
		Result:         result,
	}

	k.SetVerifyResult(ctx, verifyId[:], verifyResult)

	event := sdk.NewEvent("verifyFinished",
		sdk.NewAttribute("verifyId", verifyIdStr),
		sdk.NewAttribute("dataCommitment", dataCommitmentStr),
		sdk.NewAttribute("dataLocation", dataLocationId.String()),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)),
		sdk.NewAttribute("proofSystem", msg.ProofSystem))
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSendTaskResponse{}, nil
}
