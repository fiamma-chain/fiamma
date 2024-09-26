package keeper

import (
	"context"
	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveVK(goCtx context.Context, msg *types.MsgRemoveVK) (*types.MsgRemoveVKResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	committeeAddress := k.GetCommitteeAddress(ctx)

	if msg.Creator != committeeAddress {
		return nil, types.ErrUnauthorized
	}

	err := k.Keeper.RemoveVK(ctx, msg.Vk)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemoveVKResponse{}, nil
}
