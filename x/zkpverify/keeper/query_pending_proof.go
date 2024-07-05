package keeper

import (
	"context"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PendingProof(goCtx context.Context, req *types.QueryPendingProofRequest) (*types.QueryPendingProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pendingProofs, err := k.GetPendingProofs(ctx)
	if err != nil {
		return nil, types.ErrPendingProofs
	}

	return &types.QueryPendingProofResponse{PendingProofs: pendingProofs}, nil
}
