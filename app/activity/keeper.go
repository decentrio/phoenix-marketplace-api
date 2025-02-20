package activity

import (
	"gorm.io/gorm"

	types "phoenix-api/types/activity"
)

type Keeper struct {
	dbHandler *gorm.DB
	types.UnimplementedActivityQueryServer
}

func NewKeeper(db *gorm.DB) *Keeper {
	return &Keeper{
		dbHandler: db,
	}
}

var _ types.ActivityQueryServer = Keeper{}

