package keeper_test

import (
	"testing"

	keepertest "fiamma/testutil/keeper"

	"github.com/stretchr/testify/require"

	"fiamma/x/zkpverify/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ZkpVerifyKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
