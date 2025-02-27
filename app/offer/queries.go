package offer

import (
	"context"

	app "phoenix-marketplace-api/app"
	dbtypes "phoenix-marketplace-api/database"
	auctionTypes "phoenix-marketplace-api/types/auction"
	nftTypes "phoenix-marketplace-api/types/nft"
	types "phoenix-marketplace-api/types/offer"
)

func (k Keeper) OffersAtNft(ctx context.Context, request *types.OffersAtNftRequest) (*types.OffersAtNftResponse, error) {
	var data []*types.Offer
	var auctions []*dbtypes.Auction

	query := k.dbHandler.Table(app.AUCTION_TABLE).
		Where("status = ?", "active").
		Where("nid = ?", request.NftId)

	err := query.Find(&auctions).Error
	if err != nil {
		return &types.OffersAtNftResponse{}, err
	}

	for _, auction := range auctions {
		var bids []*nftTypes.Bid
		if err = k.dbHandler.Table(app.BID_TABLE).
			Where("auction_id = ?", auction.Id).First(&bids).Error; err == nil {
			for _, bid := range bids {
				data = append(data, &types.Offer{
					OfferId: bid.Hash,
					Price:   float32(bid.BidAmount),
					Bidder:  bid.Bidder,
				})
			}
		}
	}

	return &types.OffersAtNftResponse{
		Offers: data,
	}, nil
}
