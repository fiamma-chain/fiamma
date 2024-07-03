package keeper

import (
	"context"
	"encoding/hex"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BitVMWitness(goCtx context.Context, req *types.QueryBitVMWitnessRequest) (*types.QueryBitVMWitnessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	witness, found := k.GetBitVMWitness(ctx, []byte(req.ProofId))
	if !found {
		return nil, types.ErrBitVMWitnessNotFound
	}

	return &types.QueryBitVMWitnessResponse{Witness: hex.EncodeToString(witness)}, nil
}
