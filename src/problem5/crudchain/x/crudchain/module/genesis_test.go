package crudchain_test

import (
	"testing"

	keepertest "crudchain/testutil/keeper"
	"crudchain/testutil/nullify"
	crudchain "crudchain/x/crudchain/module"
	"crudchain/x/crudchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ResourceList: []types.Resource{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ResourceCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudchainKeeper(t)
	crudchain.InitGenesis(ctx, k, genesisState)
	got := crudchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ResourceList, got.ResourceList)
	require.Equal(t, genesisState.ResourceCount, got.ResourceCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
