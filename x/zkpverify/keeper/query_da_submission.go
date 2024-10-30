package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fiamma-chain/fiamma/x/zkpverify/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DaSubmissionData(goCtx context.Context, req *types.QueryDaSubmissionDataRequest) (*types.QueryDaSubmissionDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proofId, err := hex.DecodeString(req.ProofId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid proof id")
	}

	data, found := k.GetDASubmissionData(ctx, proofId)
	if !found {
		return nil, status.Error(codes.NotFound, "da submission data not found")
	}

	return &types.QueryDaSubmissionDataResponse{
		DaSubmissionData: &data,
	}, nil
}

func (k Keeper) DaSubmissionResult(goCtx context.Context, req *types.QueryDaSubmissionResultRequest) (*types.QueryDaSubmissionResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proofId, err := hex.DecodeString(req.ProofId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid proof id")
	}

	result, found := k.GetDASubmissionResult(ctx, proofId)
	if !found {
		return nil, status.Error(codes.NotFound, "da submission result not found")
	}

	return &types.QueryDaSubmissionResultResponse{
		DaSubmissionResult: &result,
	}, nil
}
