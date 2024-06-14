package zkproof

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"fiamma/testutil/sample"
	zkproofsimulation "fiamma/x/zkproof/simulation"
	"fiamma/x/zkproof/types"
)

// avoid unused import issue
var (
	_ = zkproofsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitGnarkPlonk = "op_weight_msg_submit_gnark_plonk"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitGnarkPlonk int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	zkproofGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&zkproofGenesis)
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
		zkproofsimulation.SimulateMsgSubmitGnarkPlonk(am.accountKeeper, am.bankKeeper, am.keeper),
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
				zkproofsimulation.SimulateMsgSubmitGnarkPlonk(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
