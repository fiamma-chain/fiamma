package zkpverify

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "fiamma/api/fiamma/zkpverify"
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
					RpcMethod:      "ProofData",
					Use:            "get-proof-data [proof_id]",
					Short:          "Query Proof data stored in the fiamma by proof_id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_id"}},
				},

				{
					RpcMethod:      "BitVMWitness",
					Use:            "get-bitvm-witness [proof_id]",
					Short:          "Query bitvm witness stored in the fiamma by proof_id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_id"}},
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
					RpcMethod:      "SubmitProof",
					Use:            "submit-proof [proof_system] [proof] [public_input] [vk]",
					Short:          "Send a zkp proof verify tx" + "\n" + "Currently supported proof systems: " + "[PlonkBn254, PlonkBls12_381, Groth16Bn254, Groth16Bn254_BitVM, SP1]",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_system"}, {ProtoField: "proof"}, {ProtoField: "public_input"}, {ProtoField: "vk"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
