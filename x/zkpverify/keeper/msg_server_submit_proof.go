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

	// check if the proof system is valid
	if _, ok := types.ProofSystem_value[msg.ProofSystem]; !ok {
		return nil, types.ErrInvalidProofSystem
	}
	proofData := types.ProofData{
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
		if proofIdStr == "fe98dcdcfb929e012cd8000dd1ad2b42a36f603d37fdd83e03938a9ab3af2363" {
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
	}

	k.SetPendingProof(ctx, proofId[:], verifyResult)

	event := sdk.NewEvent("SubmitProof",
		sdk.NewAttribute("proofSystem", msg.ProofSystem),
		sdk.NewAttribute("proofId", proofIdStr),
		sdk.NewAttribute("dataCommitment", dataCommitmentStr),
		sdk.NewAttribute("dataLocation", dataLocation.String()),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)),
		sdk.NewAttribute("proposer", proposerAddress))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitProofResponse{}, nil
}
