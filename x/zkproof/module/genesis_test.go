package zkproof_test

import (
	"testing"

	keepertest "fiamma/testutil/keeper"
	"fiamma/testutil/nullify"
	zkproof "fiamma/x/zkproof/module"
	"fiamma/x/zkproof/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZkproofKeeper(t)
	zkproof.InitGenesis(ctx, k, genesisState)
	got := zkproof.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
