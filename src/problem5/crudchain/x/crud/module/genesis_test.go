package crud_test

import (
	"testing"

	keepertest "crudchain/testutil/keeper"
	"crudchain/testutil/nullify"
	crud "crudchain/x/crud/module"
	"crudchain/x/crud/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudKeeper(t)
	crud.InitGenesis(ctx, k, genesisState)
	got := crud.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
