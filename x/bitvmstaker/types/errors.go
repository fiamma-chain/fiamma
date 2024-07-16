package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/bitvmstaker module sentinel errors
var (
	ErrInvalidSigner  = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrUnauthorized   = sdkerrors.Register(ModuleName, 1101, "msg creator must be a committee address")
	ErrStakerNotFound = sdkerrors.Register(ModuleName, 1102, "staker not found")
)
