package keeper

import (
	"context"
	"encoding/binary"

	"crudchain/x/crudchain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetFilmCount get the total number of film
func (k Keeper) GetFilmCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.FilmCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetFilmCount set the total number of film
func (k Keeper) SetFilmCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.FilmCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendFilm appends a film in the store with a new id and update the count
func (k Keeper) AppendFilm(
	ctx context.Context,
	film types.Film,
) uint64 {
	// Create the film
	count := k.GetFilmCount(ctx)

	// Set the ID of the appended value
	film.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilmKey))
	appendedValue := k.cdc.MustMarshal(&film)
	store.Set(GetFilmIDBytes(film.Id), appendedValue)

	// Update film count
	k.SetFilmCount(ctx, count+1)

	return count
}

// SetFilm set a specific film in the store
func (k Keeper) SetFilm(ctx context.Context, film types.Film) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilmKey))
	b := k.cdc.MustMarshal(&film)
	store.Set(GetFilmIDBytes(film.Id), b)
}

// GetFilm returns a film from its id
func (k Keeper) GetFilm(ctx context.Context, id uint64) (val types.Film, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilmKey))
	b := store.Get(GetFilmIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFilm removes a film from the store
func (k Keeper) RemoveFilm(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilmKey))
	store.Delete(GetFilmIDBytes(id))
}

// GetAllFilm returns all film
func (k Keeper) GetAllFilm(ctx context.Context) (list []types.Film) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FilmKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Film
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetFilmIDBytes returns the byte representation of the ID
func GetFilmIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.FilmKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
