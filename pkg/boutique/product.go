package boutique

import "github.com/panupakm/boutique-go/pkg/money"

// Product is a Product model.
type Product struct {
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Picture     string      `json:"picture,omitempty"`
	PriceUsd    money.Money `json:"price,omitempty"`
	Categories  []string    `json:"categories,omitempty"`
}
