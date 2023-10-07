package data

import (
	"context"
	"time"

	"github.com/panupakm/boutique-go/app/catalog/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo, NewMongo)

// Data .
type Data struct {
	db  *mongo.Database
	log *log.Helper
}

// NewMongo
func NewMongo(conf *conf.Data) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.MongoDb.Uri+"/"+conf.MongoDb.Database))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client.Database(conf.MongoDb.Database)
}

// NewData .
func NewData(database *mongo.Database, c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "catalog/data"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	d := &Data{
		db:  database,
		log: log,
	}

	return d, func() {
		if err := d.db.Client().Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}, nil
}
