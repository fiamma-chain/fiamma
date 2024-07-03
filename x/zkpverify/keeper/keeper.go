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

// SetVerifyResult stores proof information
func (k Keeper) SetVerifyResult(ctx sdk.Context, verifyId []byte, verifyResult types.VerifyResult) {
	store := k.verifyResultStore(ctx)
	bz := k.cdc.MustMarshal(&verifyResult)
	store.Set(verifyId, bz)
}

// GetVerifyResult retrieves proof information
func (k Keeper) GetVerifyResult(ctx sdk.Context, verifyId []byte) (types.VerifyResult, bool) {
	store := k.verifyResultStore(ctx)
	bz := store.Get(verifyId)
	if bz == nil {
		return types.VerifyResult{}, false
	}
	var verifyResult types.VerifyResult
	k.cdc.MustUnmarshal(bz, &verifyResult)
	return verifyResult, true
}

// SetVerifyData stores proof information
func (k Keeper) SetVerifyData(ctx sdk.Context, verifyId []byte, verifyData types.VerifyData) {
	store := k.verifyDataStore(ctx)
	bz := k.cdc.MustMarshal(&verifyData)
	store.Set(verifyId, bz)
}

// GetVerifyData retrieves proof information
func (k Keeper) GetVerifyData(ctx sdk.Context, verifyId []byte) (types.VerifyData, bool) {
	store := k.verifyDataStore(ctx)
	bz := store.Get(verifyId)
	if bz == nil {
		return types.VerifyData{}, false
	}
	var verifyData types.VerifyData
	k.cdc.MustUnmarshal(bz, &verifyData)
	return verifyData, true
}

// SetBitVMWitness stores witness data
func (k Keeper) SetBitVMWitness(ctx sdk.Context, verifyId []byte, witnessData []byte) {
	store := k.bitVMWitnessStore(ctx)
	store.Set(verifyId, witnessData)
}

// GetBitVMWitness retrieves witness data from the chain
func (k Keeper) GetBitVMWitness(ctx sdk.Context, verifyId []byte) ([]byte, bool) {
	store := k.bitVMWitnessStore(ctx)
	bz := store.Get(verifyId)
	if bz == nil {
		return nil, false
	}
	return bz, true
}

func (k Keeper) verifyDataStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.VerifyDataKey)
}

func (k Keeper) verifyResultStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.VerifyResultKey)
}

func (k Keeper) bitVMWitnessStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.BitVMWitnessKey)
}
