package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SubmitVerifyData(ctx context.Context, verifyId []byte, verifyData types.VerifyData) (string, types.DataLocationId, error) {
	// We submit data to DA, if submission fails, then we will store the data on our own chain
	// Currently we only support NubitDA
	dataCommitments, err := k.SubmitVerifyDataToDA(ctx, verifyId, verifyData)
	if err == nil {
		dataCommitmentStr := hex.EncodeToString(dataCommitments[0])
		return dataCommitmentStr, types.NubitDA, nil

	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// If submission to DA fails, we store the data on our own chain
	k.SetVerifyData(sdkCtx, verifyId[:], verifyData)
	return "", types.Fiamma, nil

}

func (k Keeper) SubmitVerifyDataToDA(ctx context.Context, verifyId []byte, verifyData types.VerifyData) ([][]byte, error) {

	// Create a new array to store the proof data
	// This array will be used to store the proof data that is submitted to the Nubit chain
	verifySubmitData := [][]byte{}

	// In order to realize the idempotency of all nodes,
	// It is necessary to get the data from the proof node first.
	// If the data is already on the chain, the data will not be submitted again.
	dataProofs, err := k.nubitDA.GetBlobProofs(ctx, [][]byte{verifyId})
	if err == nil {
		return dataProofs, nil
	}

	byteArray, err := json.Marshal(verifyData)
	if err != nil {
		return nil, err
	}

	// Append the proof data to the proof data array
	// This will add the proof data to the proof data array
	verifySubmitData = append(verifySubmitData, byteArray)

	// Submit the proof data to the Nubit chain
	dataCommitments, err := k.nubitDA.SubmitBlobs(ctx, verifySubmitData)
	return dataCommitments, err
}
