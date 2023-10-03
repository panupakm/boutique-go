package tests

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	api "github.com/panupakm/boutique-go/api/email"
)

func TestGrpcSendEmail(t *testing.T) {
	type args struct {
		msg string
	}

	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9000"))
	assert.NoError(t, err)

	email := api.NewEmailClient(conn)
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
