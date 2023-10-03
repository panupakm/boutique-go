package tests

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/stretchr/testify/require"

	api "github.com/panupakm/boutique-go/api/email"
)

func TestHttpSendEmail(t *testing.T) {
	type args struct {
		msg string
	}

	client, err := http.NewClient(context.Background(), http.WithEndpoint("localhost:8001"))
	require.NoError(t, err)

	email := api.NewEmailHTTPClient(client)
	request := api.SendOrderConfirmationRequest{
		Email: "anonymous@example.com",
		Order: &api.OrderResult{
			OrderId:            "order_id",
			ShippingTrackingId: "shipping_tracking_id",
			ShippingCost: &api.Money{
				CurrencyCode: "USD",
				Units:        100,
			},
			ShippingAddress: &api.Address{
				StreetAddress: "111",
				City:          "New York",
				State:         "NY",
				Country:       "United State",
				ZipCode:       10001,
			},
			Items: []*api.OrderItem{
				{
					Item: &api.CartItem{
						ProductId: "product_id_1",
						Quantity:  1,
					},
					Cost: &api.Money{
						CurrencyCode: "USD",
						Units:        60,
					},
				},
				{
					Item: &api.CartItem{
						ProductId: "product_id_2",
						Quantity:  1,
					},
					Cost: &api.Money{
						CurrencyCode: "USD",
						Units:        40,
					},
				},
			},
		},
	}
	res, err := email.SendOrderConfirmation(context.Background(), &request)
	require.NoError(t, err)
	require.NotNil(t, res)
}
