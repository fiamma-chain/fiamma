package keeper

import (
	"context"

	"fiamma/x/zkproof/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PendingProofByType(goCtx context.Context, req *types.QueryPendingProofByTypeRequest) (*types.QueryPendingProofByTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryPendingProofByTypeResponse{}, nil
}
