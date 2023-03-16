package v2_0_0 //nolint:revive,stylecheck

// CreateUpgradeHandler creates an SDK upgrade handler for v2_0_0
//func CreateUpgradeHandler(
//	mm *module.Manager,
//	configurator module.Configurator,
//) upgradetypes.UpgradeHandler {
//	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
//		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
//		if err != nil {
//			return newVM, err
//		}
//		return newVM, err
//	}
//}
//
//func CreateUpgradeHandlerV2(
//	mm *module.Manager,
//	configurator module.Configurator,
//) upgradetypes.UpgradeHandler {
//	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
//		// Refs:
//		// - https://docs.cosmos.network/master/building-modules/upgrade.html#registering-migrations
//		// - https://docs.cosmos.network/master/migrations/chain-upgrade-guide-044.html#chain-upgrade
//
//		// Add Interchain Accounts host module
//		// set the ICS27 consensus version so InitGenesis is not run
//		fromVM[icatypes.ModuleName] = mm.Modules[icatypes.ModuleName].ConsensusVersion()
//
//		// create ICS27 Controller submodule params, controller module not enabled.
//		controllerParams := icacontrollertypes.Params{}
//
//		// create ICS27 Host submodule params
//
//		// create ICS27 Host submodule params
//		hostParams := icahosttypes.Params{
//			HostEnabled: true,
//			AllowMessages: []string{
//				sdk.MsgTypeURL(&banktypes.MsgSend{}),
//				sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
//				sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
//				sdk.MsgTypeURL(&stakingtypes.MsgCreateValidator{}),
//				sdk.MsgTypeURL(&stakingtypes.MsgEditValidator{}),
//				sdk.MsgTypeURL(&distrtypes.MsgWithdrawDelegatorReward{}),
//				sdk.MsgTypeURL(&distrtypes.MsgSetWithdrawAddress{}),
//				sdk.MsgTypeURL(&distrtypes.MsgWithdrawValidatorCommission{}),
//				sdk.MsgTypeURL(&distrtypes.MsgFundCommunityPool{}),
//				sdk.MsgTypeURL(&govtypes.MsgVote{}),
//				sdk.MsgTypeURL(&authz.MsgExec{}),
//				sdk.MsgTypeURL(&authz.MsgGrant{}),
//				sdk.MsgTypeURL(&authz.MsgRevoke{}),
//			},
//		}
//		// No changes in existsing module and their states,
//		// This upgrades adds new modules and new states in the existing store
//
//		icamodule, correctTypecast := mm.Modules[icatypes.ModuleName].(ica.AppModule)
//		if !correctTypecast {
//			panic("mm.Modules[icatypes.ModuleName] is not of type ica.AppModule")
//		}
//		icamodule.InitModule(ctx, controllerParams, hostParams)
//
//		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
//		if err != nil {
//			return newVM, err
//		}
//		return newVM, err
//	}
//}
