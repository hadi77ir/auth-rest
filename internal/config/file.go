package config

import (
	"errors"
	"os"

	"github.com/hadi77ir/go-env"
	"gopkg.in/yaml.v3"
)

func ReadConfig[T any](path string) (*T, error) {
	if path == "" {
		return nil, errors.New("no config file was found")
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	expanded, err := env.ExpandEnv(string(file))
	if err != nil {
		return nil, err
	}
	file = []byte(expanded)
	var cfg T
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func WriteConfig[T any](path string, cfg *T) error {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}
