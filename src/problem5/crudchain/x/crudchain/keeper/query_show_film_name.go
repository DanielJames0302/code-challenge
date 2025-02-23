package keeper

import (
	"context"

	"crudchain/x/crudchain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowFilmName(goCtx context.Context, req *types.QueryShowFilmNameRequest) (*types.QueryShowFilmNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	filmStore := prefix.NewStore(store, types.KeyPrefix(types.FilmKey))

	iterator := filmStore.Iterator(nil, nil)
	defer iterator.Close()

	var films []*types.Film
	for ; iterator.Valid(); iterator.Next() {
		film := new(types.Film)
		if err := k.cdc.Unmarshal(iterator.Value(), film); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		if film.Name == req.Name {
			films = append(films, film)
		}
	}

	return &types.QueryShowFilmNameResponse{Films: films}, nil
}
