package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateStaker(goCtx context.Context, msg *types.MsgCreateStaker) (*types.MsgCreateStakerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateStakerResponse{}, nil
}
