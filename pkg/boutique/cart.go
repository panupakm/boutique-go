package boutique

type CartItem struct {
	ProductId string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}
