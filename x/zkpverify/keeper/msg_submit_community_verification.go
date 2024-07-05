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

	verifyResult, found := k.GetVerifyResult(ctx, proofId[:])
	if !found {
		k.Logger().Info("Error finding proof id:", "error", msg.ProofId)
		return nil, types.ErrGetProofId
	}

	if verifyResult.Result != msg.VerifyResult {
		k.Logger().Info("Inconsistent with verification result:", "error", msg.ProofId)
		return nil, types.ErrVerifyResult
	}

	if verifyResult.Status == types.VerificationStatus_DEFINITIVEVALIDATION {
		k.Logger().Info("Error exceeding verification period:", "error", msg.ProofId)
		return nil, types.ErrVerifyPeriod
	}

	k.Logger().Info("Proof verification result for community:", "result", msg.VerifyResult)

	verifyResult.CommunityVerificationCount++
	if verifyResult.CommunityVerificationCount < VerificationCountLimit {
		verifyResult.Status = types.VerificationStatus_COMMUNITYVALIDATION
	} else {
		verifyResult.Status = types.VerificationStatus_DEFINITIVEVALIDATION
	}
	k.Logger().Info("Proof verification status for community:", "status", verifyResult.Status)

	k.SetVerifyResult(ctx, proofId[:], verifyResult)

	event := sdk.NewEvent("verifyFinished",
		sdk.NewAttribute("proofId", msg.ProofId),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(msg.VerifyResult)),
		sdk.NewAttribute("status", strconv.Itoa(int(verifyResult.Status))),
		sdk.NewAttribute("CommunityVerificationCount", strconv.Itoa(int(verifyResult.CommunityVerificationCount))))
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitCommunityVerificationResponse{}, nil
}
