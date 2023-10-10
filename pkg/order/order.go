package order

import (
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/pkg/address"
	"github.com/panupakm/boutique-go/pkg/cart"
	"github.com/panupakm/boutique-go/pkg/money"
)

type OrderResult struct {
	OrderId            string          `bson:"order_id" json:"order_id"`
	ShippingTrackingId string          `bson:"shipping_tracking_id" json:"shipping_tracking_id"`
	ShippingCost       money.Money     `bson:"shipping_cost" json:"shipping_cost"`
	ShippingAddress    address.Address `bson:"shipping_address" json:"shipping_address"`
	Items              []OrderItem     `bson:"items" json:"items"`
}

type OrderItem struct {
	Item cart.CartItem `bson:"item" json:"item"`
	Cost money.Money   `bson:"cost" json:"cost"`
}

func ToProtoItem(in *OrderItem, out *spb.OrderItem) {
	out.Item = &spb.CartItem{
		ProductId: in.Item.ProductId,
		Quantity:  in.Item.Quantity,
	}
	out.Cost = &spb.Money{
		CurrencyCode: in.Cost.CurrencyCode,
		Units:        in.Cost.Units,
		Nanos:        in.Cost.Nanos,
	}
}

func ToProtoResult(in *OrderResult, out *spb.OrderResult) {
	out.OrderId = in.OrderId
	out.ShippingTrackingId = in.ShippingTrackingId
	out.ShippingCost = &spb.Money{}
	out.ShippingAddress = &spb.Address{}
	money.ToProto(&in.ShippingCost, out.ShippingCost)
	address.ToProto(&in.ShippingAddress, out.ShippingAddress)
	out.Items = make([]*spb.OrderItem, len(in.Items))
	for i, item := range in.Items {
		out.Items[i] = &spb.OrderItem{
			Item: &spb.CartItem{
				ProductId: item.Item.ProductId,
				Quantity:  item.Item.Quantity,
			},
			Cost: &spb.Money{
				CurrencyCode: item.Cost.CurrencyCode,
				Units:        item.Cost.Units,
				Nanos:        item.Cost.Nanos,
			},
		}
	}
}
