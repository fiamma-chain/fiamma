package zkpverify

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/fiamma-chain/fiamma/api/fiamma/zkpverify"
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
					RpcMethod:      "PendingProofByNamespace",
					Use:            "pending-proof-by-namespace [namespace]",
					Short:          "Query pending-proof by namespace",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "namespace"}},
				},

				{
					RpcMethod:      "ProofData",
					Use:            "get-proof-data [proof_id]",
					Short:          "Query Proof data stored in the fiamma by proof_id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_id"}},
				},

				{
					RpcMethod:      "VerifyResult",
					Use:            "get-verify-result [proof_id]",
					Short:          "Query Proof verified result by proof_id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_id"}},
				},

				{
					RpcMethod:      "VerifyResultsByNamespace",
					Use:            "get-verify-results-by-namespace [namespace]",
					Short:          "Query Proof verified results by namespace",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "namespace"}},
				},

				{
					RpcMethod:      "DASubmissionQueue",
					Use:            "get-da-submission-queue",
					Short:          "Query DA submission queue",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "DASubmitter",
					Use:            "get-da-submitter",
					Short:          "Query DA submitter address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "BitVMChallengeData",
					Use:            "get-bitvm-challenge-data [proof_id]",
					Short:          "Query BitVM challenge data by proof_id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_id"}},
				},

				{
					RpcMethod:      "DaSubmissionData",
					Use:            "get-da-submission-data [proof_id]",
					Short:          "Query DA submission data by proof_id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_id"}},
				},

				{
					RpcMethod:      "DaSubmissionResult",
					Use:            "get-da-submission-result [proof_id]",
					Short:          "Query DA submission result by proof_id",
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
					Use:            "submit-proof [namespace] [proof_system] [proof] [public_input] [vk] [data_location]",
					Short:          "Send a zkp proof verify tx" + "\n" + "Currently supported proof systems: " + "[GROTH16_BN254_BITVM, FFPLONK_BN254_BITVM]",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "namespace"}, {ProtoField: "proof_system"}, {ProtoField: "proof"}, {ProtoField: "public_input"}, {ProtoField: "vk"}, {ProtoField: "data_location"}},
				},
				{
					RpcMethod:      "SubmitCommunityVerification",
					Use:            "submit-community-verification [proof_id] [verify_result]",
					Short:          "submit a community zkp proof verify tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof_id"}, {ProtoField: "verify_result"}},
				},

				{
					RpcMethod:      "UpdateDASubmitter",
					Use:            "update-da-submitter [da_submitter]",
					Short:          "Update DA submitter address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "da_submitter"}},
				},

				{
					RpcMethod:      "UpdateDASubmissionResults",
					Use:            "update-da-submission-results [da_submission_result]",
					Short:          "Update DA submission results",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "da_submission_result"}},
				},

				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
