package v5_0_0

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

//func UpdateDutchLendAuctions(
//	ctx sdk.Context,
//	liquidationkeeper liquidationkeeper.Keeper,
//	auctionkeeper auctionkeeper.Keeper,
//) {
//	lockedVaults := liquidationkeeper.GetLockedVaults(ctx)
//	for _, v := range lockedVaults {
//		if v.Kind != nil {
//			err := auctionkeeper.LendDutchActivator(ctx, v)
//			if err != nil {
//				return
//			}
//		}
//	}
//}

// CreateUpgradeHandler creates an SDK upgrade handler for v5_0_0
func CreateUpgradeHandlerV500(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// This change is only for testnet upgrade
		delete(fromVM, "lendV1")
		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)

		if err != nil {
			return newVM, err
		}
		return newVM, err
	}
}
