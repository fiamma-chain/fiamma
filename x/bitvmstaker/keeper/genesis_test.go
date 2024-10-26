package keeper_test

import (
	"testing"

	keppertest "github.com/fiamma-chain/fiamma/testutil/keeper"
	"github.com/fiamma-chain/fiamma/testutil/sample"

	"github.com/fiamma-chain/fiamma/x/bitvmstaker/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesis := types.GenesisState{
		Params:           types.Params{},
		CommitteeAddress: sample.AccAddress(),
		StakerAddresses:  []string{sample.ValAddress()},
	}
	k, ctx := keppertest.BitvmstakerKeeper(t)
	err := k.InitGenesis(ctx, genesis)
	require.NoError(t, err)
	got := k.ExportGenesis(ctx)
	require.NotNil(t, got)
}
