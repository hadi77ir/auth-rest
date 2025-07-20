package config

import (
	"auth-rest/internal/app"
	"github.com/hadi77ir/go-env"
)

const DefaultConfigPath = "${WEB_CONFIG_PATH:-config.yaml}"

type ConfigKeyType int

const ConfigKey ConfigKeyType = 0xCCFF00

func FromGlobals(globals *app.AppGlobals) *Config {
	return app.Value[*Config](globals, ConfigKey)
}

func Setup(globals *app.AppGlobals, configPath string) (*Config, error) {
	if configPath == "" {
		var err error
		configPath, err = env.ExpandEnv(DefaultConfigPath)
		if err != nil {
			return nil, err
		}
	}
	cfg, err := ReadConfig[Config](configPath)
	if err != nil {
		return nil, err
	}
	globals.Set(ConfigKey, cfg)
	return cfg, nil
}
