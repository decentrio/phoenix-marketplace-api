package offer

import (
	"context"

	app "phoenix-api/app"
	types "phoenix-api/types/offer"
)

func (k Keeper) OffersAtNft(ctx context.Context, request *types.OffersAtNftRequest) (*types.OffersAtNftResponse, error) {
	var data []*types.Offer

	query := k.dbHandler.Table(app.OFFER_TABLE).
		Where("nft_id = ?", request.NftId)

	err := query.Find(&data).Error
	if err != nil {
		return &types.OffersAtNftResponse{}, err
	}

	return &types.OffersAtNftResponse{
		Data: data,
	}, nil
}
