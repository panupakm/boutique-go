//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/panupakm/boutique-go/app/email/internal/biz"
	"github.com/panupakm/boutique-go/app/email/internal/conf"
	"github.com/panupakm/boutique-go/app/email/internal/server"
	"github.com/panupakm/boutique-go/app/email/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, newApp))
}
