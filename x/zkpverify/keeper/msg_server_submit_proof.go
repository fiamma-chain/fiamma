package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitProof(goCtx context.Context, msg *types.MsgSubmitProof) (*types.MsgSubmitProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return &types.MsgSubmitProofResponse{}, err
	}

	// check if the vk is registered
	if !k.bitvmstakerKeeper.IsVKRegistered(ctx, msg.Vk) {
		return nil, types.ErrVKNotRegistered
	}

	proofData := types.ProofData{
		Namespace:    msg.Namespace,
		ProofSystem:  types.ProofSystem(types.ProofSystem_value[msg.ProofSystem]),
		Proof:        msg.Proof,
		PublicInput:  msg.PublicInput,
		Vk:           msg.Vk,
		DataLocation: types.DataLocation(types.DataLocation_value[msg.DataLocation]),
	}

	proofId, err := k.GetProofId(proofData)
	if err != nil {
		k.Logger().Info("Error getting proof id:", "error", err)
		return nil, types.ErrGetProofId

	}
	// The chain first verifies the correctness of the proofs submitted by the user, and saves the results.
	// The observer may challenge the result at a later stage.
	result, witness := k.verifyProof(&proofData)
	k.AddPendingProofIndex(ctx, proofId[:])

	proofIdHex := hex.EncodeToString(proofId[:])

	// TODO: remove this
	// This is a buggy proofId for testing the bitvm challenge process
	if proofIdHex == "1735e881fa5e58408e4710a4e8cbea0a7995f029eefdf85d7e59775b0b6c44c5" {
		result = false
	}

	// get the proposer address
	currentHeight := ctx.BlockHeight()
	proposerAddress := k.GetBlockProposer(ctx, currentHeight)

	if proofData.DataLocation != types.DataLocation_FIAMMA {
		// enqueue the proof for data availability submission and bitvm challenge
		daSubmissionData := types.DASubmissionData{
			ProofId:   proofIdHex,
			ProofData: &proofData,
		}
		k.EnqueueDASubmission(ctx, proofId[:], daSubmissionData)
	} else {
		// store the proof data in the fiamma store
		k.SetProofData(ctx, proofId[:], proofData)
	}

	bitVMChallengeData := types.BitVMChallengeData{
		Witness:  witness,
		Proposer: proposerAddress,
	}
	k.SetBitVMChallengeData(ctx, proofId[:], bitVMChallengeData)

	// store verify data in the store
	verifyResult := types.VerifyResult{
		ProofId:                    proofIdHex,
		ProofSystem:                proofData.ProofSystem,
		Result:                     result,
		Status:                     types.VerificationStatus_INITIAL_VALIDATION,
		CommunityVerificationCount: uint64(0),
		Namespace:                  proofData.Namespace,
	}
	k.SetVerifyResult(ctx, proofId[:], verifyResult)

	event := sdk.NewEvent("SubmitProof",
		sdk.NewAttribute("namespace", proofData.Namespace),
		sdk.NewAttribute("proofSystem", msg.ProofSystem),
		sdk.NewAttribute("proofId", proofIdHex),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSubmitProofResponse{}, nil
}

// GetProofId returns the proof id
func (k Keeper) GetProofId(proofData types.ProofData) ([32]byte, error) {
	var buf bytes.Buffer
	buf.Write([]byte(proofData.Namespace))
	buf.Write([]byte(proofData.ProofSystem.String()))
	buf.Write(proofData.Proof)
	buf.Write(proofData.PublicInput)
	buf.Write(proofData.Vk)
	buf.Write([]byte(proofData.DataLocation.String()))

	hash := sha256.Sum256(buf.Bytes())
	return hash, nil
}
