package bitvmstaker_test

import (
	"testing"

	keepertest "github.com/fiamma-chain/fiamma/testutil/keeper"
	"github.com/fiamma-chain/fiamma/testutil/nullify"
	"github.com/fiamma-chain/fiamma/testutil/sample"

	bitvmstaker "github.com/fiamma-chain/fiamma/x/bitvmstaker/module"
	"github.com/fiamma-chain/fiamma/x/bitvmstaker/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:           types.DefaultParams(),
		CommitteeAddress: sample.AccAddress(),
		StakerAddresses:  []string{sample.ValAddress()},

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BitvmstakerKeeper(t)
	bitvmstaker.InitGenesis(ctx, k, genesisState)
	got := bitvmstaker.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
