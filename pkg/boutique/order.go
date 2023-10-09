package boutique

import (
	"github.com/panupakm/boutique-go/pkg/cart"
	"github.com/panupakm/boutique-go/pkg/money"
)

type OrderResult struct {
	OrderId            string      `json:"order_id"`
	ShippingTrackingId string      `json:"shipping_tracking_id"`
	ShippingCost       money.Money `json:"shipping_cost"`
	ShippingAddress    Address     `json:"shipping_address"`
	Items              []OrderItem `json:"items"`
}

type OrderItem struct {
	Item cart.CartItem `json:"item"`
	Cost money.Money   `json:"cost"`
}
