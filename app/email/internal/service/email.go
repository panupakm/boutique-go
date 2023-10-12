package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	api "github.com/panupakm/boutique-go/api/email"
	"github.com/panupakm/boutique-go/app/email/internal/biz"
)

type EmailService struct {
	api.UnimplementedEmailServer

	ec  *biz.EmailUseCase
	log *log.Helper
}

func NewEmailService(ec *biz.EmailUseCase, logger log.Logger) *EmailService {
	return &EmailService{
		ec:  ec,
		log: log.NewHelper(log.With(logger, "module", "email/service")),
	}
}

func (s *EmailService) SendOrderConfirmation(ctx context.Context, req *api.SendOrderConfirmationRequest) (*api.Empty, error) {
	err := s.ec.SendOrderConfirmation(ctx, req.GetEmail(), req)
	return &api.Empty{}, err
}
