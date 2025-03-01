package auction

import (
	"context"
	"strconv"
	"strings"

	app "phoenix-marketplace-api/app"
	dbtypes "phoenix-marketplace-api/database"
	types "phoenix-marketplace-api/types/auction"
)

func (k Keeper) AuctionsAvailable(ctx context.Context, request *types.AuctionsAvailableRequest) (*types.AuctionsAvailableResponse, error) {
	var auctions []*dbtypes.Auction
	var auctionInfos []*types.AuctionInfo

	query := k.dbHandler.Table(app.AUCTION_TABLE).
		Where("status = ?", "Active")

	err := query.Find(&auctions).Error
	if err != nil {
		return &types.AuctionsAvailableResponse{}, err
	}

	for _, auction := range auctions {
		splitNid := strings.Split(auction.Nid, "@")
		auctionInfos = append(auctionInfos, &types.AuctionInfo{
			AuctionId:  strconv.FormatUint(auction.ID, 10),
			NftId:      splitNid[1],
			Collection: splitNid[0],
			CurrentBid: float32(auction.HighestBid) / app.DIVIDE_DECIMALS,
			EndTime:    auction.EndTime,
		})
	}

	return &types.AuctionsAvailableResponse{
		Auctions: auctionInfos,
	}, nil
}

func (k Keeper) Auctions(ctx context.Context, request *types.AuctionsRequest) (*types.AuctionsResponse, error) {
	var auctions []*dbtypes.Auction
	var auctionInfos []*types.AuctionInfo

	query := k.dbHandler.Table(app.AUCTION_TABLE)
	if request.Status != "" {
		query = query.Where("status = ?", request.Status)
	}

	err := query.Find(&auctions).Error
	if err != nil {
		return &types.AuctionsResponse{}, err
	}

	for _, auction := range auctions {
		splitNid := strings.Split(auction.Nid, "@")
		auctionInfos = append(auctionInfos, &types.AuctionInfo{
			AuctionId:  strconv.FormatUint(auction.ID, 10),
			NftId:      splitNid[1],
			Collection: splitNid[0],
			CurrentBid: float32(auction.HighestBid) / app.DIVIDE_DECIMALS,
			EndTime:    auction.EndTime,
		})
	}

	return &types.AuctionsResponse{
		Auctions: auctionInfos,
	}, nil
}

func (k Keeper) AuctionDetail(ctx context.Context, request *types.AuctionDetailRequest) (*types.AuctionDetailResponse, error) {
	var auction dbtypes.Auction
	err := k.dbHandler.Table(app.AUCTION_TABLE).Where("id = ?", request.AuctionId).First(&auction).Error
	if err != nil {
		return &types.AuctionDetailResponse{}, err
	}

	var bids []*types.Bid
	err = k.dbHandler.Table(app.BID_TABLE).
		Select("bidder", "bid_amount").
		Where("auction_id = ?", request.AuctionId).
		Order("timestamp DESC").
		Find(&bids).Error
	if err != nil {
		return &types.AuctionDetailResponse{}, err
	}
	splitNid := strings.Split(auction.Nid, "@")
	return &types.AuctionDetailResponse{
		Data: &types.AuctionDetail{
			AuctionId:    request.AuctionId,
			NftId:        splitNid[1],
			Collection:   splitNid[0],
			CurrentBid:   float32(auction.HighestBid) / app.DIVIDE_DECIMALS,
			EndTime:      auction.EndTime,
			Amount:       auction.Amount,
			AuctionToken: auction.AuctionToken,
			Status:       auction.Status,
			BuyNowPrice:  float32(auction.BuyNowPrice) / app.DIVIDE_DECIMALS,
			MinimumPrice: float32(auction.MinimumPrice) / app.DIVIDE_DECIMALS,
			Bids:         bids,
		},
	}, nil
}
