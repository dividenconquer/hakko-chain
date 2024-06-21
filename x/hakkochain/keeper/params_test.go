package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "hakko-chain/testutil/keeper"
	"hakko-chain/x/hakkochain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HakkochainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
