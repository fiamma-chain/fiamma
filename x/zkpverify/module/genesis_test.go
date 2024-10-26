package zkpverify_test

import (
	"testing"

	zkpverify "github.com/fiamma-chain/fiamma/x/zkpverify/module"
	"github.com/fiamma-chain/fiamma/x/zkpverify/types"

	keepertest "github.com/fiamma-chain/fiamma/testutil/keeper"
	"github.com/fiamma-chain/fiamma/testutil/nullify"
	"github.com/fiamma-chain/fiamma/testutil/sample"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:      types.DefaultParams(),
		DaSubmitter: sample.AccAddress(),
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
