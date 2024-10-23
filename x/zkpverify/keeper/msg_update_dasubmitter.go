package keeper

import (
	"context"

	"fiamma/x/zkpverify/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateDASubmitter(goCtx context.Context, msg *types.MsgUpdateDASubmitter) (*types.MsgUpdateDASubmitterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.ValAddressFromBech32(msg.DaSubmitter)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	daSubmitter := k.GetDASubmitter(ctx)

	if msg.Creator != daSubmitter {
		return nil, types.ErrMsgCreatorMustBeDASubmitter
	}

	k.SetDASubmitter(ctx, msg.DaSubmitter)

	return &types.MsgUpdateDASubmitterResponse{}, nil
}
