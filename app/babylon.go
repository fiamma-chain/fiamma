package app

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/babylonchain/babylon-sdk/x/babylon"
	bbnkeeper "github.com/babylonchain/babylon-sdk/x/babylon/keeper"
	bbntypes "github.com/babylonchain/babylon-sdk/x/babylon/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// registerBabylonModules register Babylon keepers and non dependency inject modules.
func (app *App) registerBabylonModules() error {
	// set up non depinject support modules store keys
	if err := app.RegisterStores(
		storetypes.NewKVStoreKey(bbntypes.StoreKey),
		storetypes.NewMemoryStoreKey(bbntypes.MemStoreKey),
	); err != nil {
		return err
	}

	app.BabylonKeeper = bbnkeeper.NewKeeper(
		app.appCodec,
		app.GetKey(bbntypes.StoreKey),
		app.GetMemKey(bbntypes.MemStoreKey),
		app.BankKeeper,
		app.StakingKeeper,
		&app.WasmKeeper, // ensure this is a pointer as we instantiate the keeper a bit later
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	// register babylon modules
	if err := app.RegisterModules(
		babylon.NewAppModule(app.appCodec, app.BabylonKeeper),
	); err != nil {
		return err
	}

	return nil
}
