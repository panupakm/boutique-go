package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type Card struct {
	Id              uuid.UUID
	Name            string
	Number          string
	Ccv             int32
	ExpirationYear  int32
	ExpirationMonth int32
}

// CardRepo is a Greater repo.
type CardRepo interface {
	Create(context.Context, uuid.UUID, *Card) *Card
	FindByID(context.Context, uuid.UUID) *Card
	ListByOwner(context.Context, uuid.UUID) []*Card
	Delete(context.Context, uuid.UUID)
}

type CardUsecase struct {
	repo CardRepo
	log  *log.Helper
}

// NewCardUsecase new a Card usecase.
func NewCardUsecase(repo CardRepo, logger log.Logger) *CardUsecase {
	return &CardUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
