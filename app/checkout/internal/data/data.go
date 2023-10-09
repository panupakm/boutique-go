package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	cartApi "github.com/panupakm/boutique-go/api/cart"
	catalogApi "github.com/panupakm/boutique-go/api/catalog"
	"github.com/panupakm/boutique-go/app/checkout/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCartServiceClient, NewCatalogServiceClient, NewCartRepo, NewCatalogRepo)

// Data .
type Data struct {
	cc  cartApi.CartServiceClient
	clc catalogApi.CatalogClient
}

// NewCartServiceClient create a client to connect to the cart service.
func NewCartServiceClient(c *conf.Data) cartApi.CartServiceClient {
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
	return cartApi.NewCartServiceClient(conn)
}

// NewCatalogClient create a client to connect to the catalog service.
func NewCatalogServiceClient(c *conf.Data) catalogApi.CatalogClient {
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
	return catalogApi.NewCatalogClient(conn)
}

// NewData .
func NewData(cc cartApi.CartServiceClient, clc catalogApi.CatalogClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		cc:  cc,
		clc: clc,
	}, cleanup, nil
}
