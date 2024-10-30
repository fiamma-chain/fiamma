package keeper

import (
	"github.com/fiamma-chain/fiamma/x/zkpverify/types"
)

var _ types.QueryServer = Keeper{}
