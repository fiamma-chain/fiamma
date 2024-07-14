package simulation

import (
	"math/rand"

	"fiamma/x/bitvmstaker/keeper"
	"fiamma/x/bitvmstaker/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSlashStaker(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSlashStaker{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SlashStaker simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "SlashStaker simulation not implemented"), nil, nil
	}
}
