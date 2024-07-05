package keeper

import (
	"context"
	"encoding/hex"
	"strconv"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	VerificationCountLimit uint64 = 4
)

func (k msgServer) SubmitCommunityVerification(goCtx context.Context, msg *types.MsgSubmitCommunityVerification) (*types.MsgSubmitCommunityVerificationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proofSystemId, err := types.ProofSystemIdFromString(msg.ProofSystem)
	if err != nil {
		k.Logger().Info("Error parsing proof system:", "error", err)
		return nil, types.ErrInvalidProofSystem

	}
	proofData := types.ProofData{
		ProofSystem: uint64(proofSystemId),
		Proof:       msg.Proof,
		PublicInput: msg.PublicInput,
		Vk:          msg.Vk,
	}

	proofId, err := k.GetProofId(proofData)
	if err != nil {
		k.Logger().Info("Error getting proof id:", "error", err)
		return nil, types.ErrGetProofId
	}

	verifyResult, found := k.GetVerifyResult(ctx, proofId[:])
	if !found {
		k.Logger().Info("Error finding proof id:", "error", hex.EncodeToString(proofId[:]))
		return nil, types.ErrInvalidProofId
	}

	if verifyResult.Status == types.VerificationStatus_DEFINITIVEVALIDATION {
		k.Logger().Info("Error exceeding verification period:", "error", hex.EncodeToString(proofId[:]))
		return nil, types.ErrVerifyPeriod
	}

	// The chain first verifies the correctness of the proofs submitted by the user, and saves the results.
	// The observer may challenge the result at a later stage.
	result, witness := k.verifyProof(&proofData)
	if !result {
		k.Logger().Info("Error verifying proof result:", "error", hex.EncodeToString(proofId[:]))
		return nil, types.ErrVerifyResult
	}
	verifyResult.Result = result

	// store witness if the proof system is BitVM
	if proofData.ProofSystem == uint64(types.Groth16Bn254_BitVM) {
		k.SetBitVMWitness(ctx, proofId[:], witness)
	}

	k.Logger().Info("Proof verification result for community:", "result", result)

	verifyResult.CommunityVerificationCount++
	if verifyResult.CommunityVerificationCount < VerificationCountLimit {
		verifyResult.Status = types.VerificationStatus_COMMUNITYVALIDATION
	} else {
		verifyResult.Status = types.VerificationStatus_DEFINITIVEVALIDATION
	}
	k.Logger().Info("Proof verification status for community:", "status", verifyResult.Status)

	k.SetVerifyResult(ctx, proofId[:], verifyResult)

	event := sdk.NewEvent("verifyFinished",
		sdk.NewAttribute("proofId", hex.EncodeToString(proofId[:])),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)),
		sdk.NewAttribute("proofSystem", msg.ProofSystem),
		sdk.NewAttribute("status", strconv.Itoa(int(verifyResult.Status))),
		sdk.NewAttribute("CommunityVerificationCount", strconv.Itoa(int(verifyResult.CommunityVerificationCount))))
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitCommunityVerificationResponse{}, nil
}
