// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/panupakm/boutique-go/app/email/internal/biz"
	"github.com/panupakm/boutique-go/app/email/internal/conf"
	"github.com/panupakm/boutique-go/app/email/internal/server"
	"github.com/panupakm/boutique-go/app/email/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	emailUseCase := biz.NewEmailUseCase(data, logger)
	emailService := service.NewEmailService(emailUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, emailService, logger)
	httpServer := server.NewHTTPServer(confServer, emailService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}