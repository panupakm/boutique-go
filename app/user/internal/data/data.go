package data

import (
	"context"

	"github.com/panupakm/boutique-go/app/user/internal/conf"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent/migrate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewUserRepo, NewCardRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: entClient,
	}, cleanup, nil
}

func NewEntClient(c *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "user/data/ent"))
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
