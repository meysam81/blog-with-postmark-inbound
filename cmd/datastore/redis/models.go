package redis

import (
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	r                *redis.Client
	cspCiolationsKey string
}

type Builder struct{}
