package keeper

import (
	"context"

	"crudchain/x/crudchain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FilmAll(ctx context.Context, req *types.QueryAllFilmRequest) (*types.QueryAllFilmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var films []types.Film

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	filmStore := prefix.NewStore(store, types.KeyPrefix(types.FilmKey))

	pageRes, err := query.Paginate(filmStore, req.Pagination, func(key []byte, value []byte) error {
		var film types.Film
		if err := k.cdc.Unmarshal(value, &film); err != nil {
			return err
		}

		films = append(films, film)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFilmResponse{Film: films, Pagination: pageRes}, nil
}

func (k Keeper) Film(ctx context.Context, req *types.QueryGetFilmRequest) (*types.QueryGetFilmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	film, found := k.GetFilm(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetFilmResponse{Film: film}, nil
}
