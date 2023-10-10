package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/panupakm/boutique-go/app/user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (u *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	return nil, nil
}

func (u *userRepo) FindByID(context.Context, string) (*biz.User, error) {
	return nil, nil
}
func (u *userRepo) Update(context.Context, *biz.User) (*biz.User, error) {
	return nil, nil
}
