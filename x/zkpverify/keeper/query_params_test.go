package keeper_test

import (
	"testing"

	keepertest "github.com/fiamma-chain/fiamma/testutil/keeper"

	"github.com/stretchr/testify/require"

	"github.com/fiamma-chain/fiamma/x/zkpverify/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.ZkpVerifyKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
