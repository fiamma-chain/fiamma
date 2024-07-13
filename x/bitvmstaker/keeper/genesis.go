package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, gs types.GenesisState) error {
	k.SetParams(ctx, gs.Params)
	k.SetCommitteeAddress(ctx, gs.CommitteeAddress)

	return nil
}

// ExportGenesis returns the module's exported genesis
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:           k.GetParams(ctx),
		CommitteeAddress: k.GetCommitteeAddress(ctx),
	}
}
