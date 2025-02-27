package auction

import (
	"gorm.io/gorm"

	types "phoenix-marketplace-api/types/auction"
)

type Keeper struct {
	dbHandler *gorm.DB
	types.UnimplementedAuctionQueryServer
}

func NewKeeper(db *gorm.DB) *Keeper {
	return &Keeper{
		dbHandler: db,
	}
}

var _ types.AuctionQueryServer = Keeper{}
