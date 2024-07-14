package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SlashStaker(goCtx context.Context, msg *types.MsgSlashStaker) (*types.MsgSlashStakerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSlashStakerResponse{}, nil
}
