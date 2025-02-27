package app


import (
	"encoding/json"
	"io"
	"net/http"

	"google.golang.org/protobuf/types/known/structpb"
)

const (
	PAGE_SIZE        = 10
	ACCOUNT_TABLE    = "accounts"
	COLLECTION_TABLE = "collections"
	NFT_TABLE        = "nfts"
	AUCTION_TABLE    = "auctions"
	PRICE_TABLE      = "prices"
	ACTIVITY_TABLE   = "activities"
	OFFER_TABLE      = "offers"
	BALANCE_TABLE    = "balances"
	VOLUME_TABLE     = "volumes"
	BID_TABLE        = "bids"
	DIVIDE_DECIMALS  = 10000000
)


func FetchAndParseJSON(uri string) *structpb.Struct {
	resp, err := http.Get(uri)
	if err != nil {
		return createErrorStruct(uri)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return createErrorStruct(uri)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return createErrorStruct(uri)
	}

	metadata, err := structpb.NewStruct(data)
	if err != nil {
		return createErrorStruct(uri)
	}

	return metadata
}

func createErrorStruct(uri string) *structpb.Struct {
	errorStruct, _ := structpb.NewStruct(map[string]interface{}{
		"uri": uri,
	})
	return errorStruct
}
