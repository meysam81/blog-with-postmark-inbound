package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/google/uuid"
	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/filestore"
	"github.com/redis/go-redis/v9"
)

func (*Builder) NewFilestore(ctx context.Context, cfg *config.Config) (filestore.Filestore, error) {
	client, err := newRedisClient(cfg)
	if err != nil {
		return nil, err
	}

	rf := &redisFilestore{client: client}

	if err := rf.runMigrations(ctx); err != nil {
		return nil, err
	}

	return rf, nil
}

func newRedisClient(cfg *config.Config) (*redis.Client, error) {
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

	return redis.NewClient(opts), nil
}

func (r *redisFilestore) runMigrations(ctx context.Context) error {

	if err := r.client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis during migration: %w", err)
	}
	return nil
}

func (r *redisFilestore) Save(data []byte, ext string) (string, error) {

	filename := uuid.New().String() + ext

	key := fmt.Sprintf("filestore:attachment:%s", filename)

	ctx := context.Background()
	if err := r.client.Set(ctx, key, data, 0).Err(); err != nil {
		return "", fmt.Errorf("failed to save file to Redis: %w", err)
	}

	return filename, nil
}

func (r *redisFilestore) Load(filename string) (string, error) {

	key := fmt.Sprintf("filestore:attachment:%s", filename)

	ctx := context.Background()
	content, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("file not found: %s", filename)
		}
		return "", fmt.Errorf("failed to load file from Redis: %w", err)
	}

	return content, nil
}
