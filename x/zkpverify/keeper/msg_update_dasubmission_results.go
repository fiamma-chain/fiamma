package keeper

import (
	"context"
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
		k.SetDASubmissionResult(ctx, result)
		k.DequeueDASubmission(ctx, result.ProofId)
	}

	return &types.MsgUpdateDASubmissionResultsResponse{}, nil
}
