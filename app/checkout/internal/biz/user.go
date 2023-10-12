package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panupakm/boutique-go/pkg/user"
)

type UserRepo interface {
	Get(ctx context.Context, userId string) (*user.User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "user/biz"))}
}

func (uc *UserUseCase) Get(ctx context.Context, id string) (*user.User, error) {
	fmt.Println("--------------------------------")
	fmt.Printf("%s\n", id)
	fmt.Printf("%v\n", uc)
	fmt.Printf("\n\n\n")
	return uc.repo.Get(ctx, id)
}
