package user

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	api "github.com/panupakm/boutique-go/api/user"
	"github.com/panupakm/boutique-go/pkg/boutique/generators"
	"github.com/panupakm/boutique-go/pkg/util"
	"github.com/panupakm/boutique-go/tests"
	serviceclients "github.com/panupakm/boutique-go/tests/service-clients"
)

var client api.UserClient

func TestMain(m *testing.M) {
	tests.SetUp()

	client = api.NewUserClient(serviceclients.GrpcClientMaps["user"])

	code := m.Run()
	tests.TearDown()
	os.Exit(code)
}

func TestGrpcCreateUser(t *testing.T) {
	ctx := context.Background()
	username := generators.Username()
	email := generators.Email()
	res, err := client.CreateUser(ctx, &api.CreateUserReq{
		Username: username,
		Password: "12345678",
		Email:    email,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotEmpty(t, res.Id, res.Username)
	require.Equal(t, username, res.Username)
}

func TestGrpcAddCards(t *testing.T) {
	ctx := context.Background()
	username := generators.Username()
	email := generators.Email()

	resuser, _ := client.CreateUser(ctx, &api.CreateUserReq{
		Username: username,
		Password: "12345678",
		Email:    email,
	})

	rescard, err := client.AddCard(ctx, &api.AddCardReq{
		UserId:          resuser.Id,
		Name:            generators.Name(),
		CardNumber:      util.GetRandomNumStr(16),
		ExpirationYear:  int32(util.GetNumberRange(2021, 2030)),
		ExpirationMonth: int32(util.GetNumberRange(1, 12)),
		Ccv:             int32(util.GetNumberRange(100, 999)),
	})

	require.NoError(t, err)
	require.NotEmpty(t, rescard.CardId)
}

func TestGrpcListCards(t *testing.T) {
	ctx := context.Background()
	username := generators.Username()
	email := generators.Email()

	resuser, _ := client.CreateUser(ctx, &api.CreateUserReq{
		Username: username,
		Password: "12345678",
		Email:    email,
	})

	testReqs := make([]*api.AddCardReq, 2)
	for i, _ := range testReqs {
		req := api.AddCardReq{
			UserId:          resuser.Id,
			Name:            generators.Name(),
			CardNumber:      util.GetRandomNumStr(16),
			ExpirationYear:  int32(util.GetNumberRange(2021, 2030)),
			ExpirationMonth: int32(util.GetNumberRange(1, 12)),
			Ccv:             int32(util.GetNumberRange(100, 999)),
		}
		testReqs[i] = &req
		client.AddCard(ctx, &req)
	}

	res, err := client.ListCards(ctx, &api.ListCardsReq{
		UserId: resuser.Id,
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 2, len(res.Cards))
	require.Equal(t, testReqs[0].CardNumber, res.Cards[0].CreditCardNumber)
	require.Equal(t, testReqs[1].CardNumber, res.Cards[1].CreditCardNumber)
}
