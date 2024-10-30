package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSubmitProof{}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgSubmitProof) ValidateBasic() error {
	// check if the namespace is valid
	if len(msg.Namespace) == 0 {
		return ErrInvalidNamespace
	}

	// check if the proof system is valid
	if _, ok := ProofSystem_value[msg.ProofSystem]; !ok {
		return ErrInvalidProofSystem
	}

	// check if the data location is valid
	if _, ok := DataLocation_value[msg.DataLocation]; !ok {
		return ErrInvalidDataLocation
	}

	// check if the proof is valid
	if len(msg.Proof) == 0 {
		return ErrInvalidProof
	}

	// check if the public input is valid
	if len(msg.PublicInput) == 0 {
		return ErrInvalidPublicInput
	}

	// check if the vk is valid
	if len(msg.Vk) == 0 {
		return ErrInvalidVk
	}

	return nil
}
