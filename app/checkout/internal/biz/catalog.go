package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/panupakm/boutique-go/pkg/product"
)

type CatalogRepo interface {
	GetProduct(ctx context.Context, productId string) (product.Product, error)
}

type CatalogUseCase struct {
	repo CatalogRepo
	log  *log.Helper
}

func NewCatalogUseCase(repo CatalogRepo, logger log.Logger) *CatalogUseCase {
	return &CatalogUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "catalog/usercase"))}
}

func (cu *CatalogUseCase) GetProduct(ctx context.Context, id string) (product.Product, error) {
	return cu.repo.GetProduct(ctx, id)
}
