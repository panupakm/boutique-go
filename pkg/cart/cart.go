package cart

import (
	shared "github.com/panupakm/boutique-go/api/shared"
)

type CartItem struct {
	ProductId string `bson:"product_id" json:"product_id"`
	Quantity  int32  `bson:"quantity" json:"quantity"`
}

type Cart struct {
	UserId string     `bson:"user_id" json:"user_id"`
	Items  []CartItem `bson:"items,omitempty" json:"items,omitempty"`
}

func ToCartBiz(in *shared.Cart, out *Cart) {
	out.UserId = in.UserId
	out.Items = make([]CartItem, len(in.Items))
	for i, item := range in.Items {
		out.Items[i] = CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
}

func ToCartProto(in *Cart, out *shared.Cart) {
	out.UserId = in.UserId
	out.Items = make([]*shared.CartItem, len(in.Items))
	for i, item := range in.Items {
		ToCartItemProto(&item, out.Items[i])
	}
}

func ToCartItemProto(in *CartItem, out *shared.CartItem) {
	out.ProductId = in.ProductId
	out.Quantity = in.Quantity
}
