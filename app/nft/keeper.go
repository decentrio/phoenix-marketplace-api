package nft

import (
	"gorm.io/gorm"

	types "phoenix-marketplace-api/types/nft"
)

type Keeper struct {
	dbHandler *gorm.DB
	types.UnimplementedNftQueryServer
}

func NewKeeper(db *gorm.DB) *Keeper {
	return &Keeper{
		dbHandler: db,
	}
}

var _ types.NftQueryServer = Keeper{}
