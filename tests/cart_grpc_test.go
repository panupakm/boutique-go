package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	api "github.com/panupakm/boutique-go/api/cart"
	"github.com/panupakm/boutique-go/api/shared"
)

func TestGrpcAddItem(t *testing.T) {
	type args struct {
		msg string
	}

	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9002"))
	assert.NoError(t, err)

	client := api.NewCartServiceClient(conn)
	userId := fmt.Sprintf("user_id_%d", time.Now().UnixMilli())
	request := api.AddItemRequest{
		UserId: userId,
		Item: &shared.CartItem{
			ProductId: "product_id_1",
			Quantity:  1,
		},
	}
	res, err := client.AddItem(context.Background(), &request)
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = client.AddItem(context.Background(), &request)
	require.NoError(t, err)
	require.NotNil(t, res)

	cart, err := client.GetCart(context.Background(), &api.GetCartRequest{
		UserId: userId,
	})
	require.NoError(t, err)
	require.Equal(t, 1, len(cart.Items))
	require.Equal(t, "product_id_1", cart.Items[0].ProductId)
	require.Equal(t, int32(2), cart.Items[0].Quantity)
}

func TestGrpcEmptyCart(t *testing.T) {
	type args struct {
		msg string
	}

	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9002"))
	assert.NoError(t, err)

	client := api.NewCartServiceClient(conn)
	userId := fmt.Sprintf("user_id_%d", time.Now().UnixMilli())
	request := api.AddItemRequest{
		UserId: userId,
		Item: &shared.CartItem{
			ProductId: "product_id_3",
			Quantity:  1,
		},
	}
	res, err := client.AddItem(context.Background(), &request)
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = client.EmptyCart(context.Background(), &api.EmptyCartRequest{
		UserId: userId,
	})
	require.NoError(t, err)

	cart, err := client.GetCart(context.Background(), &api.GetCartRequest{
		UserId: userId,
	})
	require.NoError(t, err)
	require.Equal(t, 0, len(cart.Items))
}
