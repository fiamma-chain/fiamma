package cmd

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

type DAConfig struct {
	RpcURL  string `mapstructure:"rpc"`
	AuthKey string `mapstructure:"authkey"`
}

func defaultFiammaDAConfig() DAConfig {
	return DAConfig{
		RpcURL:  "",
		AuthKey: "",
	}
}

type FiammaAppConfig struct {
	serverconfig.Config `mapstructure:",squash"`

	Wasm wasmtypes.WasmConfig `mapstructure:"wasm"`

	DAConfig DAConfig `mapstructure:"da-config"`
}

func DefaultFiammaConfig() *FiammaAppConfig {
	return &FiammaAppConfig{
		Config:   *serverconfig.DefaultConfig(),
		Wasm:     wasmtypes.DefaultWasmConfig(),
		DAConfig: defaultFiammaDAConfig(),
	}
}

func DefaultFiammaTemplate() string {
	return serverconfig.DefaultConfigTemplate + wasmtypes.DefaultConfigTemplate() + `
###############################################################################
###                      Fiamma DA configuration                      		###
###############################################################################

[da-config]

# RpcURL is the rpc url of the nubit DA node
rpc = "http://127.0.0.1:26658"

# AuthKey is the key to authenticate the nubit DA node
# You can refer to this document for nubit da configuration 
# https://docs.nubit.org/nubit-da/run-a-node
authkey = ""
`
}

func ParseDAOptionsFromConfig(opts servertypes.AppOptions) *DAConfig {
	rpcInterface := opts.Get("da-config.rpc")

	if rpcInterface == nil {
		panic("Nubit Data available rpc config should be in options")
	}

	nubitRPC, err := cast.ToStringE(rpcInterface)

	if err != nil {
		panic("Nubit data available rpc config should be valid string")
	}

	authKeyInterface := opts.Get("da-config.authkey")

	if authKeyInterface == nil {
		panic("Nubit Data available authkey config should be in options")
	}

	nubitAuthKey, err := cast.ToStringE(authKeyInterface)

	if err != nil {
		panic("Nubit data available authkey config should be valid string")
	}

	return &DAConfig{
		RpcURL:  nubitRPC,
		AuthKey: nubitAuthKey,
	}
}
