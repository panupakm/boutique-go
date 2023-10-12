package service

import (
	"github.com/go-kratos/kratos/v2/log"
	api "github.com/panupakm/boutique-go/api/user"
	"github.com/panupakm/boutique-go/app/user/internal/biz"
)

// CardService is a card service.
type CardService struct {
	api.UnimplementedUserServer

	uc  *biz.CardUsecase
	log *log.Helper
}

// NewCardService new a card service.
func NewCardService(uc *biz.CardUsecase, logger log.Logger) *CardService {
	return &CardService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "card/service")),
	}
}
