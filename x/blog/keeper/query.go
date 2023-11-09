package keeper

import (
	"galactic-empire/x/blog/types"
)

var _ types.QueryServer = Keeper{}
