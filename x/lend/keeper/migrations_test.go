package keeper_test

import (
	v4types "github.com/comdex-official/comdex/x/lend/migrations/v5.0.0.beta/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (s *KeeperTestSuite) TestMigrateLends() {
	oldLend := v4types.LendAsset{
		ID:                  1,
		AssetID:             1,
		PoolID:              1,
		Owner:               "cosmos1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t",
		AmountIn:            sdk.NewCoin("uasset1", newInt(10000000000)),
		LendingTime:         time.Now(),
		UpdatedAmountIn:     newInt(10000000000),
		AvailableToBorrow:   newInt(10000000000),
		Reward_Accumulated:  newInt(100),
		AppID:               1,
		GlobalIndex:         sdk.OneDec(),
		LastInteractionTime: time.Now(),
		CPoolName:           "abc",
	}
	s.keeper.SetOldLend(s.ctx, oldLend)
	err := s.migrator.MigrateTo5_0_0beta(s.ctx)
	if err != nil {
		return
	}
}
