package main

import (
	"context"
	"flag"
	"time"

	"github.com/panupakm/boutique-go/pkg/boutique/generators"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	uri := flag.String("uri", "mongodb://user:password@localhost:27017", "uri to connect to mongodb")
	dbName := flag.String("db", "catalog", "database for collections")
	num := flag.Uint("num", 100, "number of generated products")

	flag.Parse()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	struri := *uri + "/" + *dbName
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(struri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	database := client.Database(*dbName)
	generators.GenerateProducts(ctx, database, "products", "", *num, true)
}
