package keeper_test

import (
	keppertest "fiamma/testutil/keeper"
	"fiamma/x/bitvmstaker/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesis := types.GenesisState{
		Params:           types.Params{},
		CommitteeAddress: "testAddress",
	}
	k, ctx := keppertest.BitvmstakerKeeper(t)
	k.InitGenesis(ctx, genesis)
	got := k.ExportGenesis(ctx)
	require.NotNil(t, got)
	require.Equal(t, got.CommitteeAddress, "testAddress")
}
