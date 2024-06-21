package zkpverify_test

import (
	"testing"

	keepertest "fiamma/testutil/keeper"
	"fiamma/testutil/nullify"
	zkpverify "fiamma/x/zkpverify/module"
	"fiamma/x/zkpverify/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZkpVerifyKeeper(t)
	zkpverify.InitGenesis(ctx, k, genesisState)
	got := zkpverify.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
