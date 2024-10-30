package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/fiamma-chain/fiamma/x/zkpverify/types"
)

func (k Keeper) DASubmitter(goCtx context.Context, req *types.QueryDASubmitterRequest) (*types.QueryDASubmitterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	daSubmitter := k.GetDASubmitter(ctx)

	return &types.QueryDASubmitterResponse{DaSubmitter: daSubmitter}, nil
}
