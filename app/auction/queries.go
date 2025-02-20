package auction

import (
	"context"
	"time"

	app "phoenix-api/app"
	types "phoenix-api/types/auction"
)

func (k Keeper) AuctionsAvailable(ctx context.Context, request *types.AuctionsAvailableRequest) (*types.AuctionsAvailableResponse, error) {
	var data []*types.Auction
	now := time.Now()
	query := k.dbHandler.Table(app.AUCTION_TABLE).
		Where("end_time > ?", now)

	err := query.Find(&data).Error
	if err != nil {
		return &types.AuctionsAvailableResponse{}, err
	}

	return &types.AuctionsAvailableResponse{
		Data: data,
	}, nil
}
