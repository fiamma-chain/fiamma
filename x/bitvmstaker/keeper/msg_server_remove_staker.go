package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RemoveStaker(goCtx context.Context, msg *types.MsgRemoveStaker) (*types.MsgRemoveStakerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgRemoveStakerResponse{}, err
	}

	if msg.Creator != k.GetCommitteeAddress(ctx) {
		return nil, types.ErrUnauthorized
	}

	stakerInfo, found := k.GetStaker(ctx, msg.StakerAddress)

	if !found {
		return nil, types.ErrStakerNotFound
	}

	k.DeleteStaker(ctx, stakerInfo.StakerAddress)

	return &types.MsgRemoveStakerResponse{}, nil
}
