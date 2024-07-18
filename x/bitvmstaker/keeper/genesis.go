package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, gs types.GenesisState) error {
	if err := gs.Validate(); err != nil {
		panic(err)
	}
	err := k.SetParams(ctx, gs.Params)
	if err != nil {
		return err
	}
	k.SetCommitteeAddress(ctx, gs.CommitteeAddress)

	for _, staker := range gs.StakerAddresses {
		k.AppendStaker(ctx, types.StakerInfo{StakerAddress: staker})
	}

	return nil
}

// ExportGenesis returns the module's exported genesis
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:           k.GetParams(ctx),
		CommitteeAddress: k.GetCommitteeAddress(ctx),
	}
}
