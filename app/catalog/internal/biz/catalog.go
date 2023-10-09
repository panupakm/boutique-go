package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/panupakm/boutique-go/pkg/product"
)

// ProductRepo is a Cart repo.
type ProductRepo interface {
	Query(context.Context, string, int, string) ([]product.Product, error)
	GetProduct(context.Context, string) (product.Product, error)
}

// CartUsecase is a Cart usecase.
type CatalogUsecase struct {
	repo ProductRepo
	log  *log.Helper
}

// NewCartUsecase new a Cart usecase.
func NewCatalogUsecase(repo ProductRepo, logger log.Logger) *CatalogUsecase {
	return &CatalogUsecase{repo: repo, log: log.NewHelper(logger)}
}

// AddItem add cart item into a cart for a user.
func (uc *CatalogUsecase) ListProducts(ctx context.Context, pageSize int, pageToken string) ([]product.Product, error) {
	list, err := uc.repo.Query(ctx, "", pageSize, pageToken)
	if err != nil {
		return []product.Product{}, err
	}

	return list, err
}

// GetCart get cart of a user.
func (uc *CatalogUsecase) GetProduct(ctx context.Context, productId string) (product.Product, error) {
	p, err := uc.repo.GetProduct(ctx, productId)
	if err != nil {
		return product.Product{}, err
	}
	return p, nil
}

// EmptyCart clear all items in a cart.
func (uc *CatalogUsecase) SearchProducts(ctx context.Context, query string, pageSize int, pageToken string) ([]product.Product, error) {
	list, err := uc.repo.Query(ctx, query, pageSize, pageToken)
	if err != nil {
		return []product.Product{}, err
	}

	return list, err
}
