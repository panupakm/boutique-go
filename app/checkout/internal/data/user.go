package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"

	userapi "github.com/panupakm/boutique-go/api/user"
	"github.com/panupakm/boutique-go/app/checkout/internal/biz"
	user "github.com/panupakm/boutique-go/pkg/user"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "cart/data")),
	}
}

func (r *userRepo) Get(ctx context.Context, userId string) (*user.User, error) {
	reply, err := r.data.userc.GetUser(ctx, &userapi.GetUserReq{
		Id: userId,
	})

	if err != nil {
		return nil, err
	}

	return &user.User{
		Id:   uuid.MustParse(reply.Id),
		Name: reply.Username,
	}, nil
}
