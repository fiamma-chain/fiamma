package keeper

import (
	"context"

	"fiamma/x/bitvmstaker/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetCommitteeAddress set the committee address to KVStore
func (k Keeper) SetCommitteeAddress(ctx context.Context, committeeAddress string) {
	store := k.committeeAddressStore(ctx)
	store.Set(types.CommitteeAddressKey, []byte(committeeAddress))
}

// GetCommitteeAddress gets the committee address from KVStore
func (k Keeper) GetCommitteeAddress(ctx context.Context) string {
	store := k.committeeAddressStore(ctx)
	address := store.Get(types.CommitteeAddressKey)
	if address == nil {
		panic("stored committee address should not have been nil")
	}
	return string(address)
}

// committeeAddressStore returns the KVStore of the committee address
func (k Keeper) committeeAddressStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, []byte{})
}
