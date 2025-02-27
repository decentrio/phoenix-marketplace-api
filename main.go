package main

import (
	"context"
	"log"
	"net"
	"net/http"

	_ "phoenix-marketplace-api/docs/statik"
	"github.com/rakyll/statik/fs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"phoenix-marketplace-api/app"
	"phoenix-marketplace-api/database"

	"phoenix-marketplace-api/app/activity"
	"phoenix-marketplace-api/app/auction"
	"phoenix-marketplace-api/app/collection"
	"phoenix-marketplace-api/app/nft"
	"phoenix-marketplace-api/app/offer"
	activitytypes "phoenix-marketplace-api/types/activity"
	auctiontypes "phoenix-marketplace-api/types/auction"
	collectiontypes "phoenix-marketplace-api/types/collection"
	nfttypes "phoenix-marketplace-api/types/nft"
	offertypes "phoenix-marketplace-api/types/offer"
)

func initModule() []app.AppModule {
	dbHandler := database.NewDBHandler()

	activityKeeper := activity.NewKeeper(dbHandler)
	auctionKeeper := auction.NewKeeper(dbHandler)
	collectionKeeper := collection.NewKeeper(dbHandler)
	nftKeeper := nft.NewKeeper(dbHandler)
	offerKeeper := offer.NewKeeper(dbHandler)

	modules := []app.AppModule{
		activity.NewAppModule(*activityKeeper),
		auction.NewAppModule(*auctionKeeper),
		collection.NewAppModule(*collectionKeeper),
		nft.NewAppModule(*nftKeeper),
		offer.NewAppModule(*offerKeeper),
	}

	return modules
}

func runGRPCServer() error {
	lis, err := net.Listen("tcp", ":6070")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	modules := initModule()
	bookApp := app.NewApp(s, modules)
	bookApp.RegisterServices()
	return s.Serve(lis)
}

func runHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := activitytypes.RegisterActivityQueryHandlerFromEndpoint(ctx, mux, ":6070", opts)
	if err != nil {
		return err
	}

	err = auctiontypes.RegisterAuctionQueryHandlerFromEndpoint(ctx, mux, ":6070", opts)
	if err != nil {
		return err
	}

	err = collectiontypes.RegisterCollectionQueryHandlerFromEndpoint(ctx, mux, ":6070", opts)
	if err != nil {
		return err
	}

	err = nfttypes.RegisterNftQueryHandlerFromEndpoint(ctx, mux, ":6070", opts)
	if err != nil {
		return err
	}

	err = offertypes.RegisterOfferQueryHandlerFromEndpoint(ctx, mux, ":6070", opts)
	if err != nil {
		return err
	}

	http.Handle("/", mux)
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)

	// Serve Swagger UI

	http.Handle("/public/", http.StripPrefix("/public/", staticServer))

	log.Println("HTTP server listening on :6060")
	return http.ListenAndServe(":6060", nil)
}

func main() {
	go func() {
		if err := runGRPCServer(); err != nil {
			log.Fatalf("failed to run gRPC server: %v", err)
		}
	}()
	if err := runHTTPServer(); err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}
