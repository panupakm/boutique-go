package data

import (
	"context"
	"encoding/json"

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
	if result != "" {
		err = json.Unmarshal([]byte(result), cacheCart)
	}

	if err != nil && err != redis.Nil {
		r.log.Errorf("Result %s Error Unmarshal %s", result, err.Error())
		return err
	}

	updated := false
	for i, cartItem := range cacheCart.Items {
		if cartItem.ProductId == item.ProductId {
			cacheCart.Items[i].Quantity += item.Quantity
			updated = true
			break
		}
	}

	if !updated {
		cacheCart.Items = append(cacheCart.Items, biz.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	marshal, err := json.Marshal(cacheCart)
	if err != nil {
		r.log.Errorf("fail to set cart cache:json.Marshal(%v) error(%v)", cacheCart, err)
		return err
	}

	err = r.data.redisCli.Set(ctx, key, string(marshal), -1).Err()
	if err != nil {
		r.log.Errorf("fail to set cart cache:redis.Set(%v) error(%v)", cacheCart, err)
		return err
	}

	return nil
}

func (r *cartRepo) Get(ctx context.Context, userId string) (*biz.Cart, error) {
	key := toCartKey(userId)
	result, err := r.data.redisCli.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return &biz.Cart{}, nil
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

func (r *cartRepo) Empty(ctx context.Context, userId string) error {
	key := toCartKey(userId)
	r.data.redisCli.Del(context.Background(), key)
	return nil
}
