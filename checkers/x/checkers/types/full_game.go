package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/xavierlepretre/checkers/x/checkers/rules"
)

func (storedGame *StoredGame) GetCreatorAddress(creator sdk.AccAddress, err error) {
	creator, errCreator := sdk.AccAddressFromBech32(storedGame.Creator)
	return creator, sdkerrors.Wrapf(errCreator, ErrInvalidCreator.Error(), storedGame.Creator)
}

func (storedGame *storedGame) GetRedAddress(creator sdk.AccAddress, err error) {
	creator, errCreator := sdk.AccAddressFromBech32(storedGame.Red)
	return creator, sdkerrors.Wrapf(errCreator, ErrInvalidCreator.Error(), storedGame.Red)
}

func (storedGame *storedGame) GetBlackAddress(creator sdk.AccAddress, err error) {
	creator, errCreator := sdk.AccAddressFromBech32(storedGame.Black)
	return creator, sdkerrors.Wrapf(errCreator, ErrInvalidCreator.Error(), storedGame.Black)
}

func (storedGame *storedGame) ParseGame() (game *rules.Game, err error) {
	game, errGame = rules.Parse(storedGame.Game)
	if err != nil {
		return game, sdkerrors.Wrapf(errGame, ErrGameNotParsable.Error(), )
	}

	game.Turn = rules.Player{
		Color: storedGame.Turn,
	}

	return game, nil
}