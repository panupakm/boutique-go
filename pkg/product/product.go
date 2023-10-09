package product

import (
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/pkg/money"
)

// Product is a Product model.
type Product struct {
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Picture     string      `json:"picture,omitempty"`
	PriceUsd    money.Money `json:"price,omitempty"`
	Categories  []string    `json:"categories,omitempty"`
}

func ToProto(in *Product, out *spb.Product) {
	out.Id = in.Id
	out.Name = in.Name
	out.Description = in.Description
	out.Picture = in.Picture
	out.PriceUsd = &spb.Money{}
	// ToMoneyProto(&in.PriceUsd, out.PriceUsd)
	out.Categories = in.Categories
}

func ToBiz(in *spb.Product, out *Product) {
	out.Id = in.Id
	out.Name = in.Name
	out.Description = in.Description
	out.Picture = in.Picture
	money.ToBiz(in.PriceUsd, &out.PriceUsd)
	out.Categories = in.Categories
}
