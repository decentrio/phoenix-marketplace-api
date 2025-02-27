package activity

import (
	types "phoenix-marketplace-api/types/activity"

	"google.golang.org/grpc"
)

type AppModule struct {
	keeper Keeper
}

func NewAppModule(
	keeper Keeper,
) AppModule {
	return AppModule{
		keeper: keeper,
	}
}

func (am AppModule) RegisterServices(server *grpc.Server) {
	types.RegisterActivityQueryServer(server, am.keeper)
}
