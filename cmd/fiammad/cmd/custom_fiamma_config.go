package cmd

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
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
