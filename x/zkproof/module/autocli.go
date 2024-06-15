package zkproof

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "fiamma/api/fiamma/zkproof"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "PendingProof",
					Use:            "pending-proof",
					Short:          "Query pending-proof",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "PendingProofByType",
					Use:            "pending-proof-by-type",
					Short:          "Query pending-proof-by-type",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "AllProofTypes",
					Use:            "all-proof-types",
					Short:          "Query all-proof-types",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "SubmitGnarkPlonk",
					Use:            "submit-gnark-plonk [proof-id] [proof] [public-inputs] [verifying-key] [meta-data]",
					Short:          "Send a submit-gnark-plonk tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proofId"}, {ProtoField: "proof"}, {ProtoField: "publicInputs"}, {ProtoField: "verifyingKey"}, {ProtoField: "metaData"}},
				},
				{
					RpcMethod:      "VerifyProof",
					Use:            "verify-proof [proof-id] [result]",
					Short:          "Send a verify-proof tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proofId"}, {ProtoField: "result"}},
				},
				{
					RpcMethod:      "SubmitSp1",
					Use:            "submit-sp-1 [proof-id] [proof] [elf] [meta-data]",
					Short:          "Send a submit-sp1 tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proofId"}, {ProtoField: "proof"}, {ProtoField: "elf"}, {ProtoField: "metaData"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
