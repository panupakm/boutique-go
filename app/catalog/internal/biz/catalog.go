package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/panupakm/boutique-go/pkg/boutique"
)

// ProductRepo is a Cart repo.
type ProductRepo interface {
	Query(context.Context, string, int, string) ([]boutique.Product, error)
	GetProduct(context.Context, string) (boutique.Product, error)
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
func (uc *CatalogUsecase) ListProducts(ctx context.Context, pageSize int, pageToken string) ([]boutique.Product, error) {
	list, err := uc.repo.Query(ctx, "", pageSize, pageToken)
	if err != nil {
		return []boutique.Product{}, err
	}

	return list, err
}

// GetCart get cart of a user.
func (uc *CatalogUsecase) GetProduct(ctx context.Context, productId string) (boutique.Product, error) {
	p, err := uc.repo.GetProduct(ctx, productId)
	if err != nil {
		return boutique.Product{}, err
	}
	return p, nil
}

// EmptyCart clear all items in a cart.
func (uc *CatalogUsecase) SearchProducts(ctx context.Context, query string, pageSize int, pageToken string) ([]boutique.Product, error) {
	list, err := uc.repo.Query(ctx, query, pageSize, pageToken)
	if err != nil {
		return []boutique.Product{}, err
	}

	return list, err
}
