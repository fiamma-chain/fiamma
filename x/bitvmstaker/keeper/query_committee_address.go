package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CommitteeAddress(goCtx context.Context, req *types.QueryCommitteeAddressRequest) (*types.QueryCommitteeAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	committeeAddress := k.GetCommitteeAddress(ctx)

	return &types.QueryCommitteeAddressResponse{CommitteeAddress: committeeAddress}, nil
}
