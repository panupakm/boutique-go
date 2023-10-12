package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/panupakm/boutique-go/app/user/internal/biz"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent"
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

func entToBiz(in *ent.User, out *biz.User) {
	out.Id = in.ID
	out.Name = in.Name
	out.Email = in.Email
	out.PasswordHash = in.PasswordHash
}
func bizToEnt(in *biz.User, out *ent.User) {
	out.ID = in.Id
	out.Name = in.Name
	out.Email = in.Email
	out.PasswordHash = in.PasswordHash
}

func (u *userRepo) AddCards(ctx context.Context, id uuid.UUID, bcards []*biz.Card) {
	user, err := u.data.db.User.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	ecards := make([]*ent.Card, len(bcards))
	for i, card := range bcards {
		ecards[i] = &ent.Card{}
		cardBizToEnt(card, ecards[i])

	}

	user.Update().AddCards(ecards...).Save(ctx)
}

func (u *userRepo) Create(ctx context.Context, user *biz.User) (*biz.User, error) {
	euser, err := u.data.db.User.Create().
		SetPasswordHash(user.PasswordHash).
		SetEmail(user.Email).
		SetName(user.Name).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	var buser biz.User
	entToBiz(euser, &buser)
	return &buser, nil
}

func (u *userRepo) FindByID(ctx context.Context, id uuid.UUID) (*biz.User, error) {
	euser, err := u.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	var buser biz.User
	entToBiz(euser, &buser)
	return &buser, nil
}

func (u *userRepo) Update(ctx context.Context, user *biz.User) (*biz.User, error) {
	update := u.data.db.User.UpdateOneID(user.Id)
	if user.Name != "" {
		update.SetName(user.Name)
	}
	if user.Email != "" {
		update.SetEmail(user.Email)
	}
	if user.PasswordHash != "" {
		update.SetPasswordHash(user.PasswordHash)
	}

	euser, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	var buser biz.User
	entToBiz(euser, &buser)
	return &buser, nil
}
