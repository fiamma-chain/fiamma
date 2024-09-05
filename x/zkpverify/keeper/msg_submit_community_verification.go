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

	// check if the proof id is valid
	if len(msg.ProofId) == 0 {
		return nil, types.ErrInvalidProofId
	}

	proofId, err := hex.DecodeString(msg.ProofId)
	if err != nil {
		k.Logger().Info("Error decoding proof id:", "error", err)
		return nil, types.ErrInvalidProofId
	}

	verifyResult, found := k.GetVerifyResult(ctx, proofId[:])
	if !found {
		k.Logger().Info("Error finding proof id:", "error", msg.ProofId)
		return nil, types.ErrGetProofId
	}

	// Check if the proof is still pending
	if verifyResult.Status != types.VerificationStatus_INITIAL_VALIDATION &&
		verifyResult.Status != types.VerificationStatus_COMMUNITY_VALIDATION {
		k.Logger().Info("Proof is not in pending status:", "proofId", msg.ProofId, "status", verifyResult.Status)
		return nil, types.ErrProofNotPending
	}

	if verifyResult.Result != msg.VerifyResult {
		k.Logger().Info("Inconsistent with verification result:", "error", msg.ProofId)
		return nil, types.ErrVerifyResult
	}

	verifyResult.CommunityVerificationCount++
	if verifyResult.CommunityVerificationCount < VerificationCountLimit {
		verifyResult.Status = types.VerificationStatus_COMMUNITY_VALIDATION
	} else {
		verifyResult.Status = types.VerificationStatus_DEFINITIVE_VALIDATION
		k.RemovePendingProofIndex(ctx, proofId)
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
