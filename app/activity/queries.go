package activity

import (
	"context"
	"encoding/json"

	app "phoenix-api/app"
	types "phoenix-api/types/activity"

	"google.golang.org/protobuf/types/known/structpb"
)

func (k Keeper) ActivityAtNft(ctx context.Context, request *types.ActivityAtNftRequest) (*types.ActivityAtNftResponse, error) {
	var activities []*types.Activity
	var res []*types.ActivityInfo

	query := k.dbHandler.Table(app.ACTIVITY_TABLE).
		Where("nid = ?", request.NftId)

	err := query.Find(&activities).Error
	if err != nil {
		return &types.ActivityAtNftResponse{}, err
	}

	for _, activity := range activities {
		detail := &structpb.Struct{}
		if err := json.Unmarshal([]byte(activity.Details), detail); err != nil {
			return &types.ActivityAtNftResponse{}, err
		}
		res = append(res, &types.ActivityInfo{
			Type:      activity.Type,
			Timestamp: activity.Details,
			Details:   detail,
		})
	}

	return &types.ActivityAtNftResponse{
		Activity: res,
	}, nil
}
