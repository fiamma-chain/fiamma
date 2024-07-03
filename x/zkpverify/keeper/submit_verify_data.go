package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SubmitVerifyData(ctx context.Context, verifyData types.VerifyData) ([32]byte, [][]byte, types.DataLocationId, error) {
	// We submit data to DA, if submission fails, then we will store the data on our own chain
	// Currently we only support NubitDA
	verifyId, dataCommitments, err := k.SubmitVerifyDataToDA(ctx, verifyData)
	if err == nil {
		return verifyId, dataCommitments, types.NubitDA, nil

	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// If submission to DA fails, we store the data on our own chain
	k.SetVerifyData(sdkCtx, verifyId[:], verifyData)
	return verifyId, nil, types.Fiamma, nil

}

func (k Keeper) SubmitVerifyDataToDA(ctx context.Context, verifyData types.VerifyData) ([32]byte, [][]byte, error) {

	// Create a new array to store the proof data
	// This array will be used to store the proof data that is submitted to the Nubit chain
	verifySubmitData := [][]byte{}

	// Currently submit one blob at a time.
	byteArray, err := json.Marshal(verifyData)
	if err != nil {
		return [32]byte{}, nil, err
	}

	verifyId := sha256.Sum256(byteArray)

	// In order to realize the idempotency of all nodes,
	// It is necessary to get the data from the proof node first.
	// If the data is already on the chain, the data will not be submitted again.
	dataProofs, err := k.nubitDA.GetBlobProofs(ctx, [][]byte{verifyId[:]})
	if err == nil {
		return verifyId, dataProofs, nil
	}

	// Append the proof data to the proof data array
	// This will add the proof data to the proof data array
	verifySubmitData = append(verifySubmitData, byteArray)

	// Submit the proof data to the Nubit chain
	dataCommitments, err := k.nubitDA.SubmitBlobs(ctx, verifySubmitData)
	if err != nil {
		return [32]byte{}, nil, err
	}

	// Return the proof id and data commitment
	// This will return the proof id that is generated when the proof data is submitted to the Nubit chain
	return verifyId, dataCommitments, nil
}
