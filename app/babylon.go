package app

import (
	"cosmossdk.io/core/appmodule"
	storetypes "cosmossdk.io/store/types"
	"github.com/babylonchain/babylon-sdk/x/babylon"
	bbnkeeper "github.com/babylonchain/babylon-sdk/x/babylon/keeper"
	bbntypes "github.com/babylonchain/babylon-sdk/x/babylon/types"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
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

// RegisterBabylon registers the Babylon module and its interfaces with the provided
// interface registry. It returns a map of module names to their corresponding
// AppModule implementations.
func RegisterBabylon(registry cdctypes.InterfaceRegistry) map[string]appmodule.AppModule {
	modules := map[string]appmodule.AppModule{
		bbntypes.ModuleName: babylon.AppModule{},
	}
	for name, m := range modules {
		module.CoreAppModuleBasicAdaptor(name, m).RegisterInterfaces(registry)
	}
	return modules
}
