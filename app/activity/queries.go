package activity

import (
	"context"

	app "phoenix-api/app"
	types "phoenix-api/types/activity"
)

func (k Keeper) ActivityAtNft(ctx context.Context, request *types.ActivityAtNftRequest) (*types.ActivityAtNftResponse, error) {
	var data []*types.Activity

	query := k.dbHandler.Table(app.ACTIVITY_TABLE).
		Where("nft_id = ?", request.NftId)

	err := query.Find(&data).Error
	if err != nil {
		return &types.ActivityAtNftResponse{}, err
	}

	return &types.ActivityAtNftResponse{
		Data: data,
	}, nil
}
