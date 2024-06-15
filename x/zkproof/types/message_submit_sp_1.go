package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitSp1{}

func NewMsgSubmitSp1(creator string, proofId string, proof string, elf string, metaData string) *MsgSubmitSp1 {
	return &MsgSubmitSp1{
		Creator:  creator,
		ProofId:  proofId,
		Proof:    proof,
		Elf:      elf,
		MetaData: metaData,
	}
}

func (msg *MsgSubmitSp1) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
