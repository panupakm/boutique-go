package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/panupakm/boutique-go/app/user/internal/biz"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent/card"
	"github.com/panupakm/boutique-go/app/user/internal/data/ent/user"
)

type cardRepo struct {
	data *Data
	log  *log.Helper
}

func NewCardRepo(data *Data, logger log.Logger) biz.CardRepo {
	return &cardRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "card/repo")),
	}
}

func cardEntToBiz(in *ent.Card, out *biz.Card) {
	out.Id = in.ID
	out.Number = in.Number
	out.Ccv = int32(in.Ccv)
}
func cardBizToEnt(in *biz.Card, out *ent.Card) {
	out.ID = in.Id
	out.Number = in.Number
	out.Ccv = int(in.Ccv)
}

func (c *cardRepo) Create(ctx context.Context, ownerId uuid.UUID, card *biz.Card) *biz.Card {
	po, err := c.data.db.Card.Create().
		SetNumber(card.Number).
		SetCcv(int(card.Ccv)).
		SetOwnerID(ownerId).
		SetName(card.Name).
		SetExpirationYear(int(card.ExpirationYear)).
		SetExpirationMonth(int(card.ExpirationMonth)).
		Save(ctx)

	if err != nil {
		c.log.Errorf("Create card failed: %v", err)
		panic(err)
	}

	return &biz.Card{
		Id: po.ID,
	}
}

func (c *cardRepo) FindByID(ctx context.Context, id uuid.UUID) *biz.Card {
	ecard, err := c.data.db.Card.Get(ctx, id)
	if err != nil {
		panic(err)
	}

	bcard := biz.Card{}
	cardEntToBiz(ecard, &bcard)
	return &bcard
}

func (c *cardRepo) ListByOwner(ctx context.Context, id uuid.UUID) []*biz.Card {
	ecards, err := c.data.db.Card.Query().
		Where(card.HasOwnerWith(user.ID(id))).
		All(ctx)
	if err != nil {
		panic(err)
	}

	bcards := make([]*biz.Card, len(ecards))
	for i, c := range ecards {
		bcards[i] = &biz.Card{}
		cardEntToBiz(c, bcards[i])
	}
	return bcards
}

func (c *cardRepo) Delete(ctx context.Context, id uuid.UUID) {
	_, err := c.data.db.Card.Delete().Where(card.ID(id)).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
