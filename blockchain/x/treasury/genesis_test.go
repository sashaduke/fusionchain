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
package treasury_test

import (
	"testing"

	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/testutil/nullify"
	"github.com/qredo/fusionchain/x/treasury"
	"github.com/qredo/fusionchain/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	keepers := keepertest.NewTest(t)
	tk := keepers.TreasuryKeeper
	ctx := keepers.Ctx
	treasury.InitGenesis(ctx, *tk, genesisState)
	got := treasury.ExportGenesis(ctx, *tk)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
