package keeper

import (
	"context"
	"encoding/hex"

	"github.com/fiamma-chain/fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BitVMChallengeData(goCtx context.Context, req *types.QueryBitVMChallengeDataRequest) (*types.QueryBitVMChallengeDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	proofId, err := hex.DecodeString(req.ProofId)
	if err != nil {
		return nil, types.ErrInvalidProofId
	}

	chanllengeData, found := k.GetBitVMChallengeData(ctx, proofId)
	if !found {
		return nil, types.ErrBitVMChallengeDataNotFound
	}

	return &types.QueryBitVMChallengeDataResponse{BitvmChallengeData: &chanllengeData}, nil
}
