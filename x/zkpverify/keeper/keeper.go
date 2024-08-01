package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"fiamma/nubitda"
	"fiamma/x/zkpverify/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		stakingKeeper types.StakingKeeper

		// Nubit Data Availability
		nubitDA *nubitda.NubitDA
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	stakingKeeper types.StakingKeeper,
	nubitDA *nubitda.NubitDA,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:           cdc,
		storeService:  storeService,
		authority:     authority,
		logger:        logger,
		nubitDA:       nubitDA,
		stakingKeeper: stakingKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetVerifyResult stores proof verification information
func (k Keeper) SetVerifyResult(ctx sdk.Context, proofId []byte, verifyResult types.VerifyResult) {
	store := k.verifyResultStore(ctx)
	bz := k.cdc.MustMarshal(&verifyResult)
	store.Set(proofId, bz)
}

// GetVerifyResult retrieves proof verification information
func (k Keeper) GetVerifyResult(ctx sdk.Context, proofId []byte) (types.VerifyResult, bool) {
	store := k.verifyResultStore(ctx)
	bz := store.Get(proofId)
	if bz == nil {
		return types.VerifyResult{}, false
	}
	var verifyResult types.VerifyResult
	k.cdc.MustUnmarshal(bz, &verifyResult)
	return verifyResult, true
}

// SetProofData stores proof information
func (k Keeper) SetProofData(ctx sdk.Context, proofId []byte, proofData types.ProofData) {
	store := k.proofDataStore(ctx)
	bz := k.cdc.MustMarshal(&proofData)
	store.Set(proofId, bz)
}

// GetProofData retrieves proof information
func (k Keeper) GetProofData(ctx sdk.Context, proofId []byte) (types.ProofData, bool) {
	store := k.proofDataStore(ctx)
	bz := store.Get(proofId)
	if bz == nil {
		return types.ProofData{}, false
	}
	var proofData types.ProofData
	k.cdc.MustUnmarshal(bz, &proofData)
	return proofData, true
}

// SetBitVMChallengeData stores witness data
func (k Keeper) SetBitVMChallengeData(ctx sdk.Context, proofId []byte, challengeData types.BitVMChallengeData) {
	store := k.BitVMChallengeDataStore(ctx)
	bz := k.cdc.MustMarshal(&challengeData)
	store.Set(proofId, bz)
}

// GetBitVMChallengeData retrieves witness data from the chain
func (k Keeper) GetBitVMChallengeData(ctx sdk.Context, proofId []byte) (types.BitVMChallengeData, bool) {
	store := k.BitVMChallengeDataStore(ctx)
	bz := store.Get(proofId)
	if bz == nil {
		return types.BitVMChallengeData{}, false
	}
	var challengeData types.BitVMChallengeData
	k.cdc.MustUnmarshal(bz, &challengeData)
	return challengeData, true
}

// SetPendingProof stores pending proof verification information
func (k Keeper) SetPendingProof(ctx sdk.Context, proofId []byte, verifyResult types.VerifyResult) {
	store := k.pendingProofsStore(ctx)
	bz := k.cdc.MustMarshal(&verifyResult)
	store.Set(proofId, bz)
}

// GetPendingProof retrieves pending proof verification information
func (k Keeper) GetPendingProof(ctx sdk.Context, proofId []byte) (types.VerifyResult, bool) {
	store := k.pendingProofsStore(ctx)
	bz := store.Get(proofId)
	if bz == nil {
		return types.VerifyResult{}, false
	}
	var verifyResult types.VerifyResult
	k.cdc.MustUnmarshal(bz, &verifyResult)
	return verifyResult, true
}

// DeletePendingProof deletes pending proof verification information
func (k Keeper) DeletePendingProof(ctx sdk.Context, proofId []byte) {
	store := k.pendingProofsStore(ctx)
	store.Delete(proofId)
}

func (k Keeper) GetPendingProofs(ctx context.Context, req *types.QueryPendingProofRequest) (*types.QueryPendingProofResponse, error) {
	store := k.pendingProofsStore(ctx)
	var verifyResults []*types.VerifyResult
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var verifyResult types.VerifyResult
		if err := k.cdc.Unmarshal(value, &verifyResult); err != nil {
			return err
		}
		verifyResults = append(verifyResults, &verifyResult)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryPendingProofResponse{PendingProofs: verifyResults, Pagination: pageRes}, nil
}

func (k Keeper) SetBlockProposer(ctx context.Context, height int64, proposer string) {
	store := k.blockProposerStore(ctx)
	store.Set(sdk.Uint64ToBigEndian(uint64(height)), []byte(proposer))
}

func (k Keeper) GetBlockProposer(ctx context.Context, height int64) string {
	store := k.blockProposerStore(ctx)
	bz := store.Get(sdk.Uint64ToBigEndian(uint64(height)))
	return string(bz)
}

func (k Keeper) proofDataStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.ProofDataKey)
}

func (k Keeper) verifyResultStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.VerifyResultKey)
}

func (k Keeper) BitVMChallengeDataStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BitVMChallengeDataKey)
}

func (k Keeper) pendingProofsStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.PendingProofsKey)
}

func (k Keeper) blockProposerStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BlockProposerKey)
}
