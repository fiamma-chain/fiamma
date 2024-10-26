package keeper

import (
	"context"
	"encoding/hex"

	"github.com/fiamma-chain/fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
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

func (k Keeper) VerifyResultsByNamespace(goCtx context.Context, req *types.QueryVerifyResultsByNamespaceRequest) (*types.QueryVerifyResultsByNamespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var verifyResults []*types.VerifyResult
	verifyResultStore := k.VerifyResultStore(ctx)

	pageRes, err := query.Paginate(verifyResultStore, req.Pagination, func(key []byte, value []byte) error {
		var verifyResult types.VerifyResult
		if err := k.cdc.Unmarshal(value, &verifyResult); err != nil {
			return err
		}

		if verifyResult.Namespace == req.Namespace {
			verifyResults = append(verifyResults, &verifyResult)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryVerifyResultsByNamespaceResponse{VerifyResults: verifyResults, Pagination: pageRes}, nil
}
