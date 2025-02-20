package offer

import (
	"gorm.io/gorm"

	types "phoenix-api/types/offer"
)

type Keeper struct {
	dbHandler *gorm.DB
	types.UnimplementedOfferQueryServer
}

func NewKeeper(db *gorm.DB) *Keeper {
	return &Keeper{
		dbHandler: db,
	}
}

var _ types.OfferQueryServer = Keeper{}

