package tests

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	api "github.com/panupakm/boutique-go/api/catalog"
	sharedApi "github.com/panupakm/boutique-go/api/shared"
)

func TestGrpcListProducts(t *testing.T) {
	type args struct {
		msg string
	}

	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9003"))
	assert.NoError(t, err)

	client := api.NewCatalogClient(conn)
	res, err := client.ListProducts(context.Background(), &sharedApi.Empty{})
	require.NoError(t, err)
	require.NotNil(t, res)
}
