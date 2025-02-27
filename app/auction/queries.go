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
