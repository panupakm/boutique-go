package catalog

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"

	api "github.com/panupakm/boutique-go/api/catalog"
	"github.com/panupakm/boutique-go/pkg/boutique/generators"
	"github.com/panupakm/boutique-go/pkg/util"
	"github.com/panupakm/boutique-go/tests"
	testmongo "github.com/panupakm/boutique-go/tests/mongo"
	serviceclients "github.com/panupakm/boutique-go/tests/service-clients"
)

var client api.CatalogClient
var mongoDb *mongo.Database

func TestMain(m *testing.M) {
	tests.SetUp()

	client = api.NewCatalogClient(serviceclients.GrpcClientMaps["catalog"])
	mongoDb = testmongo.MongoDBMaps["catalog"]

	code := m.Run()
	tests.TearDown()
	os.Exit(code)
}

func TestGrpcListProducts(t *testing.T) {
	res, err := client.ListProducts(context.Background(), &api.ListProductsRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.LessOrEqual(t, 100, len(res.Products))
}

func TestGrpcListProductsWithPagination(t *testing.T) {
	res1, err := client.ListProducts(context.Background(), &api.ListProductsRequest{
		PageSize: 10,
	})
	require.NoError(t, err)
	require.NotNil(t, res1)
	require.Equal(t, 10, len(res1.Products))

	res2, err := client.ListProducts(context.Background(), &api.ListProductsRequest{
		PageSize:  10,
		PageToken: res1.Products[len(res1.Products)-1].Id,
	})

	require.NoError(t, err)
	require.NotNil(t, res2)
	require.Equal(t, 10, len(res1.Products))
	require.NotEqual(t, res1.Products[0].Id, res2.Products[0].Id)
}

func TestGrpcGetProduct(t *testing.T) {

	uuids := generators.GenerateProducts(context.Background(), mongoDb, "products", "", 1, false)

	res, err := client.GetProduct(context.Background(), &api.GetProductRequest{
		Id: uuids[0],
	})
	require.NoError(t, err)
	require.Equal(t, uuids[0], res.Id)
}

func TestGrpcSearchProduct(t *testing.T) {
	n := uint(10)
	searchText := util.GetRandomStr(10)

	generators.GenerateProducts(context.Background(), mongoDb, "products", searchText+" %s ", n, false)

	res, err := client.SearchProducts(context.Background(), &api.SearchProductsRequest{
		Query: searchText,
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, int(n), len(res.Results))
}

func TestGrpcSearchProductWithPagination(t *testing.T) {
	n := uint(15)
	searchText := util.GetRandomStr(10)

	generators.GenerateProducts(context.Background(), mongoDb, "products", searchText+" %s ", n, false)

	res, err := client.SearchProducts(context.Background(), &api.SearchProductsRequest{
		Query:    searchText,
		PageSize: 10,
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, int(10), len(res.Results))

	res, err = client.SearchProducts(context.Background(), &api.SearchProductsRequest{
		Query:     searchText,
		PageSize:  10,
		PageToken: res.Results[len(res.Results)-1].Id,
	})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, int(5), len(res.Results))
}
