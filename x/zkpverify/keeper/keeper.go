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

// SetBitVMWitness stores witness data
func (k Keeper) SetBitVMWitness(ctx sdk.Context, proofId []byte, witnessData []byte) {
	store := k.bitVMWitnessStore(ctx)
	store.Set(proofId, witnessData)
}

// GetBitVMWitness retrieves witness data from the chain
func (k Keeper) GetBitVMWitness(ctx sdk.Context, proofId []byte) ([]byte, bool) {
	store := k.bitVMWitnessStore(ctx)
	bz := store.Get(proofId)
	if bz == nil {
		return nil, false
	}
	return bz, true
}

func (k Keeper) proofDataStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.ProofDataKey)
}

func (k Keeper) verifyResultStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.VerifyResultKey)
}

func (k Keeper) bitVMWitnessStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BitVMWitnessKey)
}

func (k Keeper) GetPendingProofs(ctx context.Context, req *types.QueryPendingProofRequest) (*types.QueryPendingProofResponse, error) {
	store := k.verifyResultStore(ctx)
	var verifyResults []*types.VerifyResult
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var verifyResult types.VerifyResult
		if err := k.cdc.Unmarshal(value, &verifyResult); err != nil {
			return err
		}
		if verifyResult.Status != types.VerificationStatus_DEFINITIVE_VALIDATION {
			verifyResults = append(verifyResults, &verifyResult)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryPendingProofResponse{PendingProofs: verifyResults, Pagination: pageRes}, nil
}
