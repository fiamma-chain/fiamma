package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/zkpverify module sentinel errors
var (
	ErrInvalidSigner              = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidProof               = sdkerrors.Register(ModuleName, 1101, "invalid proof")
	ErrInvalidPublicInput         = sdkerrors.Register(ModuleName, 1102, "invalid public input")
	ErrInvalidVk                  = sdkerrors.Register(ModuleName, 1103, "invalid vk")
	ErrInvalidNamespace           = sdkerrors.Register(ModuleName, 1104, "invalid namespace")
	ErrSubmitProof                = sdkerrors.Register(ModuleName, 1105, "error submitting proof to DA")
	ErrInvalidProofSystem         = sdkerrors.Register(ModuleName, 1106, "invalid proof system")
	ErrProofDataNotFound          = sdkerrors.Register(ModuleName, 1107, "proof data not found")
	ErrBitVMChallengeDataNotFound = sdkerrors.Register(ModuleName, 1108, "bitvm witness not found")
	ErrGetProofId                 = sdkerrors.Register(ModuleName, 1109, "error getting proof id")
	ErrInvalidProofId             = sdkerrors.Register(ModuleName, 1110, "invalid proof id")
	ErrVerifyPeriod               = sdkerrors.Register(ModuleName, 1111, "err exceeding verification period")
	ErrVerifyResult               = sdkerrors.Register(ModuleName, 1112, "error verifying proof result")
	ErrPendingProofs              = sdkerrors.Register(ModuleName, 1113, "none pending proofs")
	ErrVerifyResultNotFound       = sdkerrors.Register(ModuleName, 1114, "verify result not found")
	ErrInvalidVerifyResult        = sdkerrors.Register(ModuleName, 1115, "invalid verify result")
	ErrProofNotPending            = sdkerrors.Register(ModuleName, 1116, "proof not pending")
)
