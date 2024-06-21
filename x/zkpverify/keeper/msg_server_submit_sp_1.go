package keeper

import (
	"context"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitSp1(goCtx context.Context, msg *types.MsgSubmitSp1) (*types.MsgSubmitSp1Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitSp1Response{}, nil
}
