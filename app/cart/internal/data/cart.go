package data

import (
	"context"
	"encoding/json"
	"time"

	"github.com/panupakm/boutique-go/app/cart/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type cartRepo struct {
	data *Data
	log  *log.Helper
}

// NewCartRepo .
func NewCartRepo(data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func toCartKey(userId string) string {
	return "cart:" + userId
}

func (r *cartRepo) AddItem(ctx context.Context, userId string, item *biz.CartItem) error {
	key := toCartKey(userId)
	result, err := r.data.redisCli.Get(ctx, key).Result()

	if err != nil && err != redis.Nil {
		r.log.Errorf("Result %s Error getting %s", result, err.Error())
		return err
	}
	var cacheCart = &biz.Cart{}
	err = json.Unmarshal([]byte(result), cacheCart)

	cacheCart.Items = append(cacheCart.Items, biz.CartItem{
		ProductId: item.ProductId,
		Quantity:  item.Quantity,
	})

	marshal, err := json.Marshal(cacheCart)
	if err != nil {
		r.log.Errorf("fail to set cart cache:json.Marshal(%v) error(%v)", cacheCart, err)
	}

	err = r.data.redisCli.Set(ctx, key, string(marshal), time.Minute*30).Err()
	if err != nil {
		r.log.Errorf("fail to set cart cache:redis.Set(%v) error(%v)", cacheCart, err)
	}

	if err != nil {
		return err
	}
	return nil
}

func (r *cartRepo) GetCart(ctx context.Context, userId string) (*biz.Cart, error) {
	key := toCartKey(userId)
	result, err := r.data.redisCli.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	var cacheCart = &biz.Cart{}
	err = json.Unmarshal([]byte(result), cacheCart)
	if err != nil {
		r.log.Errorf("Unmarshal cart: %w", err)
		return nil, nil
	}

	return cacheCart, nil
}

func (r *cartRepo) Empty(context.Context, string) error {
	return nil
}
