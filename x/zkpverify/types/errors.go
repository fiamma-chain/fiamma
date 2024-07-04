package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/zkpverify module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidVerifyData    = sdkerrors.Register(ModuleName, 1101, "invalid proof verify data")
	ErrSubmitProof          = sdkerrors.Register(ModuleName, 1102, "error submitting proof to DA")
	ErrInvalidProofSystem   = sdkerrors.Register(ModuleName, 1103, "invalid proof system")
	ErrProofDataNotFound    = sdkerrors.Register(ModuleName, 1104, "proof data not found")
	ErrBitVMWitnessNotFound = sdkerrors.Register(ModuleName, 1105, "bitvm witness not found")
	ErrGetVerifyId          = sdkerrors.Register(ModuleName, 1106, "error getting verifyId")
)
