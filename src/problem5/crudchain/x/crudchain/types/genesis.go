package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ResourceList: []Resource{},
		FilmList:     []Film{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in resource
	resourceIdMap := make(map[uint64]bool)
	resourceCount := gs.GetResourceCount()
	for _, elem := range gs.ResourceList {
		if _, ok := resourceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for resource")
		}
		if elem.Id >= resourceCount {
			return fmt.Errorf("resource id should be lower or equal than the last id")
		}
		resourceIdMap[elem.Id] = true
	}
	// Check for duplicated ID in film
	filmIdMap := make(map[uint64]bool)
	filmCount := gs.GetFilmCount()
	for _, elem := range gs.FilmList {
		if _, ok := filmIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for film")
		}
		if elem.Id >= filmCount {
			return fmt.Errorf("film id should be lower or equal than the last id")
		}
		filmIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
