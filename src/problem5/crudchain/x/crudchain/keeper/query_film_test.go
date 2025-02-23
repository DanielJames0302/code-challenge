package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "crudchain/testutil/keeper"
	"crudchain/testutil/nullify"
	"crudchain/x/crudchain/types"
)

func TestFilmQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CrudchainKeeper(t)
	msgs := createNFilm(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetFilmRequest
		response *types.QueryGetFilmResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetFilmRequest{Id: msgs[0].Id},
			response: &types.QueryGetFilmResponse{Film: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetFilmRequest{Id: msgs[1].Id},
			response: &types.QueryGetFilmResponse{Film: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetFilmRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Film(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestFilmQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CrudchainKeeper(t)
	msgs := createNFilm(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllFilmRequest {
		return &types.QueryAllFilmRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FilmAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Film), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Film),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FilmAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Film), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Film),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.FilmAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Film),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.FilmAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
