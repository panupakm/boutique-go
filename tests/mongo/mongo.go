package mongo

import (
	"context"

	"github.com/panupakm/boutique-go/tests/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
var MongoDBMaps = map[string]*mongo.Database{}

func InitMongoDbClient(ctx context.Context, conf config.TestConfig) {
	var err error
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(conf.MongoDb.Uri))
	if err != nil {
		panic(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	for _, name := range conf.MongoDb.Databases {
		MongoDBMaps[name] = mongoClient.Database(name)
	}
}

func CloseMongoDBClient(ctx context.Context) {
	mongoClient.Disconnect(ctx)
}
