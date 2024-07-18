package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker will persist the current header and validator set as a historical entry
// and prune the oldest entry based on the HistoricalEntries parameter
func (k *Keeper) BeginBlocker(ctx context.Context) {

}

// Called every block, update validator set
func (k *Keeper) EndBlocker(ctx context.Context) error {
	allValidators, err := k.stakingKeeper.GetAllValidators(ctx)
	if err != nil {
		return err
	}

	for _, validator := range allValidators {
		valAddr := validator.GetOperator()
		_, found := k.GetStaker(ctx, valAddr)
		if !found {
			k.Logger().Info("validator is not a bitvm staker", "validator", valAddr)
			conAddress, err := validator.GetConsAddr()
			if err != nil { // should never happen
				k.Logger().Error("failed to get consensus address", "validator", valAddr)
				return err
			}
			k.stakingKeeper.Jail(ctx, types.ConsAddress(conAddress))
		}
	}
	return nil

}
