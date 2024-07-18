package keeper_test

import (
	"testing"

	keppertest "fiamma/testutil/keeper"
	"fiamma/testutil/sample"
	"fiamma/x/bitvmstaker/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesis := types.GenesisState{
		Params:           types.Params{},
		CommitteeAddress: sample.AccAddress(),
	}
	k, ctx := keppertest.BitvmstakerKeeper(t)
	err := k.InitGenesis(ctx, genesis)
	require.NoError(t, err)
	got := k.ExportGenesis(ctx)
	require.NotNil(t, got)
}
