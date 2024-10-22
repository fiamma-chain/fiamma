package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"fiamma/x/zkpverify/types"
)

func (k Keeper) DASubmissionQueue(goCtx context.Context, req *types.QueryDASubmissionQueueRequest) (*types.QueryDASubmissionQueueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	queue, pageRes, err := k.GetDASubmissionQueue(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &types.QueryDASubmissionQueueResponse{DaSubmissionData: queue, Pagination: pageRes}, nil
}
