package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/panupakm/boutique-go/pkg/util"
)

type User struct {
	Id           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Create(context.Context, *User) (*User, error)
	FindByID(context.Context, uuid.UUID) (*User, error)
	Update(context.Context, *User) (*User, error)
	AddCards(context.Context, uuid.UUID, []*Card)
}

type UserUsecase struct {
	urepo UserRepo
	crepo CardRepo
	log   *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(u UserRepo, c CardRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		urepo: u,
		crepo: c,
		log:   log.NewHelper(logger),
	}
}

// CreateUser creates a Greeter, and returns the new Greeter.
func (uc *UserUsecase) CreateUser(ctx context.Context, password string, user *User) (*User, error) {
	hash, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = hash
	return uc.urepo.Create(ctx, user)
}

// GetUser get user
func (uc *UserUsecase) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	return uc.urepo.FindByID(ctx, id)
}

// AddCard a new card for a user
func (uc *UserUsecase) AddCard(ctx context.Context, userId uuid.UUID, card *Card) *Card {
	return uc.crepo.Create(ctx, userId, card)
}

// ListCardsByOwner a list of cards for a user
func (uc *UserUsecase) ListCardsByOwner(ctx context.Context, userId uuid.UUID) []*Card {
	return uc.crepo.ListByOwner(ctx, userId)
}
