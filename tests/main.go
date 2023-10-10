package tests

import (
	"context"
	"encoding/json"
	"os"

	"github.com/panupakm/boutique-go/tests/config"
	"github.com/panupakm/boutique-go/tests/mongo"
	serviceclients "github.com/panupakm/boutique-go/tests/service-clients"
)

func SetUp() {
	file, err := os.Open("../config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var conf config.TestConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&conf); err != nil {
		panic(err)
	}

	ctx := context.Background()
	mongo.InitMongoDbClient(ctx, conf)
	serviceclients.InitGrpcClients(ctx, conf)
}

func TearDown() {
	ctx := context.Background()
	serviceclients.CloseGrpcClients(ctx)
	mongo.CloseMongoDBClient(ctx)
}
