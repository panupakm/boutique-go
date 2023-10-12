package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/panupakm/boutique-go/api/shared"
	api "github.com/panupakm/boutique-go/api/user"
	"github.com/panupakm/boutique-go/app/user/internal/biz"
)

// UserService is a user service.
type UserService struct {
	api.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "user/service")),
	}
}

// CreateUser implements user.userServer.
func (u *UserService) CreateUser(ctx context.Context, req *api.CreateUserReq) (*api.CreateUserReply, error) {
	user, err := u.uc.CreateUser(ctx, req.Password, &biz.User{
		Name:  req.Username,
		Email: req.Email,
	})

	if err != nil {
		fmt.Printf("Create user failed: %s\n", err)
		return nil, err
	}

	return &api.CreateUserReply{
		Id:       user.Id.String(),
		Username: user.Name,
	}, nil
}

// AddCard implements user.AddCard
func (u *UserService) AddCard(ctx context.Context, in *api.AddCardReq) (*api.AddCardReply, error) {
	defer func() {
		if r := recover(); r != nil {
			u.log.Infof("Recovered from panic:", r)
		}
	}()

	card := u.uc.AddCard(ctx, uuid.MustParse(in.UserId), &biz.Card{
		Name:            in.Name,
		Number:          in.CardNumber,
		Ccv:             in.Ccv,
		ExpirationYear:  in.ExpirationYear,
		ExpirationMonth: in.ExpirationMonth,
	})
	return &api.AddCardReply{
		CardId: card.Id.String(),
	}, nil
}

// ListCards implements card.ListCards.
func (u *UserService) ListCards(ctx context.Context, in *api.ListCardsReq) (*api.ListCardsReply, error) {
	cards := u.uc.ListCardsByOwner(ctx, uuid.MustParse(in.UserId))
	var replyCards = make([]*shared.CreditCardInfo, len(cards))
	for i, card := range cards {
		replyCards[i] = &shared.CreditCardInfo{
			CreditCardNumber:          card.Number,
			CreditCardCvv:             card.Ccv,
			CreditCardExpirationYear:  card.ExpirationYear,
			CreditCardExpirationMonth: card.ExpirationMonth,
		}
	}
	return &api.ListCardsReply{
		Cards: replyCards,
	}, nil
}
