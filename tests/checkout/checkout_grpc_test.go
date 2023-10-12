package checkout

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	cartapi "github.com/panupakm/boutique-go/api/cart"
	catalogapi "github.com/panupakm/boutique-go/api/catalog"
	checkoutapi "github.com/panupakm/boutique-go/api/checkout"
	shared "github.com/panupakm/boutique-go/api/shared"
	userapi "github.com/panupakm/boutique-go/api/user"
	"github.com/panupakm/boutique-go/pkg/boutique/generators"
	"github.com/panupakm/boutique-go/pkg/util"
	tests "github.com/panupakm/boutique-go/tests"
	serviceclients "github.com/panupakm/boutique-go/tests/service-clients"
	"github.com/stretchr/testify/require"
)

var cartClient cartapi.CartServiceClient
var catalogClient catalogapi.CatalogClient
var checkoutClient checkoutapi.CheckoutClient
var userClient userapi.UserClient
var Client catalogapi.CatalogClient

func TestMain(m *testing.M) {
	tests.SetUp()

	cartClient = cartapi.NewCartServiceClient(serviceclients.GrpcClientMaps["cart"])
	catalogClient = catalogapi.NewCatalogClient(serviceclients.GrpcClientMaps["catalog"])
	checkoutClient = checkoutapi.NewCheckoutClient(serviceclients.GrpcClientMaps["checkout"])
	userClient = userapi.NewUserClient(serviceclients.GrpcClientMaps["user"])

	code := m.Run()
	tests.TearDown()
	os.Exit(code)
}

func TestGrpcCheckout(t *testing.T) {

	ctx := context.Background()
	n := 10

	replyuser, err := userClient.CreateUser(ctx, &userapi.CreateUserReq{
		Username: util.GetRandomStr(10),
		Password: util.GetRandomStr(10),
		Email:    generators.Email(),
	})

	listRes, err := catalogClient.ListProducts(ctx, &catalogapi.ListProductsRequest{
		PageSize: int32(n),
	})
	require.NoError(t, err)
	require.GreaterOrEqualf(t, len(listRes.Products), n, "should have at least %d products", n)

	quantities := [10]int32{1, 1, 2, 2, 1, 1, 1, 1, 1, 1}
	prices := [10]*shared.Money{}
	for i, p := range listRes.Products {
		prices[i] = p.PriceUsd
		req := cartapi.AddItemRequest{
			UserId: replyuser.Id,
			Item: &shared.CartItem{
				ProductId: p.Id,
				Quantity:  quantities[i],
			},
		}
		_, err := cartClient.AddItem(ctx, &req)
		require.NoError(t, err)
	}

	pores, err := checkoutClient.PlaceOrder(ctx, &checkoutapi.PlaceOrderRequest{
		UserId:       replyuser.Id,
		Email:        "testcheckout@boutique.com",
		UserCurrency: "THB",
		Address: &shared.Address{
			StreetAddress: "123 Shipping St",
			City:          "Seattle",
			State:         "WA",
			Country:       "USA",
			ZipCode:       98101,
		},
		CreditCard: &shared.CreditCardInfo{
			CreditCardNumber:          "1234567890123456",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2043,
			CreditCardExpirationMonth: 12,
		},
	})

	require.NoError(t, err)
	require.NotNil(t, pores)
	require.NotNil(t, pores.Order.OrderId)
	require.NotNil(t, pores.Order.ShippingAddress)
	require.NotNil(t, pores.Order.ShippingCost)
	require.NotNil(t, pores.Order.ShippingTrackingId)
	require.Equal(t, n, len(pores.Order.Items))
	for _, item := range pores.Order.Items {
		require.NotEmpty(t, item.Item.ProductId)
		require.NotZero(t, item.Item.Quantity)
	}
}

func TestGrpcCheckoutFailed(t *testing.T) {

	ctx := context.Background()

	userId := uuid.New().String()
	_, err := checkoutClient.PlaceOrder(ctx, &checkoutapi.PlaceOrderRequest{
		UserId:       userId,
		Email:        "testcheckout@boutique.com",
		UserCurrency: "THB",
		Address: &shared.Address{
			StreetAddress: "123 Shipping St",
			City:          "Seattle",
			State:         "WA",
			Country:       "USA",
			ZipCode:       98101,
		},
		CreditCard: &shared.CreditCardInfo{
			CreditCardNumber:          "1234567890123456",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2043,
			CreditCardExpirationMonth: 12,
		},
	})

	require.Error(t, err)
}
