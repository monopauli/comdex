package v5_0_0_beta

import (
	"fmt"
	v4types "github.com/comdex-official/comdex/x/lend/migrations/v5.0.0.beta/types"
	lendtypes "github.com/comdex-official/comdex/x/lend/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//MigrationHandler codes goes here

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	err := migrateValuesLend(store, cdc)
	if err != nil {
		return err
	}
	//err = migrateValuesBorrow(store, cdc)
	//if err != nil {
	//	return err
	//}

	return nil
}
func migrateValuesLend(store sdk.KVStore, cdc codec.BinaryCodec) error {

	//keyCounter := lendtypes.LendCounterIDPrefix
	//value := store.Get(keyCounter)
	//var counter protobuftypes.UInt64Value
	//cdc.MustUnmarshal(value, &counter)

	//for ID := uint64(1); ID <= counter.GetValue(); ID++ {
	key := lendtypes.LendUserKey(1)
	oldVal := store.Get(key)
	fmt.Println("old...............")
	fmt.Println(oldVal)
	newVal := migrateValueLend(cdc, oldVal)
	fmt.Println("newval/////////")
	fmt.Println(newVal)
	//store.Delete(key) // Delete old key, value
	store.Set(key, newVal)
	//}
	return nil
}

func migrateValueLend(cdc codec.BinaryCodec, oldVal []byte) (newVal []byte) {

	// convert oldVal into lend type of previous version
	// use oldVal to create new lend of updated struct

	var lend v4types.LendAsset
	cdc.MustUnmarshal(oldVal, &lend)

	newLend := lendtypes.LendAsset{
		ID:                  lend.ID,
		AssetID:             lend.AssetID,
		PoolID:              lend.PoolID,
		Owner:               lend.Owner,
		AmountIn:            lend.AmountIn,
		LendingTime:         lend.LendingTime,
		AvailableToBorrow:   lend.AvailableToBorrow,
		AppID:               lend.AppID,
		GlobalIndex:         lend.GlobalIndex,
		LastInteractionTime: lend.LastInteractionTime,
		CPoolName:           lend.CPoolName,
	}

	newVal = cdc.MustMarshal(&newLend)
	return newVal
}

//
//func migrateValuesBorrow(store sdk.KVStore, cdc codec.BinaryCodec) error {
//
//	borrowIDKey := lendtypes.BorrowsKey
//	lendIDValue := store.Get(borrowIDKey)
//	var usersBorrowMap lendtypes.BorrowMapping
//	cdc.MustUnmarshal(lendIDValue, &usersBorrowMap)
//
//	for _, ID := range usersBorrowMap.BorrowIDs {
//		key := append(lendtypes.BorrowPairKeyPrefix, sdk.Uint64ToBigEndian(ID)...)
//		oldVal := store.Get(key)
//
//		newVal := migrateValueBorrow(cdc, oldVal)
//		store.Delete(key) // Delete old key, value
//		store.Set(key, newVal)
//	}
//	return nil
//}
//
//func migrateValueBorrow(cdc codec.BinaryCodec, oldVal []byte) (newVal []byte) {
//
//	// convert oldVal into borrow type of previous version
//	// use oldVal to create new borrow of updated struct
//
//	var borrow lendtypes.BorrowAsset
//	cdc.MustUnmarshal(oldVal, &borrow)
//
//	newBorrow := lendtypes.BorrowAsset{
//		ID:                  borrow.ID,
//		LendingID:           borrow.LendingID,
//		IsStableBorrow:      borrow.IsStableBorrow,
//		PairID:              borrow.PairID,
//		AmountIn:            borrow.AmountIn,
//		AmountOut:           borrow.AmountOut,
//		BridgedAssetAmount:  borrow.BridgedAssetAmount,
//		BorrowingTime:       borrow.BorrowingTime,
//		StableBorrowRate:    borrow.StableBorrowRate,
//		InterestAccumulated: borrow.InterestAccumulated,
//		GlobalIndex:         borrow.GlobalIndex,
//		ReserveGlobalIndex:  borrow.ReserveGlobalIndex,
//		LastInteractionTime: borrow.LastInteractionTime,
//		CPoolName:           borrow.CPoolName,
//	}
//
//	newVal = cdc.MustMarshal(&newBorrow)
//	return newVal
//}
