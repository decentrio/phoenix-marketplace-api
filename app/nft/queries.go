package nft

import (
	"context"

	app "phoenix-api/app"
	types "phoenix-api/types/nft"
)

func (k Keeper) NftsAtAccount(ctx context.Context, request *types.NftsAtAccountRequest) (*types.NftsAtAccountResponse, error) {
	var data []*types.Nft

	query := k.dbHandler.Table(app.NFT_TABLE).
		Where("owner = ?", request.AccountId)

	err := query.Find(&data).Error
	if err != nil {
		return &types.NftsAtAccountResponse{}, err
	}

	return &types.NftsAtAccountResponse{
		Data: data,
	}, nil
}
