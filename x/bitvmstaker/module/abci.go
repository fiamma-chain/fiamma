package bitvmstaker

// EndBlocker checks for any validators that need to be jailed and performs necessary actions
// func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
// 	println("hello-EndBlocker")
// 	// for _, validator := range k.StakingKeeper.GetAllValidators(ctx) {
// 	// 	valAddr := sdk.AccAddress(validator.GetOperator())
// 	// 	_, found := k.GetStaker(ctx, string(valAddr))
// 	// 	if !found {
// 	// 		// Jail the validator
// 	// 		k.StakingKeeper.Jail(ctx, []byte(valAddr.String()))
// 	// 	}
// 	// }
// }
