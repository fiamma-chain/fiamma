package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateCommitteeAddress(goCtx context.Context, msg *types.MsgUpdateCommitteeAddress) (*types.MsgUpdateCommitteeAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	committeeAddress := k.GetCommitteeAddress(ctx)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgUpdateCommitteeAddressResponse{}, err
	}

	if msg.Creator != committeeAddress {
		return nil, types.ErrUnauthorized
	}

	k.SetCommitteeAddress(ctx, msg.NewCommitteeAddress)

	return &types.MsgUpdateCommitteeAddressResponse{}, nil
}
