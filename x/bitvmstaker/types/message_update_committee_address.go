package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateCommitteeAddress{}

func NewMsgUpdateCommitteeAddress(creator string, newCommitteeAddress string) *MsgUpdateCommitteeAddress {
	return &MsgUpdateCommitteeAddress{
		Creator:             creator,
		NewCommitteeAddress: newCommitteeAddress,
	}
}

func (msg *MsgUpdateCommitteeAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.NewCommitteeAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
