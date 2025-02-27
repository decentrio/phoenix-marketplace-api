package collection

import (
	"context"

	app "phoenix-api/app"
	types "phoenix-api/types/collection"
	nftTypes "phoenix-api/types/nft"
)

func (k Keeper) CollectionsAvailable(ctx context.Context, request *types.CollectionsAvailableRequest) (*types.CollectionsAvailableResponse, error) {
	var collectionInfos []*types.CollectionAvailable
	var collections []*types.Collection

	err := k.dbHandler.Table(app.COLLECTION_TABLE).
		Find(&collections).Error

	if err != nil {
		return &types.CollectionsAvailableResponse{}, err
	}

	for _, collection := range collections {
		collectionInfos = append(collectionInfos, &types.CollectionAvailable{
			Id:          collection.Address,
			Name:        collection.Name,
			TotalItems:  collection.TotalItems,
			Description: collection.Uri,
		})
	}

	return &types.CollectionsAvailableResponse{
		Collections: collectionInfos,
	}, nil
}

func (k Keeper) CollectionsPopular(ctx context.Context, request *types.CollectionsPopularRequest) (*types.CollectionsPopularResponse, error) {
	var collectionInfos []*types.CollectionPopular
	var volumes []*nftTypes.Volume

	err := k.dbHandler.Table(app.VOLUME_TABLE).
		Order("volume DESC").Limit(10).
		Find(&volumes).Error

	if err != nil {
		return &types.CollectionsPopularResponse{}, err
	}

	for _, volume := range volumes {
		var collection types.Collection
		if err = k.dbHandler.Table(app.COLLECTION_TABLE).
			Where("address = ?", volume.Collection).First(&collection).Error; err == nil {
			collectionInfos = append(collectionInfos, &types.CollectionPopular{
				Id:          volume.Collection,
				Name:        collection.Name,
				TradeVolume: float32(volume.Volume),
				TotalItems:  collection.TotalItems,
			})
		}

	}

	return &types.CollectionsPopularResponse{
		Collections: collectionInfos,
	}, nil
}
