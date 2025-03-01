package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	app "phoenix-marketplace-api/app"
	dbtypes "phoenix-marketplace-api/database"
	types "phoenix-marketplace-api/types/collection"

	"google.golang.org/protobuf/types/known/structpb"
)

func (k Keeper) CollectionsAvailable(ctx context.Context, request *types.CollectionsAvailableRequest) (*types.CollectionsAvailableResponse, error) {
	collMap := make(map[string]bool)
	collVec := []string{}
	var auctions []*dbtypes.Auction
	var collectionInfos []*types.Collection
	var collections []*dbtypes.Collection

	err := k.dbHandler.Table(app.AUCTION_TABLE).
		Find(&auctions).Error
	if err != nil {
		return &types.CollectionsAvailableResponse{}, err
	}

	for _, auction := range auctions {
		collMap[strings.Split(auction.Nid, "@")[0]] = true
	}
	for collAddr := range collMap {
		collVec = append(collVec, collAddr)
	}

	err = k.dbHandler.Table(app.COLLECTION_TABLE).Where("address IN ?", collVec).
		Find(&collections).Error
	if err != nil {
		return &types.CollectionsAvailableResponse{}, err
	}

	for _, collection := range collections {
		metadata := &structpb.Struct{}
		err := json.Unmarshal(collection.Metadata, metadata)
		if err != nil {
			return &types.CollectionsAvailableResponse{}, err
		}
		collectionInfos = append(collectionInfos, &types.Collection{
			Id:          collection.Address,
			Name:        collection.Symbol,
			Description: collection.Name,
			Metadata:    metadata,
			TotalItems:  uint32(collection.TotalItems),
		})
	}

	return &types.CollectionsAvailableResponse{
		Collections: collectionInfos,
	}, nil
}

func (k Keeper) CollectionsPopular(ctx context.Context, request *types.CollectionsPopularRequest) (*types.CollectionsPopularResponse, error) {
	var volumes []dbtypes.Volume
	err := k.dbHandler.Table(app.VOLUME_TABLE).
		Order("timestamp DESC").
		Limit(1000).
		Find(&volumes).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query volumes: %v", err)
	}

	volumeMap := make(map[string]float32)
	for _, volume := range volumes {
		volumeMap[volume.Collection] += float32(volume.Volume) / app.DIVIDE_DECIMALS
	}

	collections := make([]string, 0, len(volumeMap))
	for collection := range volumeMap {
		collections = append(collections, collection)
	}

	var collectionInfos []dbtypes.Collection
	err = k.dbHandler.Table(app.COLLECTION_TABLE).
		Where("address IN ?", collections).
		Find(&collectionInfos).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query collections: %v", err)
	}

	response := &types.CollectionsPopularResponse{
		Collections: make([]*types.CollectionPopular, 0, len(collectionInfos)),
	}

	for _, collection := range collectionInfos {
		// Get the total trade volume from the volume map
		tradeVolume := volumeMap[collection.Address]

		// Add the collection to the response
		response.Collections = append(response.Collections, &types.CollectionPopular{
			Id:          collection.Address,
			Name:        collection.Name,
			TradeVolume: tradeVolume,
			TotalItems:  uint32(collection.TotalItems),
		})
	}

	return response, nil
}

func (k Keeper) Collections(ctx context.Context, request *types.CollectionsRequest) (*types.CollectionsResponse, error) {
	var collections []*dbtypes.Collection
	err := k.dbHandler.Table(app.COLLECTION_TABLE).
		Find(&collections).Error
	if err != nil {
		return &types.CollectionsResponse{}, err
	}

	response := &types.CollectionsResponse{}

	for _, collection := range collections {
		metadata := &structpb.Struct{}
		err := json.Unmarshal(collection.Metadata, metadata)
		if err != nil {
			return &types.CollectionsResponse{}, err
		}
		response.Collections = append(response.Collections, &types.Collection{
			Id:          collection.Address,
			Name:        collection.Symbol,
			Description: collection.Name,
			Metadata:    metadata,
			TotalItems:  uint32(collection.TotalItems),
		})
	}
	return response, nil
}

func (k Keeper) CollectionDetail(ctx context.Context, request *types.CollectionDetailRequest) (*types.CollectionDetailResponse, error) {
	var collection dbtypes.Collection
	err := k.dbHandler.Table(app.COLLECTION_TABLE).
		First(&collection).Error
	if err != nil {
		return &types.CollectionDetailResponse{}, err
	}
	metadata := &structpb.Struct{}
	err = json.Unmarshal(collection.Metadata, metadata)
	if err != nil {
		return &types.CollectionDetailResponse{}, err
	}
	response := &types.CollectionDetailResponse{
		Data: &types.CollectionDetail{
			Id:          collection.Address,
			Name:        collection.Symbol,
			Description: collection.Name,
			Metadata:    metadata,
			TotalItems:  uint32(collection.TotalItems),
		},
	}

	var nfts []*dbtypes.NFT
	err = k.dbHandler.Table(app.NFT_TABLE).
		Find(&nfts).Error
	if err != nil {
		return &types.CollectionDetailResponse{}, err
	}
	for _, nft := range nfts {
		metadata := &structpb.Struct{}
		err := json.Unmarshal(nft.Metadata, metadata)
		if err != nil {
			return &types.CollectionDetailResponse{}, err
		}
		response.Data.Nfts = append(response.Data.Nfts, &types.NFT{
			Nid:      nft.Nid,
			Id:       nft.ID,
			Metadata: metadata,
		})
	}
	return response, nil
}
