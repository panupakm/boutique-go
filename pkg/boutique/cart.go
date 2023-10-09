package boutique

type CartItem struct {
	ProductId string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type Cart struct {
	UserId string     `json:"user_id"`
	Items  []CartItem `json:"items,omitempty"`
}
