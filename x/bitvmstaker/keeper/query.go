package keeper

import (
	"github.com/fiamma-chain/fiamma/x/bitvmstaker/types"
)

var _ types.QueryServer = Keeper{}
