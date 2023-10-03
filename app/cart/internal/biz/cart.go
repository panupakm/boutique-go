package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Cart is a Cart model.
type Cart struct {
	UserId string     `json:"user_id"`
	Items  []CartItem `json:"items,omitempty"`
}

type CartItem struct {
	ProductId string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

// CartRepo is a Cart repo.
type CartRepo interface {
	AddItem(context.Context, string, *CartItem) error
	Empty(context.Context, string) error
	GetCart(context.Context, string) (*Cart, error)
}

// CartUsecase is a Cart usecase.
type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

// NewCartUsecase new a Cart usecase.
func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

// AddItem add cart item into a cart for a user.
func (uc *CartUsecase) AddItem(ctx context.Context, userId string, item *CartItem) error {
	err := uc.repo.AddItem(ctx, userId, item)
	if err != nil {
		return err
	}
	uc.log.WithContext(ctx).Infof("AddItem: %v", *item)

	return nil
}

// GetCart get cart of a user.
func (uc *CartUsecase) GetCart(ctx context.Context, userId string) (*Cart, error) {
	cart, err := uc.repo.GetCart(ctx, userId)
	if err != nil {
		return nil, err
	}
	uc.log.WithContext(ctx).Infof("GetCart: %s", userId)
	return cart, nil
}
