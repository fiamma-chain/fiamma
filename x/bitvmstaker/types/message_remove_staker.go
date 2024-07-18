package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveStaker{}

func NewMsgRemoveStaker(creator string, stakerAddress string) *MsgRemoveStaker {
	return &MsgRemoveStaker{
		Creator:       creator,
		StakerAddress: stakerAddress,
	}
}

func (msg *MsgRemoveStaker) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}
