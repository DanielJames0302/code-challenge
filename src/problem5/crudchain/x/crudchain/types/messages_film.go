package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFilm{}

func NewMsgCreateFilm(creator string, name string, description string, genre string) *MsgCreateFilm {
	return &MsgCreateFilm{
		Creator:     creator,
		Name:        name,
		Description: description,
		Genre:       genre,
	}
}

func (msg *MsgCreateFilm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFilm{}

func NewMsgUpdateFilm(creator string, id uint64, name string, description string, genre string) *MsgUpdateFilm {
	return &MsgUpdateFilm{
		Id:          id,
		Creator:     creator,
		Name:        name,
		Description: description,
		Genre:       genre,
	}
}

func (msg *MsgUpdateFilm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFilm{}

func NewMsgDeleteFilm(creator string, id uint64) *MsgDeleteFilm {
	return &MsgDeleteFilm{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteFilm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
