package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker will persist the current header and validator set as a historical entry
// and prune the oldest entry based on the HistoricalEntries parameter
func (k *Keeper) BeginBlocker(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	proposer := sdkCtx.BlockHeader().ProposerAddress
	validator, err := k.stakingKeeper.ValidatorByConsAddr(sdkCtx, proposer)
	if err != nil {
		return err
	}
	operatorAddr := validator.GetOperator()
	k.SetBlockProposer(ctx, sdkCtx.BlockHeight(), operatorAddr)
	return nil
}

// EndBlocker Called every block, update validator set
func (k *Keeper) EndBlocker(ctx context.Context) error {
	return nil
}
