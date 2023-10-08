package generators

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GenerateProducts(ctx context.Context, database *mongo.Database, collection string, fmtstr string, num uint, drop bool) []string {

	coll := database.Collection(collection)
	if drop {
		coll.Drop(ctx)
		model := mongo.IndexModel{
			Keys: bson.D{{Key: "name", Value: "text"}},
		}
		_, err := coll.Indexes().CreateOne(context.TODO(), model)
		if err != nil {
			panic(err)
		}

	}
	docs := []interface{}{}
	uuids := make([]string, 0, num)

	for i := uint(0); i < num; i++ {
		book := gofakeit.Book()
		newUuid := uuid.New().String()
		uuids = append(uuids, newUuid)
		docs = append(docs, bson.M{
			"id": newUuid,
			"name": func() string {
				if fmtstr == "" {
					return book.Title
				}
				return fmt.Sprintf(fmtstr, book.Title)
			}(),
			"description": book.Genre,
			"price_usd": bson.M{
				"currency_code": "THB",
				"units":         1000,
				"nanos":         100000000,
			},
			"categories": []string{"book"},
		})
	}

	opts := options.InsertMany().SetOrdered(false)
	_, err := coll.InsertMany(ctx, docs, opts)
	if err != nil {
		panic(err)
	}

	return uuids
}
