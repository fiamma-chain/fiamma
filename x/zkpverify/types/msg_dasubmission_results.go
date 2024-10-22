package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgUpdateDASubmissionResults{}

func (msg *MsgUpdateDASubmissionResults) ValidateBasic() error {
	if len(msg.DaSubmissionResult) == 0 {
		return ErrNoDASubmissionResults
	}
	for _, result := range msg.DaSubmissionResult {
		if _, ok := DataLocation_name[int32(result.DataLocation)]; !ok {
			return ErrInvalidDataLocation
		}
		if result.ProofId == "" {
			return ErrInvalidProofId
		}
		if result.BlockHash == "" {
			return ErrInvalidBlockHash
		}
		if result.BlockHeight == 0 {
			return ErrInvalidBlockHeight
		}
		if result.Namespace == "" {
			return ErrInvalidNamespace
		}
		if result.TxHash == "" {
			return ErrInvalidTxHash
		}
	}
	return nil
}
