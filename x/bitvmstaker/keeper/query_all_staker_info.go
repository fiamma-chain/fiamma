package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllStakerInfo(goCtx context.Context, req *types.QueryAllStakerInfoRequest) (*types.QueryAllStakerInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stakerInfos, pageRes, err := k.GetAllStakerInfo(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStakerInfoResponse{
		AllStakerInfo: stakerInfos,
		Pagination:    pageRes,
	}, nil
}
