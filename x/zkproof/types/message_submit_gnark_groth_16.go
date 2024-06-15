package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitGnarkGroth16{}

func NewMsgSubmitGnarkGroth16(creator string, proof string, publicInputs string, verifyingKey string, metaData string) *MsgSubmitGnarkGroth16 {
	return &MsgSubmitGnarkGroth16{
		Creator:      creator,
		Proof:        proof,
		PublicInputs: publicInputs,
		VerifyingKey: verifyingKey,
		MetaData:     metaData,
	}
}

func (msg *MsgSubmitGnarkGroth16) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
