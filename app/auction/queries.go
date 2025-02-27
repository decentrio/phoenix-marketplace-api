package auction

import (
	"context"

	app "phoenix-marketplace-api/app"
	dbtypes "phoenix-marketplace-api/database"
	types "phoenix-marketplace-api/types/auction"
)

func (k Keeper) AuctionsAvailable(ctx context.Context, request *types.AuctionsAvailableRequest) (*types.AuctionsAvailableResponse, error) {
	var auctions []*dbtypes.Auction
	var auctionInfos []*types.AuctionInfo

	query := k.dbHandler.Table(app.AUCTION_TABLE).
		Where("status = ?", "active")

	err := query.Find(&auctions).Error
	if err != nil {
		return &types.AuctionsAvailableResponse{}, err
	}

	for _, auction := range auctions {
		var collection string
		if err = k.dbHandler.Table(app.NFT_TABLE).
			Select("collection").
			Where("nid = ?", auction.Nid).First(&collection).Error; err == nil {
			auctionInfos = append(auctionInfos, &types.AuctionInfo{
				AuctionId:  string(auction.ID),
				NftId:      auction.Nid,
				Collection: collection,
				CurrentBid: float32(auction.HighestBid),
				EndTime:    auction.EndTime,
			})
		}

	}

	return &types.AuctionsAvailableResponse{
		Auctions: auctionInfos,
	}, nil
}
