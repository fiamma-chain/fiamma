package keeper

import (
	"context"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendTask(goCtx context.Context, msg *types.MsgSendTask) (*types.MsgSendTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendTaskResponse{}, nil
}
