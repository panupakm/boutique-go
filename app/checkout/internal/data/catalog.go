package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/panupakm/boutique-go/app/checkout/internal/biz"
	"github.com/panupakm/boutique-go/pkg/product"
)

type catalogRepo struct {
	data *Data
	log  *log.Helper
}

// NewCatalogRepo .
func NewCatalogRepo(data *Data, logger log.Logger) biz.CatalogRepo {
	return &catalogRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "catalog/data")),
	}
}

func (r *catalogRepo) GetProduct(ctx context.Context, productId string) (product.Product, error) {
	// product, err := r.data.clc.GetProduct(ctx, &catalogApi.GetProductRequest{
	// 	Id: productId,
	// })

	return product.Product{}, nil
}
