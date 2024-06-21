package zkpverify

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"fiamma/testutil/sample"
	zkpverifysimulation "fiamma/x/zkpverify/simulation"
	"fiamma/x/zkpverify/types"
)

// avoid unused import issue
var (
	_ = zkpverifysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitGnarkPlonk = "op_weight_msg_submit_gnark_plonk"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitGnarkPlonk int = 100

	opWeightMsgVerifyProof = "op_weight_msg_verify_proof"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVerifyProof int = 100

	opWeightMsgSubmitSp1 = "op_weight_msg_submit_sp_1"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitSp1 int = 100

	opWeightMsgSubmitGnarkGroth16 = "op_weight_msg_submit_gnark_groth_16"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitGnarkGroth16 int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	zkpverifyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&zkpverifyGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitGnarkPlonk int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitGnarkPlonk, &weightMsgSubmitGnarkPlonk, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitGnarkPlonk = defaultWeightMsgSubmitGnarkPlonk
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitGnarkPlonk,
		zkpverifysimulation.SimulateMsgSubmitGnarkPlonk(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVerifyProof int
	simState.AppParams.GetOrGenerate(opWeightMsgVerifyProof, &weightMsgVerifyProof, nil,
		func(_ *rand.Rand) {
			weightMsgVerifyProof = defaultWeightMsgVerifyProof
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVerifyProof,
		zkpverifysimulation.SimulateMsgVerifyProof(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitSp1 int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitSp1, &weightMsgSubmitSp1, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitSp1 = defaultWeightMsgSubmitSp1
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitSp1,
		zkpverifysimulation.SimulateMsgSubmitSp1(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitGnarkGroth16 int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitGnarkGroth16, &weightMsgSubmitGnarkGroth16, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitGnarkGroth16 = defaultWeightMsgSubmitGnarkGroth16
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitGnarkGroth16,
		zkpverifysimulation.SimulateMsgSubmitGnarkGroth16(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitGnarkPlonk,
			defaultWeightMsgSubmitGnarkPlonk,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zkpverifysimulation.SimulateMsgSubmitGnarkPlonk(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVerifyProof,
			defaultWeightMsgVerifyProof,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zkpverifysimulation.SimulateMsgVerifyProof(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitSp1,
			defaultWeightMsgSubmitSp1,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zkpverifysimulation.SimulateMsgSubmitSp1(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitGnarkGroth16,
			defaultWeightMsgSubmitGnarkGroth16,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zkpverifysimulation.SimulateMsgSubmitGnarkGroth16(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
