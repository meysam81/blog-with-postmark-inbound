package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/datastore"
	"github.com/redis/go-redis/v9"
)

func (*Builder) NewDatastore(ctx context.Context, cfg *config.Config) (datastore.Datastore, error) {
	return newRedis(cfg), nil
}

func newRedis(cfg *config.Config) *redisClient {

	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       int(cfg.RedisDB),
	}

	if cfg.RedisSSL {
		opts.TLSConfig = &tls.Config{
			InsecureSkipVerify: false,
		}
	}

	return &redisClient{r: redis.NewClient(opts)}
}
