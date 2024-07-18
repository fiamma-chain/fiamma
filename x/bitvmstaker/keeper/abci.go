package keeper

import (
	"context"
)

// BeginBlocker will persist the current header and validator set as a historical entry
// and prune the oldest entry based on the HistoricalEntries parameter
func (k *Keeper) BeginBlocker(ctx context.Context) {

}

// Called every block, update validator set
func (k *Keeper) EndBlocker(ctx context.Context) error {
	_, err := k.stakingKeeper.GetAllValidators(ctx)
	if err != nil {
		return err
	}

	// for _, validator := range allValidators {
	// 	valAddr := validator.GetOperator()
	// 	println("EndBlocker")
	// 	println("validator", valAddr)
	// 	_, found := k.GetStaker(ctx, valAddr)
	// 	if !found {
	// 		k.Logger().Info("validator is not a staker", "validator", valAddr)
	// 		validator.GetConsAddr()
	// 		consPubKey, err := validator.Get
	// 		if err != nil { // should never happen
	// 			panic(err)
	// 		}
	// 		println("Jail")
	// 		println("conAddress", sdk.ConsAddress(conAddress).String())
	// 		k.stakingKeeper.Jail(ctx, conAddress)
	// 	}
	// }
	return nil

}
