package nft

import (
	"context"
	"encoding/json"

	app "phoenix-api/app"
	types "phoenix-api/types/nft"

	"google.golang.org/protobuf/types/known/structpb"
)

func (k Keeper) NftsAtAccount(ctx context.Context, request *types.NftsAtAccountRequest) (*types.NftsAtAccountResponse, error) {
	var data []*types.NftAtAccount
	var nftIds []string

	err := k.dbHandler.Table(app.BALANCE_TABLE).
		Select("nid").
		Where("account = ?", request.AccountId).
		Find(&nftIds).Error

	if err != nil {
		return &types.NftsAtAccountResponse{}, err
	}

	for _, nftId := range nftIds {
		var nft types.Nft
		if err = k.dbHandler.Table(app.NFT_TABLE).
			Where("nid = ?", nftId).First(&nft).Error; err == nil {
			metadata := &structpb.Struct{}
			if err := json.Unmarshal([]byte(nft.Uri), metadata); err != nil {
				return &types.NftsAtAccountResponse{}, err
			}

			data = append(data, &types.NftAtAccount{
				Id:         nft.Id,
				Name:       nft.Name,
				Collection: nft.Collection,
				Metadata:   metadata,
			})
		}
	}

	return &types.NftsAtAccountResponse{
		Nfts: data,
	}, nil
}

func (k Keeper) NftsAvailable(ctx context.Context, request *types.NftsAvailableRequest) (*types.NftsAvailableResponse, error) {
	var nftInfos []*types.NftAvailable
	var nfts []*types.Nft

	err := k.dbHandler.Table(app.NFT_TABLE).
		Find(&nfts).Error

	if err != nil {
		return &types.NftsAvailableResponse{}, err
	}

	for _, nft := range nfts {
		var price uint64
		if err = k.dbHandler.Table(app.PRICE_TABLE).
			Select("price").
			Where("nid = ?", nft.Id).Order("tx_time DESC").First(&price).Error; err == nil {
			nftInfos = append(nftInfos, &types.NftAvailable{
				Id:         nft.Id,
				Name:       nft.Name,
				Collection: nft.Collection,
				Price:      float32(price),
			})
		}
	}

	return &types.NftsAvailableResponse{
		Nfts: nftInfos,
	}, nil
}

func (k Keeper) NftsPopular(ctx context.Context, request *types.NftsPopularRequest) (*types.NftsPopularResponse, error) {
	var nftInfos []*types.NftPopular
	var volumes []*types.Volume

	err := k.dbHandler.Table(app.VOLUME_TABLE).
		Order("volume DESC").Limit(10).
		Find(&volumes).Error

	if err != nil {
		return &types.NftsPopularResponse{}, err
	}

	for _, volume := range volumes {
		nftInfos = append(nftInfos, &types.NftPopular{
			Id:          volume.Nid,
			Name:        volume.Token,
			Collection:  volume.Collection,
			TradeVolume: float32(volume.Volume),
		})
	}

	return &types.NftsPopularResponse{Nfts: nftInfos}, nil
}

func (k Keeper) PriceHistory(ctx context.Context, request *types.PriceHistoryRequest) (*types.PriceHistoryResponse, error) {
	var prices []*types.PriceHistory
	var priceInfos []*types.PriceHistoryInfo

	err := k.dbHandler.Table(app.PRICE_TABLE).
		Where("nid = ?", request.NftId).
		Find(&prices).Error

	if err != nil {
		return &types.PriceHistoryResponse{}, err
	}

	for _, price := range prices {
		priceInfos = append(priceInfos, &types.PriceHistoryInfo{
			Timestamp: price.TxTime,
			Price:     float32(price.Price),
		})
	}

	return &types.PriceHistoryResponse{
		PriceHistory: priceInfos,
	}, nil
}
