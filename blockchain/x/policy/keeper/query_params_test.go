// Copyright 2023 Qredo Ltd.
// This file is part of the Fusion library.
//
// The Fusion library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Fusion library. If not, see https://github.com/qredo/fusionchain/blob/main/LICENSE
package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/policy/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keepers := keepertest.NewTest(t)
	pk := keepers.PolicyKeeper
	ctx := keepers.Ctx
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	pk.SetParams(ctx, params)

	response, err := pk.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
