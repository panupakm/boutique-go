package data

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/panupakm/boutique-go/app/catalog/internal/biz"
	"github.com/panupakm/boutique-go/pkg/boutique"
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

func (r *productRepo) Query(ctx context.Context, q string) (prods []boutique.Product, err error) {
	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: q}}}}
	cursor, err := r.productColl.Find(ctx, filter)
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
		var prod boutique.Product
		json.Unmarshal(jsonData, &prod)
		prods = append(prods, prod)
	}

	return
}

func (r *productRepo) GetProduct(ctx context.Context, id string) (boutique.Product, error) {
	prod := boutique.Product{}
	if err := r.productColl.FindOne(ctx, bson.M{"id": id}).Decode(&prod); err != nil {
		if err == mongo.ErrNoDocuments {
			return boutique.Product{}, nil
		}
		return boutique.Product{}, err
	}

	return prod, nil
}
