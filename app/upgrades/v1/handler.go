package v1

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/comdex-official/comdex/x/incentives"
	incentiveskeeper "github.com/comdex-official/comdex/x/incentives/keeper"
	incentivestypes "github.com/comdex-official/comdex/x/incentives/types"
	"github.com/comdex-official/comdex/x/liquidity"
	liquiditykeeper "github.com/comdex-official/comdex/x/liquidity/keeper"
	liquiditytypes "github.com/comdex-official/comdex/x/liquidity/types"
	"github.com/comdex-official/comdex/x/locking"
	lockingkeeper "github.com/comdex-official/comdex/x/locking/keeper"
	lockingtypes "github.com/comdex-official/comdex/x/locking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

//upgrade handler for version1
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	wasmkeeper *wasmkeeper.Keeper,
	liquiditykeeper *liquiditykeeper.Keeper,
	incentiveskeeper *incentiveskeeper.Keeper,
	lockingkeeper *lockingkeeper.Keeper,

) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {

		// Set all modules "old versions" to 1. Then the run migrations logic will
		// handle running their upgrade logics.
		fromVM := make(map[string]uint64)
		for moduleName := range mm.Modules {
			fromVM[moduleName] = 1
		}

		// EXCEPT Auth needs to run AFTER staking.
		//
		// See: https://github.com/cosmos/cosmos-sdk/issues/10591
		//
		// So we do this by making auth run last. This is done by setting auth's
		// consensus version to 2, running RunMigrations, then setting it back to 1,
		// and then running migrations again.
		fromVM[authtypes.ModuleName] = 2

		delete(fromVM, authz.ModuleName)
		delete(fromVM, liquiditytypes.ModuleName)
		delete(fromVM, incentivestypes.ModuleName)
		delete(fromVM, lockingtypes.ModuleName)

		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		params := wasmkeeper.GetParams(ctx)
		params.CodeUploadAccess = wasmtypes.AllowNobody
		wasmkeeper.SetParams(ctx, params)

		//INIT for new modules

		incentives.InitGenesis(ctx, *incentiveskeeper, incentivestypes.GenesisState{Params: incentivestypes.DefaultParams()})
		locking.InitGenesis(ctx, *lockingkeeper, lockingtypes.GenesisState{Params: lockingtypes.DefaultParams()})
		liquidity.InitGenesis(ctx, *liquiditykeeper, liquiditytypes.GenesisState{
			Params: liquiditytypes.DefaultParams()})

		// now update auth version back to v1, to run auth migration last
		newVM[authtypes.ModuleName] = 1

		ctx.Logger().Info("Now running migrations just for auth, to get auth migration to be last. " +
			"(https://github.com/cosmos/cosmos-sdk/issues/10591)")
		return mm.RunMigrations(ctx, configurator, newVM)
	}
}
