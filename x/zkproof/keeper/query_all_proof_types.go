package keeper

import (
	"context"

	"fiamma/x/zkproof/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllProofTypes(goCtx context.Context, req *types.QueryAllProofTypesRequest) (*types.QueryAllProofTypesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryAllProofTypesResponse{}, nil
}
