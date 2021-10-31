package keeper

import (
	"github.com/hello/planet/x/foo/types"
)

var _ types.QueryServer = Keeper{}
