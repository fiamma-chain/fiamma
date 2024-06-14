package keeper

import (
	"context"

	"fiamma/x/zkproof/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitGnarkPlonk(goCtx context.Context, msg *types.MsgSubmitGnarkPlonk) (*types.MsgSubmitGnarkPlonkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitGnarkPlonkResponse{}, nil
}
