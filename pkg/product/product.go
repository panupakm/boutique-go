package product

import (
	spb "github.com/panupakm/boutique-go/api/shared"
	"github.com/panupakm/boutique-go/pkg/money"
)

// Product is a Product model.
type Product struct {
	Id          string      `bson:"id,omitempty" json:"id,omitempty"`
	Name        string      `bson:"name,omitempty" json:"name,omitempty"`
	Description string      `bson:"description,omitempty" json:"description,omitempty"`
	Picture     string      `bson:"picture,omitempty" json:"picture,omitempty"`
	PriceUsd    money.Money `bson:"price_usd,omitempty" json:"price_usd,omitempty"`
	Categories  []string    `bson:"categories,omitempty" json:"categories,omitempty"`
}

func ToProto(in *Product, out *spb.Product) {
	out.Id = in.Id
	out.Name = in.Name
	out.Description = in.Description
	out.Picture = in.Picture
	out.PriceUsd = &spb.Money{}
	money.ToProto(&in.PriceUsd, out.PriceUsd)
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
