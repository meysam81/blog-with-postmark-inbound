package config

import (
	"fmt"
	"strings"

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

	DataStore string `koanf:"datastore"`
	FileStore string `koanf:"filestore"`

	DbPath      string `koanf:"dir.db"`
	StoragePath string `koanf:"dir.storage"`

	BaseUrl string `koanf:"base-url"`

	RSSEndpoint string `koanf:"endpoints.rss"`

	MetricsEnabled      bool   `koanf:"metrics.enabled"`
	MetricsAuthEnabled  bool   `koanf:"metrics.auth.enabled"`
	MetricsAuthUsername string `koanf:"metrics.auth.username"`
	MetricsAuthPassword string `koanf:"metrics.auth.password"`

	DangerouslyAcceptAllSenders bool `koanf:"dangerously-accept-all-senders"`
}

type Config = cfg

func (c *cfg) validate() error {
	errors := []string{}

	if c.DataStore != "sqlite" && c.DataStore != "redis" {
		errors = append(errors, fmt.Sprintf("invalid datastore: %s, must be either 'sqlite' or 'redis'", c.DataStore))
	}

	if c.FileStore != "filesystem" && c.FileStore != "redis" {
		errors = append(errors, fmt.Sprintf("invalid filestore: %s, must be either 'filesystem' or 'redis'", c.FileStore))
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "\n"))
	}
	return nil
}

func NewConfig() (*Config, error) {
	c, err := config.NewConfig(config.WithEnvPrefix("TARZAN_"))
	if err != nil {
		return nil, err
	}

	// default configs if none is proivded from the upstream
	err = c.Koanf().Load(confmap.Provider(map[string]interface{}{
		"base-url":                       "https://tarzan.meysam.io",
		"port":                           8000,
		"auth.username":                  "postmarkapp",
		"auth.password":                  "Secr3t!",
		"redis.host":                     "localhost",
		"redis.port":                     6379,
		"redis.db":                       0,
		"dir.db":                         "tarzan.db",
		"dir.storage":                    "storage",
		"dangerously-accept-all-senders": false,
		"datastore":                      "sqlite",
		"filestore":                      "filesystem",
		"endpoints.rss":                  "/rss.xml",
		"metrics.enabled":                true,
		"metrics.auth.enabled":           false,
	}, "."), nil)
	if err != nil {
		return nil, err
	}

	var settings cfg

	// now load the ones coming from env vars
	err = c.Build()
	if err != nil {
		return nil, err
	}

	err = c.Koanf().UnmarshalWithConf("", &settings, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true})
	if err != nil {
		return nil, err
	}

	if err := settings.validate(); err != nil {
		return nil, err
	}

	return &settings, nil
}
