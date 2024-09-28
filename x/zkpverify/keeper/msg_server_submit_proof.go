package keeper

import (
	"context"
	"encoding/hex"
	"strconv"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitProof(goCtx context.Context, msg *types.MsgSubmitProof) (*types.MsgSubmitProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentHeight := ctx.BlockHeight()
	proposerAddress := k.GetBlockProposer(ctx, currentHeight)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgSubmitProofResponse{}, err
	}

	// check if the vk is registered
	if !k.bitvmstakerKeeper.IsVKRegistered(ctx, msg.Vk) {
		return nil, types.ErrVKNotRegistered
	}

	proofData := types.ProofData{
		Namespace:   msg.Namespace,
		ProofSystem: types.ProofSystem(types.ProofSystem_value[msg.ProofSystem]),
		Proof:       msg.Proof,
		PublicInput: msg.PublicInput,
		Vk:          msg.Vk,
	}

	proofId, err := k.GetProofId(proofData)
	if err != nil {
		k.Logger().Info("Error getting proof id:", "error", err)
		return nil, types.ErrGetProofId

	}

	// submit proof data to DA
	dataCommitmentStr, dataLocation, err := k.SubmitProofData(ctx, proofId[:], proofData)
	if err != nil {
		k.Logger().Info("Error submitting proof to DA:", "error", err)
		return nil, types.ErrSubmitProof
	}

	proofIdStr := hex.EncodeToString(proofId[:])

	// The chain first verifies the correctness of the proofs submitted by the user, and saves the results.
	// The observer may challenge the result at a later stage.
	result, witness := k.verifyProof(&proofData)

	// store witness if the proof system is BitVM
	if proofData.ProofSystem == types.ProofSystem_GROTH16_BN254_BITVM {
		// TODO: remove this
		// This is a buggy proofId for testing the bitvm challenge process
		if proofIdStr == "d14058dd7a41147beb577014cb2a356d52500f300c35b5f5ae8da89e47c4ec59" {
			result = false
		}
		bitvmChallengeData := types.BitVMChallengeData{
			VerifyResult: result,
			Witness:      witness,
			Vk:           msg.Vk,
			PublicInput:  msg.PublicInput,
			Proposer:     proposerAddress,
		}
		k.SetBitVMChallengeData(ctx, proofId[:], bitvmChallengeData)

	}

	// store verify data in the store
	verifyResult := types.VerifyResult{
		ProofId:                    proofIdStr,
		ProofSystem:                proofData.ProofSystem,
		DataCommitment:             dataCommitmentStr,
		DataLocation:               dataLocation,
		Result:                     result,
		Status:                     types.VerificationStatus_INITIAL_VALIDATION,
		CommunityVerificationCount: uint64(0),
		Namespace:                  proofData.Namespace,
	}

	k.AddPendingProofIndex(ctx, proofId[:])
	k.SetVerifyResult(ctx, proofId[:], verifyResult)

	event := sdk.NewEvent("SubmitProof",
		sdk.NewAttribute("namespace", proofData.Namespace),
		sdk.NewAttribute("proofSystem", msg.ProofSystem),
		sdk.NewAttribute("proofId", proofIdStr),
		sdk.NewAttribute("dataCommitment", dataCommitmentStr),
		sdk.NewAttribute("dataLocation", dataLocation.String()),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)),
		sdk.NewAttribute("proposer", proposerAddress),
	)

	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitProofResponse{}, nil
}
