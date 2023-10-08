package config

type MongoConfig struct {
	Uri      string `json:"uri"`
	Database string `json:"database"`
}

type CatalogConfig struct {
	Mongodb MongoConfig
}

type TestConfig struct {
	Catalog CatalogConfig
}
