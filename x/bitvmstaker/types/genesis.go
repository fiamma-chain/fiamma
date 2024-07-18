package types

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CommitteeAddress: "",
		StakerAddresses:  []string{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if gs.CommitteeAddress == "" {
		return fmt.Errorf("committee_address: cannot be empty in genesis file")
	}
	_, err := sdk.AccAddressFromBech32(gs.CommitteeAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	if len(gs.StakerAddresses) == 0 {
		return fmt.Errorf("staker_addresses: cannot be empty in genesis file")
	}

	for _, sa := range gs.StakerAddresses {
		_, err := sdk.ValAddressFromBech32(sa)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
		}
	}

	return gs.Params.Validate()
}
