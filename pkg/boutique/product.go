package boutique

// Product is a Product model.
type Product struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Picture     string `json:"picture,omitempty"`
	PriceUsd    Money
	Categories  []string `json:"categories,omitempty"`
}
