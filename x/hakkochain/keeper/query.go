package keeper

import (
	"hakko-chain/x/hakkochain/types"
)

var _ types.QueryServer = Keeper{}
