package keeper

import (
	"context"
	"encoding/binary"

	"github.com/fiamma-chain/fiamma/x/bitvmstaker/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// SetCommitteeAddress set the committee address to KVStore
func (k Keeper) AppendStaker(ctx context.Context, stakerInfo types.StakerInfo) {
	count := k.GetStakerCount(ctx)
	stakerInfo.StakerIndex = count
	store := k.stakerStore(ctx)
	appendedValue := k.cdc.MustMarshal(&stakerInfo)
	store.Set([]byte(stakerInfo.StakerAddress), appendedValue)
	k.SetStakerCount(ctx, count+1)
}

func (k Keeper) GetStaker(ctx context.Context, stakerAddress string) (types.StakerInfo, bool) {
	store := k.stakerStore(ctx)
	bz := store.Get([]byte(stakerAddress))
	if bz == nil {
		return types.StakerInfo{}, false
	}
	var stakerInfo types.StakerInfo
	k.cdc.MustUnmarshal(bz, &stakerInfo)
	return stakerInfo, true
}

func (k Keeper) GetAllStakerInfo(ctx context.Context, pagination *query.PageRequest) ([]types.StakerInfo, *query.PageResponse, error) {
	store := k.stakerStore(ctx)
	var stakerInfos []types.StakerInfo
	pageRes, err := query.Paginate(store, pagination, func(key []byte, value []byte) error {
		var stakerInfo types.StakerInfo
		if err := k.cdc.Unmarshal(value, &stakerInfo); err != nil {
			return err
		}

		stakerInfos = append(stakerInfos, stakerInfo)
		return nil
	})
	return stakerInfos, pageRes, err
}

// GetCommitteeAddress gets the committee address from KVStore
func (k Keeper) DeleteStaker(ctx context.Context, stakerAddress string) {
	store := k.stakerStore(ctx)
	store.Delete([]byte(stakerAddress))
}

func (k Keeper) GetStakerCount(ctx context.Context) uint64 {
	store := k.stakerCountStore(ctx)
	bz := store.Get(types.StakerCountKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetStakerCount(ctx context.Context, count uint64) {
	store := k.stakerCountStore(ctx)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.StakerCountKey, bz)
}

// stakerStore returns the KVStore of the staker address
func (k Keeper) stakerStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, types.StakerKey)
}

// stakerCountStore returns the KVStore of the staker count
func (k Keeper) stakerCountStore(ctx context.Context) prefix.Store {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	return prefix.NewStore(storeAdapter, []byte{})
}
