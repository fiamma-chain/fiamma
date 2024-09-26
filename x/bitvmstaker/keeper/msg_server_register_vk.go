package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterVK(goCtx context.Context, msg *types.MsgRegisterVK) (*types.MsgRegisterVKResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	committeeAddress := k.GetCommitteeAddress(ctx)

	if msg.Creator != committeeAddress {
		return nil, types.ErrUnauthorized
	}
	// Register the VK
	err := k.Keeper.RegisterVK(ctx, msg.Vk)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterVKResponse{}, nil
}
