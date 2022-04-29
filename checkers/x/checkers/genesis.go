package checkers

import (
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NextGame != nil {
		k.SetNextGame(ctx, *genState.NextGame)
	}
	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, *elem)
	}
	// Set all the playerInfo
	for _, elem := range genState.PlayerInfoList {
		k.SetPlayerInfo(ctx, *elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// NOTE: ここ生成されてなかった。scaffold mapだからか？
	// genesis.StoredGameList = k.GetAllStoredGame(ctx)
	// NOTE* 代わりにこれがあったが、型を変更したからなのか上手く動かない。今はなんとなくなんとなくポインタを設定したりしてしまっている。
	// Get all storedGame
	storedGameList := k.GetAllStoredGame(ctx)
	for _, elem := range storedGameList {
		elem := elem
		genesis.StoredGameList = append(genesis.StoredGameList, &elem)
	}

	// Get all nextGame
	nextGame, found := k.GetNextGame(ctx)
	if found {
		genesis.NextGame = &nextGame
	}
	playerInfoList := k.GetAllPlayerInfo(ctx)
	for _, elem := range playerInfoList {
		elem := elem
		genesis.PlayerInfoList = append(genesis.PlayerInfoList, &elem)
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
