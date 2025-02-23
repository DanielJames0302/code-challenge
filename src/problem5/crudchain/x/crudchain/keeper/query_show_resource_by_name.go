package keeper

import (
	"context"

	"crudchain/x/crudchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowResourceByName(goCtx context.Context, req *types.QueryShowResourceByNameRequest) (*types.QueryShowResourceByNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryShowResourceByNameResponse{}, nil
}
