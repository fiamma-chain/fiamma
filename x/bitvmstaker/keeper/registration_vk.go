package keeper

import (
	"context"
	"fiamma/x/bitvmstaker/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// RegisterVK registers a new VK
func (k Keeper) RegisterVK(ctx context.Context, vk []byte) error {
	store := k.vkStore(ctx)
	store.Set(vk, []byte{1}) // Use a dummy value
	return nil
}

// RemoveVK removes a registered VK
func (k Keeper) RemoveVK(ctx context.Context, vk []byte) error {
	store := k.vkStore(ctx)
	store.Delete(vk)
	return nil
}

// IsVKRegistered checks if a VK is registered
func (k Keeper) IsVKRegistered(ctx context.Context, vk []byte) bool {
	store := k.vkStore(ctx)
	return store.Has(vk)
}

// GetRegisteredVKList retrieves the list of registered VKs with pagination
func (k Keeper) GetRegisteredVKList(ctx context.Context, pagination *query.PageRequest) ([][]byte, *query.PageResponse, error) {
	store := k.vkStore(ctx)

	var vkList [][]byte
	pageRes, err := query.Paginate(store, pagination, func(key []byte, _ []byte) error {
		vkList = append(vkList, key)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return vkList, pageRes, nil
}

// vkStore returns the KVStore for VK storage
func (k Keeper) vkStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.VkKey)
}
