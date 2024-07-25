package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateStaker(goCtx context.Context, msg *types.MsgCreateStaker) (*types.MsgCreateStakerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgCreateStakerResponse{}, err
	}

	if msg.Creator != k.GetCommitteeAddress(ctx) {
		return nil, types.ErrUnauthorized
	}
	stakerInfo := types.StakerInfo{
		StakerAddress:    msg.StakerAddress,
		StakerRegisterId: msg.RegisterId,
	}
	k.AppendStaker(ctx, stakerInfo)

	return &types.MsgCreateStakerResponse{}, nil
}
