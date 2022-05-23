package keeper

import (
	"context"
	assettypes "github.com/comdex-official/comdex/x/asset/types"
	"github.com/comdex-official/comdex/x/locker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServiceServer = (*queryServer)(nil)

type queryServer struct {
	Keeper
}

func NewQueryServiceServer(k Keeper) types.QueryServiceServer {
	return &queryServer{
		Keeper: k,
	}
}

func (q *queryServer) QueryParams(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = q.GetParams(ctx)
	)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}

func (q *queryServer) QueryLockerInfo(c context.Context, req *types.QueryLockerInfoRequest) (*types.QueryLockerInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)
	item, found := q.GetLocker(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "locker-info does not exist for id %d", req.Id)
	}

	return &types.QueryLockerInfoResponse{
		LockerInfo: item,
	}, nil
}

func (q *queryServer) QueryLockersByProductToAssetID(c context.Context, request *types.QueryLockersByProductToAssetIDRequest) (*types.QueryLockersByProductToAssetIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, found := q.GetLockerLookupTable(ctx, app.Id)

	if !found {
		return nil, status.Errorf(codes.NotFound, "no asset exists appID %d", app.Id)
	}

	var lockerIds []string
	for _, locker := range lockerLookupData.Lockers {
		if request.AssetId == locker.AssetId {
			lockerIds = append(lockerIds, locker.LockerIds...)
		}
	}
	return &types.QueryLockersByProductToAssetIDResponse{
		LockerIds: lockerIds,
	}, nil
}

func (q *queryServer) QueryLockerInfoByProductID(c context.Context, request *types.QueryLockerInfoByProductIDRequest) (*types.QueryLockerInfoByProductIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, found := q.GetLockerLookupTable(ctx, app.Id)

	if !found {
		return nil, status.Errorf(codes.NotFound, "no asset exists appID %d", app.Id)
	}

	var lockerIds []string
	for _, locker := range lockerLookupData.Lockers {
		lockerIds = locker.LockerIds
	}
	return &types.QueryLockerInfoByProductIDResponse{
		LockerIds: lockerIds,
	}, nil

}

func (q *queryServer) QueryTotalDepositByProductAssetID(c context.Context, request *types.QueryTotalDepositByProductAssetIDRequest) (*types.QueryTotalDepositByProductAssetIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, found := q.GetLockerLookupTable(ctx, app.Id)

	if !found {
		return nil, status.Errorf(codes.NotFound, "no asset exists appID %d", app.Id)
	}

	var totalDeposit uint64
	for _, locker := range lockerLookupData.Lockers {

		if request.AssetId == locker.AssetId {
			totalDeposit += locker.DepositedAmount.Uint64()
		}
	}
	return &types.QueryTotalDepositByProductAssetIDResponse{
		TotalDeposit: totalDeposit,
	}, nil
}

func (q *queryServer) QueryOwnerLockerByProductID(c context.Context, request *types.QueryOwnerLockerByProductIDRequest) (*types.QueryOwnerLockerByProductIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, _ := q.GetLockerLookupTable(ctx, app.Id)

	var lockerIds []string
	for _, locker := range lockerLookupData.Lockers {

		for _, lockerID := range locker.LockerIds {
			locker1, _ := q.GetLocker(ctx, lockerID)
			if request.Owner == locker1.Depositor {
				lockerIds = append(lockerIds, locker1.LockerId)
			}
		}
	}

	return &types.QueryOwnerLockerByProductIDResponse{
		LockerIds: lockerIds,
	}, nil
}

func (q *queryServer) QueryOwnerLockerOfAllProduct(c context.Context, request *types.QueryOwnerLockerOfAllProductRequest) (*types.QueryOwnerLockerOfAllProductResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	userlockerLookupData, _ := q.GetUserLockerAssetMapping(ctx, request.Owner)

	var lockerIds []string
	for _, locker := range userlockerLookupData.LockerAppMapping {
		for _, data := range locker.UserAssetLocker{
			lockerIds = append(lockerIds, data.LockerId)
		}
	}

	return &types.QueryOwnerLockerOfAllProductResponse{
		LockerIds: lockerIds,
	}, nil
}

