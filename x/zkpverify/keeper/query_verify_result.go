package keeper

import (
	"context"
	"encoding/hex"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifyResult(goCtx context.Context, req *types.QueryVerifyResultRequest) (*types.QueryVerifyResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proofId, err := hex.DecodeString(req.ProofId)
	if err != nil {
		return nil, types.ErrInvalidProofId
	}

	verifyResult, found := k.GetVerifyResult(ctx, proofId[:])
	if found {
		return &types.QueryVerifyResultResponse{VerifyResult: &verifyResult}, nil
	}

	return nil, types.ErrVerifyResultNotFound
}
