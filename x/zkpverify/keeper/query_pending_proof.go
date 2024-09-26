package keeper

import (
	"context"
	"fmt"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PendingProof(goCtx context.Context, req *types.QueryPendingProofRequest) (*types.QueryPendingProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	indexStore := k.PendingProofsIndexStore(ctx)
	var verifyResults []*types.VerifyResult
	pageRes, err := query.Paginate(indexStore, req.Pagination, func(key []byte, _ []byte) error {
		verifyResult, found := k.GetVerifyResult(ctx, key)
		if !found {
			return fmt.Errorf("verify result not found for proof ID %x", key)
		}
		verifyResults = append(verifyResults, &verifyResult)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryPendingProofResponse{PendingProofs: verifyResults, Pagination: pageRes}, nil
}

func (k Keeper) PendingProofByNamespace(goCtx context.Context, req *types.QueryPendingProofByNamespaceRequest) (*types.QueryPendingProofByNamespaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	indexStore := k.PendingProofsIndexStore(ctx)
	var verifyResults []*types.VerifyResult
	pageRes, err := query.FilteredPaginate(indexStore, req.Pagination, func(key []byte, _ []byte, accumulate bool) (bool, error) {
		verifyResult, found := k.GetVerifyResult(ctx, key)
		if !found {
			return false, fmt.Errorf("verify result not found for proof ID %x", key)
		}
		if verifyResult.Namespace == req.Namespace {
			if accumulate {
				verifyResults = append(verifyResults, &verifyResult)
			}
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return nil, err
	}
	return &types.QueryPendingProofByNamespaceResponse{PendingProofs: verifyResults, Pagination: pageRes}, nil
}
