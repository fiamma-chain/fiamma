package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStaker{}

func NewMsgCreateStaker(creator string, stakerAddress string, registerId uint64) *MsgCreateStaker {
	return &MsgCreateStaker{
		Creator:       creator,
		StakerAddress: stakerAddress,
		RegisterId:    registerId,
	}
}

func (msg *MsgCreateStaker) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.StakerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}
