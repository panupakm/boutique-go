package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	uri := flag.String("uri", "mongodb://user:password@localhost:27017", "uri to connect to mongodb")
	dbName := flag.String("db", "catalog", "database for collections")

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
	coll := database.Collection("products")

	// res := coll.FindOne(ctx, bson.M{"name": "Oedipus the King"})
	// var p boutique.Product
	// res.Decode(&p)
	// fmt.Println(p)

	opts := options.Find()
	cursor, _ := coll.Find(ctx, opts)
	var bsonMap []bson.M
	if err = cursor.All(ctx, &bsonMap); err != nil {
		return
	}
	fmt.Println("query num = ", len(bsonMap))
}
