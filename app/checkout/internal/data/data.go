package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	cartapi "github.com/panupakm/boutique-go/api/cart"
	catlapi "github.com/panupakm/boutique-go/api/catalog"
	userapi "github.com/panupakm/boutique-go/api/user"
	"github.com/panupakm/boutique-go/app/checkout/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCartServiceClient, NewCatalogServiceClient, NewCartRepo, NewCatalogRepo, NewUserServiceClient, NewUserRepo)

// Data .
type Data struct {
	cartc cartapi.CartServiceClient
	catlc catlapi.CatalogClient
	userc userapi.UserClient
}

// NewCartServiceClient create a client to connect to the cart service.
func NewCartServiceClient(c *conf.Data) cartapi.CartServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.CartService.Uri),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return cartapi.NewCartServiceClient(conn)
}

// NewCatalogClient create a client to connect to the catalog service.
func NewCatalogServiceClient(c *conf.Data) catlapi.CatalogClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.CatalogService.Uri),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return catlapi.NewCatalogClient(conn)
}

func NewUserServiceClient(c *conf.Data) userapi.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.UserService.Uri),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userapi.NewUserClient(conn)
}

// NewData .
func NewData(cc cartapi.CartServiceClient, clc catlapi.CatalogClient, userc userapi.UserClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		cartc: cc,
		catlc: clc,
		userc: userc,
	}, cleanup, nil
}
