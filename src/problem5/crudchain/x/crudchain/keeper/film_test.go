package keeper_test

import (
	"context"
	"testing"

	keepertest "crudchain/testutil/keeper"
	"crudchain/testutil/nullify"
	"crudchain/x/crudchain/keeper"
	"crudchain/x/crudchain/types"

	"github.com/stretchr/testify/require"
)

func createNFilm(keeper keeper.Keeper, ctx context.Context, n int) []types.Film {
	items := make([]types.Film, n)
	for i := range items {
		items[i].Id = keeper.AppendFilm(ctx, items[i])
	}
	return items
}

func TestFilmGet(t *testing.T) {
	keeper, ctx := keepertest.CrudchainKeeper(t)
	items := createNFilm(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetFilm(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestFilmRemove(t *testing.T) {
	keeper, ctx := keepertest.CrudchainKeeper(t)
	items := createNFilm(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFilm(ctx, item.Id)
		_, found := keeper.GetFilm(ctx, item.Id)
		require.False(t, found)
	}
}

func TestFilmGetAll(t *testing.T) {
	keeper, ctx := keepertest.CrudchainKeeper(t)
	items := createNFilm(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFilm(ctx)),
	)
}

func TestFilmCount(t *testing.T) {
	keeper, ctx := keepertest.CrudchainKeeper(t)
	items := createNFilm(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetFilmCount(ctx))
}
