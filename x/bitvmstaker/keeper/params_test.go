package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/fiamma-chain/fiamma/testutil/keeper"

	"github.com/fiamma-chain/fiamma/x/bitvmstaker/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BitvmstakerKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
