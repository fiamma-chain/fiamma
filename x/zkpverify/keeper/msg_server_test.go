package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/fiamma-chain/fiamma/testutil/keeper"

	"github.com/stretchr/testify/require"

	"github.com/fiamma-chain/fiamma/x/zkpverify/keeper"
	"github.com/fiamma-chain/fiamma/x/zkpverify/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.ZkpVerifyKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
