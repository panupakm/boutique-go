package config

type MongoConfig struct {
	Uri       string   `json:"uri"`
	Databases []string `json:"databases"`
}

type ClientConfig struct {
	GrpcUri       string `json:"grpc_uri"`
	HttpUri       string `json:"http_uri"`
	MongoDatabase string `json:"mongo_database"`
}

type CheckoutConfig struct {
	Uri string `json:"uri"`
}

type TestConfig struct {
	MongoDb MongoConfig             `json:"mongodb"`
	Clients map[string]ClientConfig `json:"clients"`
}
