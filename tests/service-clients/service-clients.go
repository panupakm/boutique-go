package serviceclients

import (
	"context"

	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"

	"github.com/panupakm/boutique-go/tests/config"
)

var GrpcClientMaps = make(map[string]*grpc.ClientConn, 4)

func InitGrpcClients(ctx context.Context, conf config.TestConfig) {
	var err error
	for k, v := range conf.Clients {
		GrpcClientMaps[k], err = kgrpc.DialInsecure(ctx, kgrpc.WithEndpoint(v.GrpcUri))
		if err != nil {
			panic(err)
		}
	}
}

func CloseGrpcClients(ctx context.Context) {
	for k, v := range GrpcClientMaps {
		if err := v.Close(); err != nil {
			panic(err)
		}
		delete(GrpcClientMaps, k)
	}
}
