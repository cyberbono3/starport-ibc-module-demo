package foo_test

import (
	"testing"

	keepertest "github.com/hello/planet/testutil/keeper"
	"github.com/hello/planet/x/foo"
	"github.com/hello/planet/x/foo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FooKeeper(t)
	foo.InitGenesis(ctx, *k, genesisState)
	got := foo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Equal(t, genesisState.PortId, got.PortId)
	// this line is used by starport scaffolding # genesis/test/assert
}
