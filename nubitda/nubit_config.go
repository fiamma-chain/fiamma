package nubitda

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

type DAConfig struct {
	RpcURL  string `json:"rpcURL"`
	AuthKey string `json:"authKey"`
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
