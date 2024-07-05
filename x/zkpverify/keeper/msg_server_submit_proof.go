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
		k.SetBitVMWitness(ctx, proofId[:], witness)
	}

	k.Logger().Info("Proof verification:", "result", result)

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

	k.SetVerifyResult(ctx, proofId[:], verifyResult)

	event := sdk.NewEvent("SubmitProof",
		sdk.NewAttribute("proofSystem", msg.ProofSystem),
		sdk.NewAttribute("proofId", proofIdStr),
		sdk.NewAttribute("dataCommitment", dataCommitmentStr),
		sdk.NewAttribute("dataLocation", dataLocation.String()),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)))

	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitProofResponse{}, nil
}
