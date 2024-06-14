package simulation

import (
	"math/rand"

	"fiamma/x/zkproof/keeper"
	"fiamma/x/zkproof/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSubmitGnarkPlonk(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSubmitGnarkPlonk{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SubmitGnarkPlonk simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "SubmitGnarkPlonk simulation not implemented"), nil, nil
	}
}
