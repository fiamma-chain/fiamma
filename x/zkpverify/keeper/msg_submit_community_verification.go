package keeper

import (
	"context"
	"encoding/hex"
	"strconv"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	VerificationCountLimit uint64 = 4
)

func (k msgServer) SubmitCommunityVerification(goCtx context.Context, msg *types.MsgSubmitCommunityVerification) (*types.MsgSubmitCommunityVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proofId, err := hex.DecodeString(msg.ProofId)
	if err != nil {
		k.Logger().Info("Error decoding proof id:", "error", err)
		return nil, types.ErrInvalidProofId
	}

	verifyResult, foundPending := k.GetPendingProof(ctx, proofId[:])
	if !foundPending {
		k.Logger().Info("Error finding proof id:", "error", msg.ProofId)
		return nil, types.ErrGetProofId
	}

	if verifyResult.Result != msg.VerifyResult {
		k.Logger().Info("Inconsistent with verification result:", "error", msg.ProofId)
		return nil, types.ErrVerifyResult
	}

	k.Logger().Info("Proof verification result for community:", "result", msg.VerifyResult)

	verifyResult.CommunityVerificationCount++
	if verifyResult.CommunityVerificationCount < VerificationCountLimit {
		verifyResult.Status = types.VerificationStatus_COMMUNITY_VALIDATION
		k.SetPendingProof(ctx, proofId, verifyResult)
	} else {
		verifyResult.Status = types.VerificationStatus_DEFINITIVE_VALIDATION
		k.DeletePendingProof(ctx, proofId)
	}
	k.SetVerifyResult(ctx, proofId, verifyResult)
	k.Logger().Info("Proof verification status for community:", "status", verifyResult.Status)

	event := sdk.NewEvent("SubmitCommunityVerification",
		sdk.NewAttribute("proofId", msg.ProofId),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(msg.VerifyResult)),
		sdk.NewAttribute("status", verifyResult.Status.String()),
		sdk.NewAttribute("CommunityVerificationCount", strconv.Itoa(int(verifyResult.CommunityVerificationCount))))
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitCommunityVerificationResponse{}, nil
}
