package config

import "github.com/meysam81/x/config"

type settings struct {
	Port uint32

	AuthUsername string
	AuthPassword string

	RedisHost     string
	RedisPort     uint32
	RedisDB       int
	RedisPassword string
	RedisSSL      bool
}

func NewConfig() (*settings, error) {
	c, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	s := &settings{}

	if v := c.Int64(PORT); v != 0 {
		s.Port = uint32(v)
	}

	if v := c.String(BASIC_AUTH_USERNAME); v != "" {
		s.AuthUsername = v
	}
	if v := c.String(BASIC_AUTH_PASSWORD); v != "" {
		s.AuthPassword = v
	}

	if v := c.String(REDIS_HOST); v != "" {
		s.RedisHost = v
	}
	if v := c.Int64(REDIS_PORT); v != 0 {
		s.RedisPort = uint32(v)
	}
	if v := c.Int(REDIS_DB); v != 0 {
		s.RedisDB = v
	}
	if v := c.String(REDIS_PASSWORD); v != "" {
		s.RedisPassword = v
	}
	if v := c.Bool(REDIS_SSL); v {
		s.RedisSSL = v
	}

	return s, nil
}
