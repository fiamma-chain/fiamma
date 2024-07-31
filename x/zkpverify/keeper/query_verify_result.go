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

	verifyResultPending, foundPending := k.GetPendingProof(ctx, proofId[:])
	if foundPending {
		return &types.QueryVerifyResultResponse{VerifyResult: &verifyResultPending}, nil
	}

	verifyResultDefinitive, foundDefinitive := k.GetVerifyResult(ctx, proofId[:])
	if foundDefinitive {
		return &types.QueryVerifyResultResponse{VerifyResult: &verifyResultDefinitive}, nil
	}

	return nil, types.ErrVerifyResultNotFound
}
