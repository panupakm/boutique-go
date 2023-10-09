package cart

import (
	cartApi "github.com/panupakm/boutique-go/api/cart"
)

type CartItem struct {
	ProductId string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type Cart struct {
	UserId string     `json:"user_id"`
	Items  []CartItem `json:"items,omitempty"`
}

func ToBiz(pb *cartApi.Cart, biz *Cart) {
	biz.UserId = pb.UserId
	biz.Items = make([]CartItem, len(pb.Items))
	for i, item := range pb.Items {
		biz.Items[i] = CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
}

func ToProto(biz *Cart, pb *cartApi.Cart) {
	pb.UserId = biz.UserId
	pb.Items = make([]*cartApi.CartItem, len(biz.Items))
	for i, item := range biz.Items {
		pb.Items[i] = &cartApi.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
	}
}
