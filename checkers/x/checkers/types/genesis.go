package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
        StoredGameList: []*StoredGame{},
        NextGame:       &NextGame{"", uint64(0)},
    }
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedGame
	storedGameIndexMap := make(map[string]bool)

	for _, elem := range gs.StoredGameList {
		if _, ok := storedGameIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for storedGame")
		}
		storedGameIndexMap[elem.Index] = true
	}

	return nil
}
