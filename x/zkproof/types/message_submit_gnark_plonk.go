package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitGnarkPlonk{}

func NewMsgSubmitGnarkPlonk(creator string, proofId string, proof string, publicInputs string, verifyingKey string, metaData string) *MsgSubmitGnarkPlonk {
	return &MsgSubmitGnarkPlonk{
		Creator:      creator,
		ProofId:      proofId,
		Proof:        proof,
		PublicInputs: publicInputs,
		VerifyingKey: verifyingKey,
		MetaData:     metaData,
	}
}

func (msg *MsgSubmitGnarkPlonk) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
