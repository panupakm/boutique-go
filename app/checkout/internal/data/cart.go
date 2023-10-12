package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	cartApi "github.com/panupakm/boutique-go/api/cart"
	"github.com/panupakm/boutique-go/app/checkout/internal/biz"
	"github.com/panupakm/boutique-go/pkg/cart"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "cart/data")),
	}
}

func (r *cartRepo) GetCart(ctx context.Context, userId string) (cart.Cart, error) {
	pbc, err := r.data.cartc.GetCart(ctx, &cartApi.GetCartRequest{
		UserId: userId,
	})

	if err != nil {
		return cart.Cart{}, err
	}

	var bizCart = cart.Cart{}
	cart.ToCartBiz(pbc, &bizCart)
	return bizCart, nil
}

func (r *cartRepo) EmptyCart(ctx context.Context, userId string) error {
	_, err := r.data.cartc.EmptyCart(ctx, &cartApi.EmptyCartRequest{})

	if err != nil {
		return err
	}
	return nil
}
