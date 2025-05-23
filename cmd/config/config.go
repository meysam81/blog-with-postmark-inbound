package config

import (
	"path/filepath"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"github.com/meysam81/x/config"
)

type cfg struct {
	Port int `koanf:"port"`

	AuthUsername string `koanf:"auth.username"`
	AuthPassword string `koanf:"auth.password"`

	RedisHost     string `koanf:"redis.host"`
	RedisPort     uint32 `koanf:"redis.port"`
	RedisDB       uint8  `koanf:"redis.db"`
	RedisPassword string `koanf:"redis.password"`
	RedisSSL      bool   `koanf:"redis.ssl"`

	RootPath string `koanf:"RootPath"`

	// relative
	DbPath               string `koanf:"dir.db"`
	StoragePath          string `koanf:"dir.storage"`
	AuthorizedEmailsPath string `koanf:"dir.authorized-emails"`

	// absolute
	FrontendPath string `koanf:"dir.frontend"`
}

func (c *cfg) GetDBPath() string {
	return filepath.Join(c.RootPath, c.DbPath)
}

func (c *cfg) GetStoragePath() string {
	return filepath.Join(c.RootPath, c.StoragePath)
}

func (c *cfg) GetAuthorizedEmailsPath() string {
	return filepath.Join(c.RootPath, c.AuthorizedEmailsPath)
}

type Config = cfg

func NewConfig() (*Config, error) {
	c, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	k := c.Koanf()

	// default configs if none is proivded from the upstream
	k.Load(confmap.Provider(map[string]interface{}{
		"port":                  8000,
		"auth.username":         "postmarkapp",
		"auth.password":         "Secr3t!",
		"redis.host":            "localhost",
		"redis.port":            6379,
		"redis.db":              0,
		"RootPath":              ".",
		"dir.db":                "tarzan.db",
		"dir.storage":           "storage",
		"dir.frontend":          "./public",
		"dir.authorized-emails": "deploy/authorized-emails.json",
	}, "."), nil)

	var settings cfg

	// now load the ones coming from env vars
	c.LoadConfig()

	err = k.UnmarshalWithConf("", &settings, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true})
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
