package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/zkpverify module sentinel errors
var (
	ErrInvalidSigner               = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidProof                = sdkerrors.Register(ModuleName, 1101, "invalid proof")
	ErrInvalidPublicInput          = sdkerrors.Register(ModuleName, 1102, "invalid public input")
	ErrInvalidVk                   = sdkerrors.Register(ModuleName, 1103, "invalid vk")
	ErrInvalidNamespace            = sdkerrors.Register(ModuleName, 1104, "invalid namespace")
	ErrSubmitProof                 = sdkerrors.Register(ModuleName, 1105, "error submitting proof to DA")
	ErrInvalidProofSystem          = sdkerrors.Register(ModuleName, 1106, "invalid proof system")
	ErrInvalidDataLocation         = sdkerrors.Register(ModuleName, 1107, "invalid data location")
	ErrProofDataNotFound           = sdkerrors.Register(ModuleName, 1108, "proof data not found")
	ErrBitVMChallengeDataNotFound  = sdkerrors.Register(ModuleName, 1109, "bitvm witness not found")
	ErrGetProofId                  = sdkerrors.Register(ModuleName, 1110, "error getting proof id")
	ErrInvalidProofId              = sdkerrors.Register(ModuleName, 1111, "invalid proof id")
	ErrVerifyPeriod                = sdkerrors.Register(ModuleName, 1112, "err exceeding verification period")
	ErrVerifyResult                = sdkerrors.Register(ModuleName, 1113, "error verifying proof result")
	ErrPendingProofs               = sdkerrors.Register(ModuleName, 1114, "none pending proofs")
	ErrVerifyResultNotFound        = sdkerrors.Register(ModuleName, 1115, "verify result not found")
	ErrInvalidVerifyResult         = sdkerrors.Register(ModuleName, 1116, "invalid verify result")
	ErrProofNotPending             = sdkerrors.Register(ModuleName, 1117, "proof not pending")
	ErrVKNotRegistered             = sdkerrors.Register(ModuleName, 1118, "vk in the proof system has not been registered in the committee.")
	ErrMsgCreatorMustBeDASubmitter = sdkerrors.Register(ModuleName, 1119, "msg creator must be a da submitter")
	ErrNoDASubmissionResults       = sdkerrors.Register(ModuleName, 1120, "no da submission results")
	ErrInvalidBlockHash            = sdkerrors.Register(ModuleName, 1121, "invalid block hash")
	ErrInvalidBlockHeight          = sdkerrors.Register(ModuleName, 1122, "invalid block height")
	ErrInvalidTxHash               = sdkerrors.Register(ModuleName, 1123, "invalid tx hash")
)
