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
	store := k.VerifyResultStore(ctx)
	bz := k.cdc.MustMarshal(&verifyResult)
	store.Set(proofId, bz)
}

// GetVerifyResult retrieves proof verification information
func (k Keeper) GetVerifyResult(ctx context.Context, proofId []byte) (types.VerifyResult, bool) {
	store := k.VerifyResultStore(ctx)
	bz := store.Get(proofId)
	if bz == nil {
		return types.VerifyResult{}, false
	}
	var verifyResult types.VerifyResult
	k.cdc.MustUnmarshal(bz, &verifyResult)
	return verifyResult, true
}

// SetProofData stores proof information
func (k Keeper) SetProofData(ctx context.Context, proofId []byte, proofData types.ProofData) {
	store := k.proofDataStore(ctx)
	bz := k.cdc.MustMarshal(&proofData)
	store.Set(proofId, bz)
}

// GetProofData retrieves proof information
func (k Keeper) GetProofData(ctx context.Context, proofId []byte) (types.ProofData, bool) {
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
func (k Keeper) SetBitVMChallengeData(ctx context.Context, proofId []byte, challengeData types.BitVMChallengeData) {
	store := k.bitVMChallengeDataStore(ctx)
	bz := k.cdc.MustMarshal(&challengeData)
	store.Set(proofId, bz)
}

// GetBitVMChallengeData retrieves witness data from the chain
func (k Keeper) GetBitVMChallengeData(ctx context.Context, proofId []byte) (types.BitVMChallengeData, bool) {
	store := k.bitVMChallengeDataStore(ctx)
	bz := store.Get(proofId)
	if bz == nil {
		return types.BitVMChallengeData{}, false
	}
	var challengeData types.BitVMChallengeData
	k.cdc.MustUnmarshal(bz, &challengeData)
	return challengeData, true
}

// IsPendingProof checks if a proof ID is in the pending proofs index
func (k Keeper) IsPendingProof(ctx context.Context, proofId []byte) bool {
	store := k.PendingProofsIndexStore(ctx)
	return store.Has(proofId)
}

// AddPendingProofIndex adds a proof ID to the pending proofs index
func (k Keeper) AddPendingProofIndex(ctx context.Context, proofId []byte) {
	store := k.PendingProofsIndexStore(ctx)
	store.Set(proofId, []byte{1}) // We only need to store the key, value can be a dummy byte
}

// RemovePendingProofIndex removes a proof ID from the pending proofs index
func (k Keeper) RemovePendingProofIndex(ctx context.Context, proofId []byte) {
	store := k.PendingProofsIndexStore(ctx)
	store.Delete(proofId)
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

func (k Keeper) VerifyResultStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.VerifyResultKey)
}

func (k Keeper) bitVMChallengeDataStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BitVMChallengeDataKey)
}

func (k Keeper) blockProposerStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BlockProposerKey)
}

// PendingProofsIndexStore returns a prefix store for the pending proofs index
func (k Keeper) PendingProofsIndexStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.PendingProofsIndexKey)
}
