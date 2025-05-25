package redis

import (
	"github.com/redis/go-redis/v9"
)

type redisFilestore struct {
	client *redis.Client
}

type Builder struct{}