func (q *queryServer) QueryOwnerLockerByProductToAssetID(c context.Context, request *types.QueryOwnerLockerByProductToAssetIDRequest) (*types.QueryOwnerLockerByProductToAssetIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, found := q.GetLockerLookupTable(ctx, app.Id)

	if !found {
		return nil, status.Errorf(codes.NotFound, "no asset exists appID %d", app.Id)
	}

	var lockerInfos []types.Locker
	for _, locker := range lockerLookupData.Lockers {

		if request.AssetId == locker.AssetId {
			for _, lockerID := range locker.LockerIds {
				locker1, _ := q.GetLocker(ctx, lockerID)
				if request.Owner == locker1.Depositor {
					lockerInfos = append(lockerInfos, locker1)
				}
			}
		}
	}

	return &types.QueryOwnerLockerByProductToAssetIDResponse{
		LockerInfo: lockerInfos,
	}, nil
}

func (q *queryServer) QueryLockerCountByProductID(c context.Context, request *types.QueryLockerCountByProductIDRequest) (*types.QueryLockerCountByProductIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, found := q.GetLockerLookupTable(ctx, app.Id)

	if !found {
		return nil, status.Errorf(codes.NotFound, "no asset exists appID %d", app.Id)
	}

	var lockerCount uint64
	for _, locker := range lockerLookupData.Lockers {
		for _, lockerID := range locker.LockerIds {
			_, lockerFound := q.GetLocker(ctx, lockerID)
			if lockerFound {
				lockerCount += 1
			}
		}
	}
	return &types.QueryLockerCountByProductIDResponse{
		TotalCount: lockerCount,
	}, nil

}

func (q *queryServer) QueryLockerCountByProductToAssetID(c context.Context, request *types.QueryLockerCountByProductToAssetIDRequest) (*types.QueryLockerCountByProductToAssetIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, found := q.GetLockerLookupTable(ctx, app.Id)

	if !found {
		return nil, status.Errorf(codes.NotFound, "no asset exists appID %d", app.Id)
	}

	var lockerCount uint64
	for _, locker := range lockerLookupData.Lockers {
		if request.AssetId == locker.AssetId {
			for _, lockerID := range locker.LockerIds {
				_, lockerFound := q.GetLocker(ctx, lockerID)
				if lockerFound {
					lockerCount += 1
				}
			}
		}
	}
	return &types.QueryLockerCountByProductToAssetIDResponse{
		TotalCount: lockerCount,
	}, nil
}

func (q *queryServer) QueryWhiteListedAssetIDsByProductID(c context.Context, request *types.QueryWhiteListedAssetIDsByProductIDRequest) (*types.QueryWhiteListedAssetIDsByProductIDResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)

	app, found := q.GetApp(ctx, request.ProductId)

	if !found {
		return nil, status.Errorf(codes.NotFound, "app does not exist for appID %d", request.ProductId)
	}

	lockerLookupData, found := q.GetLockerLookupTable(ctx, app.Id)

	if !found {
		return nil, status.Errorf(codes.NotFound, "no asset exists appID %d", app.Id)
	}

	var assetIds []uint64
	for _, locker := range lockerLookupData.Lockers {
		assetIds = append(assetIds, locker.AssetId)
	}

	return &types.QueryWhiteListedAssetIDsByProductIDResponse{
		AssetIds: assetIds,
	}, nil
}

func (q *queryServer) QueryWhiteListedAssetByAllProduct(c context.Context, request *types.QueryWhiteListedAssetByAllProductRequest) (*types.QueryWhiteListedAssetByAllProductResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	var (
		ctx = sdk.UnwrapSDKContext(c)
	)
	apps, found := q.asset.GetApps(ctx)

	if !found {
		return nil, status.Errorf(codes.NotFound, "ano apps exist")
	}

	var productToAll []types.ProductToAllAsset

	for _, app := range apps {
		var product types.ProductToAllAsset
		var assets []assettypes.Asset
		appData, _ := q.GetLockerProductAssetMapping(ctx, app.Id)
		for _, assetId := range appData.AssetIds {
			asset, assetFound := q.asset.GetAsset(ctx, assetId)
			if assetFound {
				assets = append(assets, asset)
			}
			product = types.ProductToAllAsset{
				ProductId: appData.AppMappingId,
				Assets:    assets,
			}
		}

		productToAll = append(productToAll, product)
	}
	return &types.QueryWhiteListedAssetByAllProductResponse{
		ProductToAllAsset: productToAll,
	}, nil
}