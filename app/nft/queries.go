package nft

import (
	"context"
	"encoding/json"
	"fmt"

	app "phoenix-marketplace-api/app"
	dbtypes "phoenix-marketplace-api/database"
	types "phoenix-marketplace-api/types/nft"

	"google.golang.org/protobuf/types/known/structpb"
)

func (k Keeper) NftsAtAccount(ctx context.Context, request *types.NftsAtAccountRequest) (*types.NftsAtAccountResponse, error) {
	// Validate the request
	if request.AccountId == "" {
		return nil, fmt.Errorf("account_id cannot be empty")
	}

	var balances []dbtypes.Balance
	err := k.dbHandler.Table(app.BALANCE_TABLE).
		Where("account = ?", request.AccountId).
		Find(&balances).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query balances: %v", err)
	}

	nids := make([]string, 0, len(balances))
	for _, balance := range balances {
		nids = append(nids, balance.Nid)
	}

	var nfts []dbtypes.NFT
	err = k.dbHandler.Table(app.NFT_TABLE).
		Where("nid IN ?", nids).
		Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query NFTs: %v", err)
	}

	response := &types.NftsAtAccountResponse{
		Nfts: make([]*types.Nft, 0, len(nfts)),
	}

	for _, nft := range nfts {
		metadata := &structpb.Struct{}
		err := json.Unmarshal(nft.Metadata, metadata)
		if err != nil {
			fmt.Printf("failed to unmarshall metadata: %v", err)
		}
		response.Nfts = append(response.Nfts, &types.Nft{
			Id:         nft.ID,
			Collection: nft.Collection,
			Metadata:   metadata,
		})
	}

	return response, nil
}

func (k Keeper) NftsAvailable(ctx context.Context, request *types.NftsAvailableRequest) (*types.NftsAvailableResponse, error) {
	var auctions []dbtypes.Auction
	query := k.dbHandler.Table(app.AUCTION_TABLE).
		Where("status = ?", "Active")

	err := query.Find(&auctions).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query auctions: %v", err)
	}

	nidSet := make(map[string]struct{})
	for _, auction := range auctions {
		nidSet[auction.Nid] = struct{}{}
	}

	// Convert the set of nids to a slice
	nids := make([]string, 0, len(nidSet))
	for nid := range nidSet {
		nids = append(nids, nid)
	}

	var nfts []dbtypes.NFT
	nftQuery := k.dbHandler.Table(app.NFT_TABLE).
		Where("nid IN ?", nids)

	if request.CollectionId != "" {
		nftQuery = nftQuery.Where("collection = ?", request.CollectionId)
	}

	err = nftQuery.Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query NFTs: %v", err)
	}

	response := &types.NftsAvailableResponse{
		Nfts: make([]*types.NftAvailable, 0, len(nfts)),
	}
	for _, nft := range nfts {
		price := dbtypes.Price{}
		err = k.dbHandler.Table(app.PRICE_TABLE).Where("nid = ?", nft.Nid).Order("tx_time DESC").First(&price).Error
		if err != nil {
			return nil, fmt.Errorf("failed to query NFT price: %v", err)
		}
		response.Nfts = append(response.Nfts, &types.NftAvailable{
			Id:         nft.ID,
			Collection: nft.Collection,
			Price:      float32(price.Price) / app.DIVIDE_DECIMALS,
		})
	}

	return response, nil
}

func (k Keeper) NftsPopular(ctx context.Context, request *types.NftsPopularRequest) (*types.NftsPopularResponse, error) {
	// Step 1: Query the volume table to get the latest 1000 records
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
		volumeMap[volume.Nid] += float32(volume.Volume) / app.DIVIDE_DECIMALS
	}

	nids := make([]string, 0, len(volumeMap))
	for nid := range volumeMap {
		nids = append(nids, nid)
	}

	// Step 4: Query the NFT table to get details for each nid
	var nfts []dbtypes.NFT
	err = k.dbHandler.Table(app.NFT_TABLE).
		Where("nid IN ?", nids).
		Find(&nfts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to query NFTs: %v", err)
	}

	// Step 5: Map the NFT details and trade volumes to the response format
	response := &types.NftsPopularResponse{
		Nfts: make([]*types.NftPopular, 0, len(nfts)),
	}

	for _, nft := range nfts {
		// Get the total trade volume from the volume map
		tradeVolume := volumeMap[nft.Nid]

		// Add the NFT to the response
		response.Nfts = append(response.Nfts, &types.NftPopular{
			Id:          nft.ID,
			Collection:  nft.Collection,
			TradeVolume: tradeVolume,
		})
	}

	return response, nil
}

