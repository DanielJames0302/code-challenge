package keeper

import (
	"crudchain/x/crud/types"
)

var _ types.QueryServer = Keeper{}
