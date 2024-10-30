package bitvmstaker

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/fiamma-chain/fiamma/testutil/sample"

	bitvmstakersimulation "github.com/fiamma-chain/fiamma/x/bitvmstaker/simulation"
	"github.com/fiamma-chain/fiamma/x/bitvmstaker/types"
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

	opWeightMsgRemoveStaker = "op_weight_msg_slash_staker"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveStaker int = 100

	opWeightMsgUpdateCommitteeAddress = "op_weight_msg_update_committee_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCommitteeAddress int = 100

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

	var weightMsgRemoveStaker int
	simState.AppParams.GetOrGenerate(opWeightMsgRemoveStaker, &weightMsgRemoveStaker, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveStaker = defaultWeightMsgRemoveStaker
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveStaker,
		bitvmstakersimulation.SimulateMsgRemoveStaker(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCommitteeAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateCommitteeAddress, &weightMsgUpdateCommitteeAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCommitteeAddress = defaultWeightMsgUpdateCommitteeAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCommitteeAddress,
		bitvmstakersimulation.SimulateMsgUpdateCommitteeAddress(am.accountKeeper, am.bankKeeper, am.keeper),
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
			opWeightMsgRemoveStaker,
			defaultWeightMsgRemoveStaker,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bitvmstakersimulation.SimulateMsgRemoveStaker(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateCommitteeAddress,
			defaultWeightMsgUpdateCommitteeAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bitvmstakersimulation.SimulateMsgUpdateCommitteeAddress(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
