package keeper

import (
	"context"

	"fiamma/x/zkpverify/types"
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
	k.SetDASubmitter(ctx, gs.DaSubmitter)

	return nil
}

// ExportGenesis returns the module's exported genesis
func (k Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:      k.GetParams(ctx),
		DaSubmitter: k.GetDASubmitter(ctx),
	}
}
