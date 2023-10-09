package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/panupakm/boutique-go/pkg/boutique"
)

type CartRepo interface {
	GetCart(ctx context.Context, userId string) (boutique.Cart, error)
	EmptyCart(ctx context.Context, userId string) error
}

type CartUseCase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUseCase(repo CartRepo, logger log.Logger) *CartUseCase {
	return &CartUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "cart/usercase"))}
}

func (cu *CartUseCase) GetCart(ctx context.Context, id string) (boutique.Cart, error) {
	return cu.repo.GetCart(ctx, id)
}

func (cu *CartUseCase) EmptyCart(ctx context.Context, id string) error {
	return cu.repo.EmptyCart(ctx, id)
}
