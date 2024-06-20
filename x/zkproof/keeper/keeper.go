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

	"fiamma/x/zkproof/types"
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
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	stakingKeeper types.StakingKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

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

// SetProofInfo stores proof information
func (k Keeper) SetVerifyData(ctx sdk.Context, proofID string, proofInfo types.Zkproof) {
	store := k.verifyDataStore(ctx)
	bz := k.cdc.MustMarshal(&proofInfo)
	store.Set([]byte(proofID), bz)
}

// GetProofInfo retrieves proof information
func (k Keeper) GetProofInfo(ctx sdk.Context, proofID string) (types.Zkproof, bool) {
	store := k.verifyDataStore(ctx)
	bz := store.Get([]byte(proofID))
	if bz == nil {
		return types.Zkproof{}, false
	}
	var proofInfo types.Zkproof
	k.cdc.MustUnmarshal(bz, &proofInfo)
	return proofInfo, true
}

func (k Keeper) verifyDataStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.VerifyDataKey)
}
