package data

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/panupakm/boutique-go/app/catalog/internal/biz"
	"github.com/panupakm/boutique-go/pkg/product"
)

type productRepo struct {
	data        *Data
	productColl *mongo.Collection
	log         *log.Helper
}

// NewProductRepo
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		data:        data,
		productColl: data.db.Collection("products"),
		log:         log.NewHelper(log.With(logger, "module", "catalog/data")),
	}
}

func (r *productRepo) Query(ctx context.Context, q string, pageSize int, pageToken string) (prods []product.Product, err error) {
	var filter bson.M = bson.M{}
	if q != "" {
		filter = bson.M{"$text": bson.M{"$search": q}}
	}

	if pageToken != "" {
		filter["id"] = bson.M{"$gt": pageToken}
	}

	opts := options.Find().SetLimit(int64(pageSize)).SetSort(bson.M{"id": 1})
	cursor, err := r.productColl.Find(ctx, filter, opts)
	if err != nil {
		return
	}

	var bsonMap []bson.M
	if err = cursor.All(ctx, &bsonMap); err != nil {
		return
	}

	for _, bsonData := range bsonMap {
		jsonData, tmpErr := json.Marshal(bsonData)
		if tmpErr != nil {
			err = tmpErr
			return
		}
		var prod product.Product
		json.Unmarshal(jsonData, &prod)
		prods = append(prods, prod)
	}

	return
}

func (r *productRepo) GetProduct(ctx context.Context, id string) (product.Product, error) {
	prod := product.Product{}
	var prodJson bson.M

	r.productColl.FindOne(ctx, bson.M{"id": id}).Decode(&prodJson)
	if err := r.productColl.FindOne(ctx, bson.M{"id": id}).Decode(&prod); err != nil {
		if err == mongo.ErrNoDocuments {
			return product.Product{}, nil
		}
		return product.Product{}, err
	}
	return prod, nil
}
