package keeper

import (
	"context"
	"fmt"

	"crudchain/x/crudchain/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateFilm(goCtx context.Context, msg *types.MsgCreateFilm) (*types.MsgCreateFilmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var film = types.Film{
		Creator:     msg.Creator,
		Name:        msg.Name,
		Description: msg.Description,
		Genre:       msg.Genre,
	}

	id := k.AppendFilm(
		ctx,
		film,
	)

	return &types.MsgCreateFilmResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateFilm(goCtx context.Context, msg *types.MsgUpdateFilm) (*types.MsgUpdateFilmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var film = types.Film{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Name:        msg.Name,
		Description: msg.Description,
		Genre:       msg.Genre,
	}

	// Checks that the element exists
	val, found := k.GetFilm(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetFilm(ctx, film)

	return &types.MsgUpdateFilmResponse{}, nil
}

func (k msgServer) DeleteFilm(goCtx context.Context, msg *types.MsgDeleteFilm) (*types.MsgDeleteFilmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetFilm(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveFilm(ctx, msg.Id)

	return &types.MsgDeleteFilmResponse{}, nil
}
