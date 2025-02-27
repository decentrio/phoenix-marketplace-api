package offer

import (
	"context"

	app "phoenix-marketplace-api/app"
	dbtypes "phoenix-marketplace-api/database"
	types "phoenix-marketplace-api/types/offer"
)

func (k Keeper) OffersAtNft(ctx context.Context, request *types.OffersAtNftRequest) (*types.OffersAtNftResponse, error) {
	var data []*types.Offer
	var auctions []*dbtypes.Auction
	nid := request.CollectionAddress + "@" + request.NftId
	query := k.dbHandler.Table(app.AUCTION_TABLE).
		Where("status = ?", "Active").
		Where("nid = ?", nid)

	err := query.Find(&auctions).Error
	if err != nil {
		return &types.OffersAtNftResponse{}, err
	}

	for _, auction := range auctions {
		var bids []*dbtypes.Bid
		if err = k.dbHandler.Table(app.BID_TABLE).
			Where("auction_id = ?", auction.ID).Find(&bids).Error; err == nil {
			for _, bid := range bids {
				data = append(data, &types.Offer{
					OfferId: bid.Hash,
					Price:   float32(bid.BidAmount) / app.DIVIDE_DECIMALS,
					Bidder:  bid.Bidder,
				})
			}
		}
	}

	return &types.OffersAtNftResponse{
		Offers: data,
	}, nil
}
