package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"

	"github.com/panupakm/boutique-go/pkg/boutique"
	"github.com/panupakm/boutique-go/pkg/cart"
	"github.com/panupakm/boutique-go/pkg/money"
)

type CheckoutUseCase struct {
	log  *log.Helper
	cluc *CatalogUseCase
	cuc  *CartUseCase
}

func NewCheckoutUseCase(cuc *CartUseCase, cluc *CatalogUseCase, logger log.Logger) *CheckoutUseCase {
	return &CheckoutUseCase{
		cuc:  cuc,
		cluc: cluc,
		log:  log.NewHelper(log.With(logger, "module", "checkout/usercase")),
	}
}

func (cu *CheckoutUseCase) PlaceOrder(ctx context.Context, userId string) (res boutique.OrderResult, err error) {
	orderId, err := uuid.NewUUID()
	if err != nil {
		return
	}
	c, err := cu.cuc.GetCart(ctx, userId)
	if err != nil {
		return
	}
	if len(c.Items) == 0 {
		err = fmt.Errorf("No cart for user %s", userId)
		return
	}

	orderItems := make([]boutique.OrderItem, len(c.Items))
	for i, item := range c.Items {
		prod, err1 := cu.cluc.GetProduct(ctx, item.ProductId)
		if err1 != nil {
			err = err1
			return
		}
		orderItems[i] = boutique.OrderItem{
			Item: cart.CartItem{
				ProductId: prod.Id,
				Quantity:  item.Quantity,
			},
			Cost: prod.PriceUsd,
		}
	}

	total := money.Money{
		CurrencyCode: "USD",
		Units:        0,
		Nanos:        0,
	}

	for _, it := range orderItems {
		multiPrice := money.MultiplySlow(it.Cost, uint32(it.Item.Quantity))
		total = money.Must(money.Sum(total, multiPrice))
	}

	//TODO: calculating shipping cost
	//TOOD: charge credit card

	_ = cu.cuc.EmptyCart(ctx, userId)

	res = boutique.OrderResult{
		OrderId:            orderId.String(),
		ShippingTrackingId: uuid.NewString(),
		ShippingCost:       money.Money{},
		ShippingAddress:    boutique.Address{},
		Items:              orderItems,
	}

	return
}