func (k Keeper) PriceHistory(ctx context.Context, request *types.PriceHistoryRequest) (*types.PriceHistoryResponse, error) {
	var prices []*dbtypes.Price
	var priceInfos []*types.PriceHistoryInfo
	nid := request.CollectionAddress + "@" + request.NftId
	err := k.dbHandler.Table(app.PRICE_TABLE).
		Where("nid = ?", nid).
		Order("tx_time DESC").
		Find(&prices).Error

	if err != nil {
		return &types.PriceHistoryResponse{}, err
	}

	for _, price := range prices {
		priceInfos = append(priceInfos, &types.PriceHistoryInfo{
			Timestamp: price.TxTime,
			Price:     float32(price.Price) / app.DIVIDE_DECIMALS,
		})
	}

	return &types.PriceHistoryResponse{
		PriceHistory: priceInfos,
	}, nil
}

func (k Keeper) AllNfts(ctx context.Context, request *types.AllNftsRequest) (*types.AllNftsResponse, error) {
	var nfts []*dbtypes.NFT
	err := k.dbHandler.Table(app.NFT_TABLE).
		Find(&nfts).Error
	if err != nil {
		return &types.AllNftsResponse{}, err
	}

	response := &types.AllNftsResponse{}
	for _, nft := range nfts {
		metadata := &structpb.Struct{}
		err := json.Unmarshal(nft.Metadata, metadata)
		if err != nil {
			fmt.Printf("failed to unmarshall metadata: %v", err)
		}
		response.Data = append(response.Data, &types.Nft{
			Id:         nft.ID,
			Collection: nft.Collection,
			Metadata:   metadata,
		})
	}
	return response, nil
}

func (k Keeper) Nfts(ctx context.Context, request *types.NftsRequest) (*types.NftsResponse, error) {
	var nfts []*dbtypes.NFT
	err := k.dbHandler.Table(app.NFT_TABLE).Where("collection = ?", request.Collection).
		Find(&nfts).Error
	if err != nil {
		return &types.NftsResponse{}, err
	}

	response := &types.NftsResponse{}
	for _, nft := range nfts {
		metadata := &structpb.Struct{}
		err := json.Unmarshal(nft.Metadata, metadata)
		if err != nil {
			fmt.Printf("failed to unmarshall metadata: %v", err)
		}
		response.Data = append(response.Data, &types.Nft{
			Id:         nft.ID,
			Collection: nft.Collection,
			Metadata:   metadata,
		})
	}
	return response, nil
}

func (k Keeper) NftVolume(ctx context.Context, request *types.NftVolumeRequest) (*types.NftVolumeResponse, error) {
	var vols []*types.VolumnInfo
	nid := request.Collection + "@" + request.Id
	err := k.dbHandler.Table(app.VOLUME_TABLE).
		Select("amount", "volume", "token", "timestamp").
		Where("nid = ?", nid).
		Order("timestamp DESC").
		Find(&vols).Error
	if err != nil {
		return &types.NftVolumeResponse{}, err
	}

	return &types.NftVolumeResponse{
		Volumes: vols,
	}, nil
}

func (k Keeper) NftDetail(ctx context.Context, request *types.NftDetailRequest) (*types.NftDetailResponse, error) {
	var nft dbtypes.NFT
	nid := request.Collection + "@" + request.Id
	err := k.dbHandler.Table(app.NFT_TABLE).
		Where("nid = ?", nid).
		First(&nft).Error
	if err != nil {
		return &types.NftDetailResponse{}, err
	}
	metadata := &structpb.Struct{}
	err = json.Unmarshal(nft.Metadata, metadata)
	if err != nil {
		fmt.Printf("failed to unmarshall metadata: %v", err)
	}
	return &types.NftDetailResponse{
		Data: &types.Nft{
			Id:         request.Id,
			Collection: request.Collection,
			Metadata:   metadata,
		},
	}, nil
}
