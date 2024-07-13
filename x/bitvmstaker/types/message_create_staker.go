package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStaker{}

func NewMsgCreateStaker(creator string, stakerAddress string) *MsgCreateStaker {
	return &MsgCreateStaker{
		Creator:       creator,
		StakerAddress: stakerAddress,
	}
}

func (msg *MsgCreateStaker) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
