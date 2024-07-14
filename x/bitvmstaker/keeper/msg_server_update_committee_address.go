package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateCommitteeAddress(goCtx context.Context, msg *types.MsgUpdateCommitteeAddress) (*types.MsgUpdateCommitteeAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateCommitteeAddressResponse{}, nil
}
