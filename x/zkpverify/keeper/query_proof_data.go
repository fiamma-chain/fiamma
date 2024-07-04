package keeper

import (
	"context"
	"encoding/hex"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProofData(goCtx context.Context, req *types.QueryProofDataRequest) (*types.QueryProofDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proofId, err := hex.DecodeString(req.ProofId)
	if err != nil {
		return nil, types.ErrInvalidProofId
	}

	proofData, found := k.GetProofData(ctx, proofId)
	if !found {
		return nil, types.ErrProofDataNotFound
	}

	return &types.QueryProofDataResponse{ProofData: &proofData}, nil
}
