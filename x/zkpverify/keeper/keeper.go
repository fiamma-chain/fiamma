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

		bitvmstakerKeeper types.BitvmstakerKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	stakingKeeper types.StakingKeeper,
	nubitDA *nubitda.NubitDA,
	bitvmstakerKeeper types.BitvmstakerKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:               cdc,
		storeService:      storeService,
		authority:         authority,
		logger:            logger,
		nubitDA:           nubitDA,
		stakingKeeper:     stakingKeeper,
		bitvmstakerKeeper: bitvmstakerKeeper,
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
func (k Keeper) SetVerifyResult(ctx sdk.Context, verifyResult types.VerifyResult) {
	store := k.VerifyResultStore(ctx)
	bz := k.cdc.MustMarshal(&verifyResult)
	store.Set([]byte(verifyResult.ProofId), bz)
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

func (k Keeper) GetDASubmissionQueue(ctx context.Context, pagination *query.PageRequest) ([]types.DASubmissionData, *query.PageResponse, error) {
	store := k.DASubmissionQueueStore(ctx)
	var daSubmissionList []types.DASubmissionData
	pageRes, err := query.Paginate(store, pagination, func(key []byte, value []byte) error {
		var daSubmission types.DASubmissionData
		k.cdc.MustUnmarshal(value, &daSubmission)
		daSubmissionList = append(daSubmissionList, daSubmission)
		return nil
	})
	return daSubmissionList, pageRes, err
}

func (k Keeper) EnqueueDASubmission(ctx context.Context, daSubmissionData types.DASubmissionData) {
	store := k.DASubmissionQueueStore(ctx)
	store.Set([]byte(daSubmissionData.ProofId), k.cdc.MustMarshal(&daSubmissionData))
}

func (k Keeper) DequeueDASubmission(ctx context.Context, proofId string) {
	store := k.DASubmissionQueueStore(ctx)
	store.Delete([]byte(proofId))
}

func (k Keeper) SetDASubmissionResult(ctx context.Context, result *types.DASubmissionResult) {
	store := k.DASubmissionResultsStore(ctx)
	store.Set([]byte(result.ProofId), k.cdc.MustMarshal(result))
}

func (k Keeper) GetDASubmissionResult(ctx context.Context, proofId string) (types.DASubmissionResult, bool) {
	store := k.DASubmissionResultsStore(ctx)
	bz := store.Get([]byte(proofId))
	if bz == nil {
		return types.DASubmissionResult{}, false
	}
	var result types.DASubmissionResult
	k.cdc.MustUnmarshal(bz, &result)
	return result, true
}

func (k Keeper) SetDASubmitter(ctx context.Context, submitter string) {
	store := k.DASubmitterStore(ctx)
	store.Set(types.DASubmitterKey, []byte(submitter))
}

func (k Keeper) GetDASubmitter(ctx context.Context) string {
	store := k.DASubmitterStore(ctx)
	bz := store.Get(types.DASubmitterKey)
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

func (k Keeper) blockProposerStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BlockProposerKey)
}

// PendingProofsIndexStore returns a prefix store for the pending proofs index
func (k Keeper) PendingProofsIndexStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.PendingProofsIndexKey)
}

// DASubmissionQueueStore returns a prefix store for the DA submission data
func (k Keeper) DASubmissionQueueStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.DASubmissionQueueKey)
}

// DASubmissionResultsStore returns a prefix store for the DA submission results
func (k Keeper) DASubmissionResultsStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.DASubmissionResultsKey)
}

// DASubmitterStore returns a prefix store for the DA submitter
func (k Keeper) DASubmitterStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, []byte{})
}
