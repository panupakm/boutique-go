package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"

	api "github.com/panupakm/boutique-go/api/email"
	"github.com/panupakm/boutique-go/app/email/internal/conf"
	"github.com/panupakm/boutique-go/app/email/internal/service"
	otel "github.com/panupakm/boutique-go/pkg/otel"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, emailer *service.EmailService, logger log.Logger) *http.Server {

	otel.SetupOTelSDK(context.Background(), "email", "1.0.0")

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	api.RegisterEmailHTTPServer(srv, emailer)

	return srv
}
