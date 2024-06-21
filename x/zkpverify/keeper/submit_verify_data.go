package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/json"

	"fiamma/x/nubit"
	"fiamma/x/zkpverify/types"
)

func (k Keeper) SubmitVerifyDataToDA(ctx context.Context, verifyData types.VerifyData) ([32]byte, [][]byte, error) {
	// Create a new NubitDA instance
	// This instance will be used to submit the proof data to the Nubit chain
	// The NubitDA instance is created with the logger instance
	nubitDA, err := nubit.NewNubitDA()
	if err != nil {
		return [32]byte{}, nil, err
	}

	// Create a new array to store the proof data
	// This array will be used to store the proof data that is submitted to the Nubit chain
	verifySubmitData := [][]byte{}

	// Currently submit one blob at a time.
	byteArray, err := json.Marshal(verifyData)
	if err != nil {
		return [32]byte{}, nil, err
	}
	verifyId := sha256.Sum256(byteArray)
	// Append the proof data to the proof data array
	// This will add the proof data to the proof data array
	verifySubmitData = append(verifySubmitData, byteArray)

	// Submit the proof data to the Nubit chain
	dataCommitments, err := nubitDA.SubmitBlobs(ctx, verifySubmitData)
	if err != nil {
		return [32]byte{}, nil, err
	}

	// Return the proof id and data commitment
	// This will return the proof id that is generated when the proof data is submitted to the Nubit chain
	return verifyId, dataCommitments, nil
}
