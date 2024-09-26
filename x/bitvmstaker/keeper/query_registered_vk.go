package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RegisteredVKList(goCtx context.Context, req *types.QueryRegisteredVKListRequest) (*types.QueryRegisteredVKListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vkList, pageRes, err := k.GetRegisteredVKList(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRegisteredVKListResponse{
		RegisteredVkList: vkList,
		Pagination:       pageRes,
	}, nil
}
