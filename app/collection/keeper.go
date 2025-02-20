package collection

import (
	"gorm.io/gorm"

	types "phoenix-api/types/collection"
)

type Keeper struct {
	dbHandler *gorm.DB
	types.UnimplementedCollectionQueryServer
}

func NewKeeper(db *gorm.DB) *Keeper {
	return &Keeper{
		dbHandler: db,
	}
}

var _ types.CollectionQueryServer = Keeper{}

