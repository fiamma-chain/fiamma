package keeper

import (
	"context"
	"encoding/hex"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateDASubmissionResults(goCtx context.Context, msg *types.MsgUpdateDASubmissionResults) (*types.MsgUpdateDASubmissionResultsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgUpdateDASubmissionResultsResponse{}, err
	}

	daSubmitter := k.GetDASubmitter(ctx)
	if msg.Creator != daSubmitter {
		return nil, types.ErrMsgCreatorMustBeDASubmitter
	}

	for _, result := range msg.DaSubmissionResult {
		proofId, err := hex.DecodeString(result.ProofId)
		if err != nil {
			k.Logger().Info("Error decoding proof id:", "error", err)
			return nil, types.ErrInvalidProofId
		}
		k.SetDASubmissionResult(ctx, proofId[:], result)
		k.DequeueDASubmission(ctx, proofId[:])
	}

	return &types.MsgUpdateDASubmissionResultsResponse{}, nil
}
