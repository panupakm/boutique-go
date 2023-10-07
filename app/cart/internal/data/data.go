package data

import (
	"context"
	"time"

	"github.com/panupakm/boutique-go/app/cart/internal/conf"
	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRedisCmd, NewCartRepo)

// Data .
type Data struct {
	redisCli redis.Cmdable
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "cart/data/redis"))
	log.Infof("redis addr: %s", conf.Redis.Addr)
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

// NewData .
func NewData(redisCmd redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	d := &Data{
		redisCli: redisCmd,
	}
	log.Infof("redisCli:%v", redisCmd)
	return d, func() {
		log.NewHelper(logger).Info("closing the data resources")
	}, nil
}
