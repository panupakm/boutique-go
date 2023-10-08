// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/panupakm/boutique-go/app/catalog/internal/biz"
	"github.com/panupakm/boutique-go/app/catalog/internal/conf"
	"github.com/panupakm/boutique-go/app/catalog/internal/data"
	"github.com/panupakm/boutique-go/app/catalog/internal/server"
	"github.com/panupakm/boutique-go/app/catalog/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	database := data.NewMongo(confData)
	dataData, cleanup, err := data.NewData(database, confData, logger)
	if err != nil {
		return nil, nil, err
	}
	productRepo := data.NewProductRepo(dataData, logger)
	catalogUsecase := biz.NewCatalogUsecase(productRepo, logger)
	catalogService := service.NewCatalogService(catalogUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, catalogService, logger)
	httpServer := server.NewHTTPServer(confServer, catalogService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}