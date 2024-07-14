package bitvmstaker

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"fiamma/testutil/sample"
	bitvmstakersimulation "fiamma/x/bitvmstaker/simulation"
	"fiamma/x/bitvmstaker/types"
)

// avoid unused import issue
var (
	_ = bitvmstakersimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateStaker = "op_weight_msg_create_staker"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStaker int = 100

	opWeightMsgSlashStaker = "op_weight_msg_slash_staker"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSlashStaker int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bitvmstakerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bitvmstakerGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateStaker int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStaker, &weightMsgCreateStaker, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStaker = defaultWeightMsgCreateStaker
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStaker,
		bitvmstakersimulation.SimulateMsgCreateStaker(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSlashStaker int
	simState.AppParams.GetOrGenerate(opWeightMsgSlashStaker, &weightMsgSlashStaker, nil,
		func(_ *rand.Rand) {
			weightMsgSlashStaker = defaultWeightMsgSlashStaker
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSlashStaker,
		bitvmstakersimulation.SimulateMsgSlashStaker(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStaker,
			defaultWeightMsgCreateStaker,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bitvmstakersimulation.SimulateMsgCreateStaker(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSlashStaker,
			defaultWeightMsgSlashStaker,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bitvmstakersimulation.SimulateMsgSlashStaker(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
