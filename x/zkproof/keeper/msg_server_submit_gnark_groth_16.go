package keeper

import (
	"context"

	"fiamma/x/zkproof/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitGnarkGroth16(goCtx context.Context, msg *types.MsgSubmitGnarkGroth16) (*types.MsgSubmitGnarkGroth16Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitGnarkGroth16Response{}, nil
}
