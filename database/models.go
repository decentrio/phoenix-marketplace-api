package database

import _ "gorm.io/gorm"

// deployer: deploy_new_collection, collection: set_collection_uri, mint, mint_batch
type Collection struct {
	Address    string `gorm:"primaryKey" json:"address,omitempty"`
	Name       string `json:"name,omitempty"`
	Symbol     string `json:"symbol,omitempty"`
	URI        string `json:"uri,omitempty"`
	TotalItems uint64 `json:"total_items,omitempty"`
}

// collection: set_uri, mint, mint_batch
type NFT struct {
	Nid        string `gorm:"primaryKey" json:"nid,omitempty"`
	ID         string `json:"id,omitempty"`
	Collection string `json:"collection,omitempty"`
	URI        string `json:"uri,omitempty"`
}

// collection: mint, mint_batch, burn, burn_batch, safe_transfer_from, safe_batch_transfer_from, buy_now, finalize_auction
type Balance struct {
	Hash    string `gorm:"primaryKey" json:"hash,omitempty"`
	Account string `json:"account,omitempty"`
	Nid     string `json:"nid,omitempty"`
	Amount  uint64 `json:"amount,omitempty"`
	TxTime  uint64 `json:"tx_time,omitempty"`
	Ledger  int    `json:"ledger,omitempty"`
}

// auction: finalize_auction, buy_now, place_bid, create_auction
type Price struct {
	Hash   string `gorm:"primaryKey" json:"hash,omitempty"`
	Nid    string `json:"nid,omitempty"`
	Price  uint64 `json:"price,omitempty"`
	Token  string `json:"token,omitempty"`
	TxTime uint64 `json:"tx_time,omitempty"`
}

// auction: create_auction, place_bid, finalize_auction, buy_now, pause, unpause
type Auction struct {
	ID           uint64 `gorm:"primaryKey" json:"id,omitempty"`
	Seller       string `json:"seller,omitempty"`
	HighestBid   uint64 `json:"highest_bid,omitempty"`
	EndTime      uint64 `json:"end_time,omitempty"`
	Status       string `json:"status,omitempty"`
	AuctionToken string `json:"auction_token,omitempty"`
	Nid          string `json:"nid,omitempty"`
	Amount       uint64 `json:"amount,omitempty"`
	MinimumPrice uint64 `json:"minimum_price,omitempty"`
	BuyNowPrice  uint64 `json:"buy_now_price,omitempty"`
}

// auction: place_bid
type Bid struct {
	Hash      string `gorm:"primaryKey" json:"hash,omitempty"`
	AuctionID uint64 `json:"auction_id,omitempty"`
	Bidder    string `json:"bidder,omitempty"`
	BidAmount uint64 `json:"bid_amount,omitempty"`
}

// collection: mint, mint_batch, burn, burn_batch, safe_transfer_from, safe_batch_transfer_from, auction: create_auction, finalize_auction, buy_now
type Activity struct {
	Nid     string `json:"nid,omitempty"`
	Type    string `json:"type,omitempty"`
	Details byte `gorm:"type:jsonb" json:"details,omitempty"`
}

// auction: finalize_auction, buy_now
type Volume struct {
	Hash       string `gorm:"primaryKey" json:"hash,omitempty"`
	Nid        string `json:"nid,omitempty"`
	Collection string `json:"collection,omitempty"`
	Amount     uint64 `json:"amount,omitempty"`
	Volume     uint64 `json:"volume,omitempty"`
	Token      string `json:"token,omitempty"`
}
