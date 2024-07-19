package keeper

import (
	"context"
)

// BeginBlocker will persist the current header and validator set as a historical entry
// and prune the oldest entry based on the HistoricalEntries parameter
func (k *Keeper) BeginBlocker(ctx context.Context) {
}

// EndBlocker Called every block, update validator set
func (k *Keeper) EndBlocker(ctx context.Context) error {
	allValidators, err := k.stakingKeeper.GetAllValidators(ctx)
	if err != nil {
		return err
	}

	for _, validator := range allValidators {
		valAddr := validator.GetOperator()
		if _, found := k.GetStaker(ctx, valAddr); !found {
			conAddress, err := validator.GetConsAddr()
			if err != nil {
				return err
			}
			return k.stakingKeeper.Jail(ctx, conAddress)
		}
	}
	return nil
}
