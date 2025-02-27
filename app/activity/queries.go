package activity

import (
	"context"
	"encoding/json"
	"strconv"

	app "phoenix-marketplace-api/app"
	dbtypes "phoenix-marketplace-api/database"
	types "phoenix-marketplace-api/types/activity"

	"google.golang.org/protobuf/types/known/structpb"
)

func (k Keeper) ActivityAtNft(ctx context.Context, request *types.ActivityAtNftRequest) (*types.ActivityAtNftResponse, error) {
	var activities []*dbtypes.Activity
	var res []*types.ActivityInfo
	nid := request.CollectionAddress + "@" + request.NftId
	query := k.dbHandler.Table(app.ACTIVITY_TABLE).
		Where("nid = ?", nid)

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
			Timestamp: strconv.FormatUint(activity.Timestamp, 10),
			Details:   detail,
		})
	}

	return &types.ActivityAtNftResponse{
		Activity: res,
	}, nil
}
