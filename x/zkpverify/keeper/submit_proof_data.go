package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SubmitProofData(ctx context.Context, proofId []byte, proofData types.ProofData) (string, types.DataLocation, error) {
	// We submit data to DA, if submission fails, then we will store the data on our own chain
	// Currently we only support NubitDA
	dataCommitments, err := k.SubmitProofDataToDA(ctx, proofId, proofData)
	if err == nil {
		dataCommitmentStr := hex.EncodeToString(dataCommitments[0])
		return dataCommitmentStr, types.DataLocation_NUBITDA, nil

	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// If submission to DA fails, we store the data on our own chain
	k.SetProofData(sdkCtx, proofId[:], proofData)
	return "", types.DataLocation_FIAMMA, nil
}

func (k Keeper) SubmitProofDataToDA(ctx context.Context, proofId []byte, proofData types.ProofData) ([][]byte, error) {
	// Create a new array to store the proof data
	// This array will be used to store the proof data that is submitted to the Nubit chain
	submitData := [][]byte{}

	// In order to realize the idempotency of all nodes,
	// It is necessary to get the data from the proof node first.
	// If the data is already on the chain, the data will not be submitted again.
	dataProofs, err := k.nubitDA.GetBlobProofs(ctx, [][]byte{proofId})
	if err == nil {
		return dataProofs, nil
	}

	byteArray, err := json.Marshal(proofData)
	if err != nil {
		return nil, err
	}

	// Append the proof data to the proof data array
	// This will add the proof data to the proof data array
	submitData = append(submitData, byteArray)

	// Submit the proof data to the Nubit chain
	dataCommitments, err := k.nubitDA.SubmitBlobs(ctx, submitData)
	return dataCommitments, err
}
